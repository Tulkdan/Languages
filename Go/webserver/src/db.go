package src

import (
	"database/sql"
	"encoding/json"
)

type DB struct {
	connection *sql.DB
}

func CreateDB(db *sql.DB) (*DB, error) {
	create := `
    CREATE TABLE IF NOT EXISTS people (
        id        TEXT NOT NULL PRIMARY KEY,
        name      TEXT,
        nickname  TEXT,
        birthdate DATETIME,
        stacks    JSONB
    );`

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return &DB{connection: db}, nil
}

func (db *DB) InsertPeople(person Person) (int64, error) {
	db.connection.Ping()
	sql := "INSERT INTO people VALUES(?,?,?,?,?);"
	statement, err := db.connection.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	stack, _ := json.Marshal(person.Stack)

	res, err := statement.Exec(person.Id, person.Name, person.Nickname, person.Birthdate, stack)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) GetPersonById(id string) (Person, error) {
	row := db.connection.QueryRow(`
        SELECT
            id,
            name,
            nickname,
            birthdate,
            stacks
        FROM people
        WHERE id = ?;`, id)

	person := Person{}

	var err error
	if err = row.Scan(&person.Id, &person.Name, &person.Nickname, &person.Birthdate, &person.Stack); err != nil {
		return Person{}, err
	}

	return person, nil
}

func (db *DB) SearchPerson(param string) ([]Person, error) {
	rows, err := db.connection.Query(`
        SELECT
            id,
            name,
            nickname,
            birthdate,
            stacks
        FROM people
        WHERE LOWER(name) LIKE '%' || ? || '%'
            OR LOWER(nickname) LIKE '%' || ? || '%'
            OR LOWER(stacks) LIKE '%' || ? || '%';`, param, param, param)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := []Person{}
	for rows.Next() {
		i := Person{}
		if err = rows.Scan(&i.Id, &i.Name, &i.Nickname, &i.Birthdate, &i.Stack); err != nil {
			return nil, err
		}

		people = append(people, i)
	}

	return people, nil
}

func (db *DB) CountAllPeople() (int, error) {
	row := db.connection.QueryRow(`SELECT COUNT(*) FROM people;`)

	var value int
	if err := row.Scan(&value); err != nil {
		return 0, err
	}

	return value, nil
}
