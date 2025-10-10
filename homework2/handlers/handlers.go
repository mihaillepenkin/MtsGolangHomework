package handlers

import (
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

const semVerVersion string = "1.2.3"

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(semVerVersion))
}

type decodeResponse struct {
	OutputString string `json:"outputString"`
}
type decodeInput struct {
	InputString string `json:"inputString"`
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input decodeInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	outputString, err := decodeToString(input.InputString)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(decodeResponse{OutputString: outputString})
}

func decodeToString(inputString string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(inputString)
	if (err != nil) {
		return "", err
	}
	outputString := string(decodedBytes)
	return outputString, nil
}

func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	timeOfWork := rand.Intn(11) + 10
	time.Sleep(time.Duration(timeOfWork) * time.Second)
	switch rand.Intn(6) {
	case 0:
		w.WriteHeader(200)
	case 1:
		w.WriteHeader(201)	
	case 2:
		w.WriteHeader(202)	
	case 3:
		w.WriteHeader(500)
	case 4:
		w.WriteHeader(501)	
	case 5:
		w.WriteHeader(502)	
	}
}



