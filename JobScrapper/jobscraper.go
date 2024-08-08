// some package description
package JobScrapper

import (
	"errors"
	"fmt"
	"io"
	"net/http" // https://pkg.go.dev/net/http
	"os"
	"strings"
)

// fetches the content of a url and returns the body as a string
func GetUrlText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("Failed to fetch url")
	}
	
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Failed to read body of fetched url")
	}

	return string(body), nil
}

// checks the frequency of the word in body
func countFreq(body string, word string) int {
	return strings.Count(strings.ToLower(body), strings.ToLower((word)))
}

func main() {
	url := "https://example.com/"
	resp, err := GetUrlText(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(countFreq(resp, KeyWords[0]))
}
