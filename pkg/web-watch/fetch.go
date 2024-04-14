package webwatch

import (
	"fmt"
	"io"
	"net/http"
)

func Fetch() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	fmt.Printf("Response body: %v", string(body))
}
