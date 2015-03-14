package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	User       string
	OauthToken string
}

func loadConfig(filename string) Config {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	render := bufio.NewReader(f)
	dec := json.NewDecoder(render)

	var c Config
	err = dec.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}
