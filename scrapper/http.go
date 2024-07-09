package main

import {
	"fmt"
	"http"
}
// https://pkg.go.dev/net/http

func getHttp(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body
}

func main() {
	
}
