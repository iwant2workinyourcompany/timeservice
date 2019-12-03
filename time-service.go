package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Time struct {
	Time     string `json:"time,omitempty"`
}

func main() {
	log.Print("Начало работы сервиса time-service")

	http.HandleFunc("/gettime", serveTime)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func serveTime(w http.ResponseWriter, r *http.Request) {
	log.Print("Вызов функции serveTime()")
	var serverTime Time
	serverTime.Time = time.Now().String()
	json.NewEncoder(w).Encode(getMSKTime())
}

func getMSKTime() string {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}

	timeInUTC := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
  
  	return timeInUTC.In(location).String()
}

