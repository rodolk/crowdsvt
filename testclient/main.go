package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const getURL = "https://mgaeogoxuc.execute-api.us-west-2.amazonaws.com/default/vtSignURL"

func main() {
	resp, err := http.Get(getURL)
	if err != nil {
		log.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	log.Println(string(bodyBytes))

	//part1 := bodyBytes[:4]
	//part2 := bodyBytes[5:]
	//aux2 := append(part1, part2...)
	//log.Println(string(bodyBytes))
	log.Println("---------------------------------")
	//log.Println(string(aux2))

	putURL := string(bodyBytes)

	filename := "test.txt"
	var fileBytes []byte

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Couldn't open file %v to upload. Here's why: %v\n", filename, err)
		os.Exit(1)
	} else {
		defer file.Close()
		fileBytes, err = io.ReadAll(file)
		if err != nil {
			log.Fatalf("Couldn't read file %v. Here's why: %v\n", filename, err)
			os.Exit(1)
		}
	}

	log.Println("-----------------------------------")
	putData := bytes.NewBuffer(fileBytes)
	putResp, err := http.NewRequest(http.MethodPut, putURL, putData)
	if err != nil {
		log.Fatalf("Failed to create PUT request: %v", err)
	}

	for name, values := range putResp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	client := &http.Client{}
	log.Println(client.Transport)
	response, err := client.Do(putResp)
	if err != nil {
		log.Fatalf("Failed to send PUT request: %v", err)
	}
	defer response.Body.Close()

	putRespBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(putRespBody))
}
