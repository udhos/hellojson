package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	// load json
	bufIn, errRead := ioutil.ReadFile("input.json")
	if errRead != nil {
		log.Fatal(errRead)
	}
	log.Printf("loaded: %s", bufIn)

	type Message struct {
		Name, Text string
	}

	var msgs []*Message

	// parse json into msgs
	if err := json.Unmarshal(bufIn, &msgs); err != nil {
		log.Fatal(err)
	}

	log.Printf("parsed into struct: %v", msgs)

	for _, m := range msgs {
		m.Text = strings.ToUpper(m.Text)
	}

	log.Printf("struct modified: %v", msgs)

	// encode msgs as json
	bufOut, errEnc := json.Marshal(&msgs)
	if errEnc != nil {
		log.Fatal(errEnc)
	}

	log.Printf("struct encoded as json: %s", bufOut)

	// save json
	if err := ioutil.WriteFile("output.json", bufOut, 0644); err != nil {
		log.Fatal(err)
	}

	log.Printf("saved")
}
