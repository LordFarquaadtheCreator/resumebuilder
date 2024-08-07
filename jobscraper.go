package main

import (
	"fmt"
	"net/http"
	"os"
)

// https://pkg.go.dev/net/http

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp.Status)
}
