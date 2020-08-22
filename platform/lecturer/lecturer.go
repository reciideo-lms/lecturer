package lecturer

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/reciideo-lms/lecturer/config"
	"github.com/reciideo-lms/lecturer/utils"
	"time"
)

type Lecturer struct {
	ID          string     `json:"id" sql:"type:uuid;primary_key"`
	Forename    string     `json:"forename"`
	Surname     string     `json:"surname"`
	Username    string     `json:"username"`
	Description string     `json:"description"`
	Platforms   []Platform `json:"platforms" gorm:"ForeignKey:LecturerID"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

type Platform struct {
	ID         string `json:"id" gorm:"primary_key"`
	LecturerID string `json:"-"`
	Platform   string `json:"platform"`
	URL        string `json:"url"`
}

func New() error {
	return config.DB.AutoMigrate(&Platform{}, &Lecturer{}).Error
}

func Add(item Lecturer) (Lecturer, error) {
	username, err := slugUsername(item.Forename, item.Surname)
	if err != nil {
		return Lecturer{}, err
	}
	item.Username = username
	if err = config.DB.Create(&item).Error; err != nil {
		return Lecturer{}, err
	}
	return item, nil
}

func GetAll() ([]Lecturer, error) {
	var lecturer []Lecturer
	if err := config.DB.Find(&lecturer).Error; err != nil {
		return nil, err
	} // TODO get related
	return lecturer, nil
}

func GetSingle(uuid string) (Lecturer, error) {
	var item Lecturer
	if err := config.DB.Where("id = ?", uuid).First(&item).Related(&item.Platforms).Error; err != nil {
		return Lecturer{}, err
	} // TODO get related
	return item, nil
}

func Delete(uuid string) error {
	var item Lecturer
	if err := config.DB.Where("id = ?", uuid).Delete(&item).Error; err != nil {
		return err
	}
	return nil
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

func (lecturer *Lecturer) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New().String())
}

func (platform *Platform) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New().String())
}
