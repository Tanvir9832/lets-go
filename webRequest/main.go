package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println("Got a error", err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
