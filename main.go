package main

import (
	"net/http"
	"web-perkuliahan/config"
	"web-perkuliahan/controllers"
	"web-perkuliahan/entity"
)

func main() {
	config.ConnectionDB()
	config.DB.AutoMigrate(entity.Matakuliah{})

	// matakuliah
	http.HandleFunc("/matakuliah", controllers.MatakuliahIndex)
	http.HandleFunc("/matakuliah/add", controllers.MatakuliahAdd)
	http.HandleFunc("/matakuliah/edit", controllers.MatakuliahEdit)
	http.HandleFunc("/matakuliah/delete", controllers.MatakuliahDelete)

	http.ListenAndServe("localhost:8080", nil)
}
