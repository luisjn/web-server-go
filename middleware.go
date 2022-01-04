package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, req *http.Request) {
			flag := true
			fmt.Println("Checking authentication...")
			if flag {
				hf(rw, req)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(req.URL, time.Since(start))
			}()
			fmt.Println("Checking authentication...")
			hf(rw, req)
		}
	}
}
