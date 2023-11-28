package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var ageArray = []int{28, 27, 18, 19, 30, 36}

func register() {
	username := generateRandomString(6)
	password := generateRandomString(12)
	indexAge := pickRandomAgeInArray(ageArray)
	user := User{
		Username:    username,
		Password:    password,
		Location:    "Hanoi",
		Description: "Generate by worker",
		Age:         ageArray[indexAge],
	}

	bodyReader, err := json.Marshal(user)
	if err != nil {
		log.Printf("error when parse request data: %s\n", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, configData.Api.Register, bytes.NewBuffer(bodyReader))
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		return
	}
	log.Printf("register with username: %s, status_code: %d\n", username, res.StatusCode)

}

func main() {
	log.Println("worker started.")
	log.Println("start read config file...")
	readConfigFile("config.yaml")
	for {
		interval := pickRandomInterval()
		register()
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
