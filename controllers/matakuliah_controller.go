package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web-perkuliahan/entity"
	"web-perkuliahan/models"
)

func MatakuliahIndex(w http.ResponseWriter, r *http.Request) {
	matakuliahs := models.GetAll()
	data := map[string]any{
		"matakuliahs": &matakuliahs,
	}

	files, err := template.ParseFiles("views/matakuliah/index.html")
	if err != nil {
		panic(err)
	}

	log.Println("| 200 | GET | Matakuliah.Index |")
	files.Execute(w, &data)
}
func MatakuliahAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("views/matakuliah/create.html")
		if err != nil {
			panic(err)
		}

		log.Println("| 200 | GET | Matakuliah.Add |")
		files.Execute(w, nil)
	}

	if r.Method == "POST" {
		var matakuliah entity.Matakuliah
		matakuliah.KodeMk = r.FormValue("KodeMk")
		matakuliah.NamaMk = r.FormValue("NamaMk")
		value := r.FormValue("Sks")
		sks, err := strconv.Atoi(value)
		if err != nil {
			sks = 0
			log.Println("error parse atoi", err.Error())
		}
		matakuliah.Sks = sks

		ok := models.Create(matakuliah)
		if !ok {
			files, err := template.ParseFiles("views/matakuliah/create.html")
			if err != nil {
				panic(err)
			}
			files.Execute(w, nil)
		}

		log.Println("| 203 | POST | Matakuliah.Add |")
		http.Redirect(w, r, "/matakuliah", http.StatusSeeOther)
	}
}
func MatakuliahEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		files, err := template.ParseFiles("views/matakuliah/update.html")
		if err != nil {
			panic(err)
		}
		id := r.URL.Query().Get("id")
		data := models.GetById(id)

		log.Println("| 200 | GET | Matakuliah.Edit |")
		files.Execute(w, data)
	}

	if r.Method == "POST" {
		var matakuliah entity.Matakuliah
		matakuliah.KodeMk = r.FormValue("KodeMk")
		matakuliah.NamaMk = r.FormValue("NamaMk")
		value := r.FormValue("Sks")
		sks, err := strconv.Atoi(value)
		if err != nil {
			sks = 0
			log.Println("error parse atoi", err.Error())
		}
		matakuliah.Sks = sks

		ok := models.Update(matakuliah)
		if !ok {
			files, err := template.ParseFiles("views/matakuliah/update.html")
			if err != nil {
				panic(err)
			}
			files.Execute(w, nil)
		}

		log.Println("| 200 | PUT | Matakuliah.Edit |")
		http.Redirect(w, r, "/matakuliah", http.StatusSeeOther)
	}
}
func MatakuliahDelete(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	err := models.Delete(id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/matakuliah", http.StatusSeeOther)
}
