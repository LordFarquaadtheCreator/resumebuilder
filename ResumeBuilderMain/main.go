// just using it for testing purposes at the moment
package main

import (
	"fmt"
	"resumebuilder/JobScrapper"
)

// issue!!!! most sites will block a http request, will need manual copy and pasted job descriptions

func main () {
	fmt.Println("Hello World")
	resp, err := JobScrapper.GetUrlText("https://www.ziprecruiter.com/c/Alexander-Interactive/Job/Web-Developer-Back-End-Development/-in-New-York,NY?jid=e3e9a99713484421")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}