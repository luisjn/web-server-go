package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello!")
}

func HandleHome(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "home")
}

func PostRequest(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(rw, "error: %v", err)
		return
	}
	fmt.Fprintf(rw, "payload: %v\n", metadata)
}

func UserPostRequest(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(rw, "error: %v", err)
		return
	}
	res, err := user.ToJson()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}
