// just using it for testing purposes at the moment
package main

import (
	"fmt"
	"resumebuilder/JobScrapper"
)

func main () {
	fmt.Println("Hello World")
	resp, err := JobScrapper.GetUrlText("https://example.com/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}