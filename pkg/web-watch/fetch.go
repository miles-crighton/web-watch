package webwatch

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Fetch() {
	resp, err := http.Get("http://google.com")
	check(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	check(err)
	fmt.Printf("Response body: %v", string(body))

	err = os.MkdirAll("page-store", 0750)
	check(err)
	file, err := os.Create("page-store/google.com")
	check(err)
	defer file.Close()

	file.Write(body)

	fmt.Println("Succesfully saved file")
}
