package models

import (
	"github.com/example/simple-REST/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Film struct {
	gorm.Model
	Id          string `json:"id"`
	Name        string `gorm:""json:"name"`
	Director      string `json:"director"`
	YearOfCreating int `json:"yearOfCreating"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Film{})
}

func (b *Film) CreateFilm() *Film {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllFilms() []Film {
	var Films []Film
	db.Find(&Films)
	return Films
}

func GetFilmById(Id int64) (*Film , *gorm.DB){
	var getFilm Film
	db:=db.Where("ID = ?", Id).Find(&getFilm)
	return &getFilm, db
}

func DeleteFilm(ID int64) Film {
	var Film Film
	db.Where("ID = ?", ID).Delete(Film)
	return Film