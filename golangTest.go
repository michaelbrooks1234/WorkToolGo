package main

import (
   "fmt"
   "net/http"
)

func main() {


	fmt.Println("hello, world.")
	client := &http.Client{
	
	}

	
	req, err := http.NewRequest("POST", "http://example.com", nil)

	fmt.Println(req);
	fmt.Println(err);

	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	resp, err := client.Do(req)

	fmt.Println(resp);
	fmt.Println(err);
}