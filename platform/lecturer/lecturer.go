package lecturer

import (
	"database/sql"
	"log"
)

type Lecturer struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
}

type Repo struct {
	DB *sql.DB
}

func New(db *sql.DB) *Repo {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS lecturer (forename TEXT, surname TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	return &Repo{
		DB: db,
	}
}

func (r *Repo) Add(item Lecturer) (Lecturer, error) {
	sqlStatement := `
		INSERT INTO lecturer (forename, surname)
		VALUES ($1, $2)`
	_, err := r.DB.Exec(sqlStatement, item.Forename, item.Surname)
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
		err = rows.Scan(&item.Forename, &item.Surname)
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
