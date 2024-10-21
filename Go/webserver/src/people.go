package src

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type InsertPerson struct {
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Birthdate time.Time `json:"birthdate"`
	Stack     []string  `json:"stack"`
}

type Person struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Birthdate time.Time `json:"birthdate"`
	Stack     []string  `json:"stack"`
}

type People struct {
	people []Person
}

func (p *People) Insert(person InsertPerson) string {
	id, _ := uuid.NewV7()
	newPerson := Person{
		Id:        id.String(),
		Name:      person.Name,
		Nickname:  person.Nickname,
		Birthdate: person.Birthdate,
		Stack:     person.Stack,
	}

	p.people = append(p.people, newPerson)

	return id.String()
}

var ErrIDNotFound = fmt.Errorf("ID not found")

func (p *People) Get(id string) (Person, error) {
	for _, p := range p.people {
		if p.Id == id {
			return p, nil
		}
	}

	return Person{}, ErrIDNotFound
}
