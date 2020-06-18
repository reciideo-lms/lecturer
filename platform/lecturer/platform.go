package lecturer

import (
	"database/sql"
	"github.com/google/uuid"
)

type Platform struct {
	Id       uuid.UUID `json:"id"`
	Platform string    `json:"platform"`
	URL      string    `json:"url"`
}

func initTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS platform (
		id UUID PRIMARY KEY UNIQUE,
		lecturerId UUID NOT NULL,
		platform TEXT NOT NULL,
		url TEXT NOT NULL,
		FOREIGN KEY (lecturerId) REFERENCES lecturer(id)
	);`)
	return err
}

func (r *Repo) addPlatform(lecturer Lecturer, platform Platform) (Platform, error) {
	platform.Id = uuid.New()
	sqlStatement := `
		INSERT INTO platform (id, lecturerId, platform, url)
		VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(sqlStatement, platform.Id, lecturer.Id, platform.Platform, platform.URL)
	if err != nil {
		return Platform{}, err
	}
	return platform, nil
}

func (r *Repo) getPlatforms(item *Lecturer) ([]Platform, error) {
	platforms := make([]Platform, 0)
	rows, err := r.DB.Query("SELECT id, platform, url FROM platform WHERE lecturerId=$1", item.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		platform := Platform{}
		err = rows.Scan(&platform.Id, &platform.Platform, &platform.URL)
		if err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}
	return platforms, nil
}

func (r *Repo) deletePlatforms(id string) error {
	_, err := r.DB.Exec("DELETE FROM platform WHERE lecturerId=$1", id)
	return err
}
