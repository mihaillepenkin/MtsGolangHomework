package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type decodeResponse struct {
	InputString string `json:"inputString"`
}

type decodeResult struct {
	OutputString string `json:"outputString"`
}

const addrs string = "http://localhost:8081/"

func Run() {
	client := http.Client{}


	//1
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", addrs + "version", nil)
	if (err != nil) {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if (err != nil) {
		log.Fatal(err)
	}
	defer func() {
		err = res.Body.Close()
		if (err != nil) {
			log.Fatal(err)
		}	
	}()
	body, err := io.ReadAll(res.Body)
	if (err != nil) {
		log.Fatal(err)
	}
	log.Println(string(body))


	//2
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	data, err := json.Marshal(decodeResponse{InputString: "SGVsbG8sIFdvcmxkIQ=="})
	if (err != nil) {
		log.Fatal(err)
	}
	req, err = http.NewRequestWithContext(ctx, "POST", addrs + "decode", bytes.NewBuffer(data))
	if (err != nil) {
		log.Fatal(err)
	}
	res, err = client.Do(req)
	if (err != nil) {
		log.Fatal(err)
	}
	defer func() {
		err = res.Body.Close()
		if (err != nil) {
			log.Fatal(err)
		}	
	}()
	body, err = io.ReadAll(res.Body)
	if (err != nil) {
		log.Fatal(err)
	}
	var output decodeResult
	err = json.Unmarshal(body, &output)
	if (err != nil) {
		log.Fatal(err)
	}
	log.Println(output.OutputString)


	//3
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req, err = http.NewRequestWithContext(ctx, "GET", addrs + "hard-op", nil)
	if (err != nil) {
		log.Fatal(err)
	}
	res, err = client.Do(req)
	if (err != nil) {
		log.Println(false)
		log.Fatal(err)
	}
	defer func() {
		err = res.Body.Close()
		if (err != nil) {
			log.Fatal(err)
		}	
	}()
	body, err = io.ReadAll(res.Body)
	if (err != nil) {
		log.Fatal(err)
	}
	log.Println(true, string(res.Status))

}