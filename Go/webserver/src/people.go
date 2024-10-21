package src

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Skills []string

func (s Skills) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return fmt.Sprintf(`["%s"]`, strings.Join(s, `","`)), nil
}

func (s *Skills) Scan(src interface{}) (err error) {
	var skills []string
	switch src.(type) {
	case string:
		err = json.Unmarshal([]byte(src.(string)), &skills)
	case []byte:
		err = json.Unmarshal(src.([]byte), &skills)
	default:
		return errors.New("Incompatible type for Skills")
	}
	if err != nil {
		return
	}
	*s = skills
	return nil
}

type InsertPerson struct {
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Birthdate time.Time `json:"birthdate"`
	Stack     Skills    `json:"stack"`
}

type Person struct {
	Id        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Nickname  string    `db:"nickname" json:"nickname"`
	Birthdate time.Time `db:"birthdate" json:"birthdate"`
	Stack     Skills    `db:"stack" json:"stack"`
}

type People struct {
	people []Person
}

func (p *People) Insert(db *DB, person InsertPerson) (string, error) {
	id, _ := uuid.NewV7()
	newPerson := Person{
		Id:        id.String(),
		Name:      person.Name,
		Nickname:  person.Nickname,
		Birthdate: person.Birthdate,
		Stack:     person.Stack,
	}

	if _, err := db.InsertPeople(newPerson); err != nil {
		return "", err
	}

	return id.String(), nil
}

var ErrIDNotFound = fmt.Errorf("ID not found")

func (p *People) Get(db *DB, id string) (Person, error) {
	d, err := db.GetPersonById(id)
	if err != nil {
		return Person{}, err
	}

	return d, nil
}
