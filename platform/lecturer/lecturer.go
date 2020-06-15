package lecturer

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/reciideo-lms/lecturer/utils"
	"log"
	"time"
)

type Lecturer struct {
	Id          uuid.UUID `json:"id"`
	Forename    string    `json:"forename"`
	Surname     string    `json:"surname"`
	Username    string    `json:"username"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Repo struct {
	DB *sql.DB
}

func New(db *sql.DB) *Repo {
	_, err := db.Exec(` CREATE TABLE IF NOT EXISTS lecturer (
 		id UUID PRIMARY KEY UNIQUE,
 		forename TEXT,
 		surname TEXT,
 		username TEXT UNIQUE,
 		description TEXT,
 		updatedAt TIMESTAMP,
 		createdAt TIMESTAMP
	 )`)
	if err != nil {
		log.Fatal(err)
	}
	return &Repo{
		DB: db,
	}
}

func (r *Repo) Add(item Lecturer) (Lecturer, error) {
	item.Id = uuid.New()
	username, err := slugUsername(item.Forename, item.Surname)
	if err != nil {
		return Lecturer{}, err
	}
	item.Username = username
	item.CreatedAt = time.Now()
	item.UpdatedAt = item.CreatedAt

	sqlStatement := `
		INSERT INTO lecturer (id, forename, surname, username, description, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = r.DB.Exec(sqlStatement, item.Id, item.Forename, item.Surname, item.Username, item.Description, item.UpdatedAt, item.CreatedAt)
	if err != nil {
		return Lecturer{}, err
	}
	return item, nil
}

func (r *Repo) GetAll() ([]*Lecturer, error) {
	rows, err := r.DB.Query("SELECT * FROM lecturer")
	if err != nil {
		return nil, err
	}

	items := make([]*Lecturer, 0)
	for rows.Next() {
		item := new(Lecturer)
		err = rows.Scan(&item.Id, &item.Forename, &item.Surname, &item.Username, &item.Description, &item.UpdatedAt, &item.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func slugUsername(forename string, surname string) (string, error) {
	sluggedForename, err := utils.SlugString(forename)
	if err != nil {
		return "", err
	}
	sluggedSurname, err := utils.SlugString(surname)
	if err != nil {
		return "", err
	}
	return utils.ConcatStrings(sluggedForename, sluggedSurname), nil
}
