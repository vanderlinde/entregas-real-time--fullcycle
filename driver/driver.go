package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func loadDrivers(file string) []byte {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error())
	}

	return data
}

func ListDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := loadDrivers("drivers.json")
	w.Write([]byte(drivers))
	http.ListenAndServe(":8081", nil)
}