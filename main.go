package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type config struct {
	UserName string `json:"user_name"`
	IntroText string `json:"intro_text"`
	ProjectList []string `json:"project_list"`
}

func main() {
	file, err := os.Open(".showcase.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	var config config
	if err = decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}
