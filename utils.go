package main

import (
	"encoding/csv"
	"gopkg.in/yaml.v3"
	"log"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	Api struct {
		Register string `yaml:"register"`
	} `yaml:"api"`
	Config struct {
		Interval []float32 `yaml:"interval"`
	} `yaml:"config"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func pickRandomAgeInArray(arr []int) int {
	res := seededRand.Intn(len(arr))
	return res
}

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Age         int    `json:"age"`
}

func pickRandomInterval() float32 {
	randomIndex := rand.Intn(len(configData.Config.Interval))
	pick := configData.Config.Interval[randomIndex]
	return pick
}

func pickRandomElementIn2dArray(data [][]string) []string {
	randomIndex := rand.Intn(len(data))
	pick := data[randomIndex]
	return pick

}

func readDataFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	return data
}

func readConfigFile(filename string) {
	config := Config{}
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error when read config file: %v", err.Error())
	}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	configData = &config
}
