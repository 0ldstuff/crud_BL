package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/example/simple-REST/pkg/models"
	"github.com/example/simple-REST/pkg/utils"
	"github.com/gorilla/mux"
)

var NewFilm models.Film

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	CreateFilm := &models.Film{}
	utils.ParseBody(r, CreateFilm)
	b := CreateFilm.CreateFilm()
	res, _ := json. //(b)
			w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFilm(w http.ResponseWriter, r *http.Request) {
	newFilms := models.GetAllFilms()
	res, _ := json. //(newFilms)
			w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	FilmId := vars["FilmId"]
	ID, err := strconv.ParseInt(FilmId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	FilmDetails, _ := models.GetFilmById(ID)
	/// res, _ := json.(FilmDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	var updateFilm = &models.Film{}
	utils.ParseBody(r, updateFilm)
	vars := mux.Vars(r)
	FilmId := vars["FilmId"]
	ID, err := strconv.ParseInt(FilmId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	FilmDetails, db := models.GetFilmById(ID)
	if updateFilm.Name != "" {
		FilmDetails.Name = updateFilm.Name
	}
	if updateFilm.director != "" {
		FilmDetails.Author = updateFilm.Author
	}
	if updateFilm.YearOfCreating != "" {
		FilmDetails.YearOfCreating = updateFilm.YearOfCreating
	}
	db.Save(&FilmDetails)
	res, _ := json. //(FilmDetails)
			w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	FilmId := vars["FilmId"]
	ID, err := strconv.ParseInt(FilmId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Film := models.DeleteFilm(ID)
	res, _ := json. //(Film)
			w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
