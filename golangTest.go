package main

import (
   "bytes"
   "encoding/json"
   "io"
   "log"
   "net/http"
   "fmt"
)

func main() {

	// fmt.Println("hello, world.")
	// client := &http.Client{
	
	// }
	
	// key := "t"
	
	// req, err := http.NewRequest("POST", "", nil)

		client := &http.Client{}

	   	postBody, _ := json.Marshal(map[string]string{
      		"id":  "",
   		})

   		responseBody := bytes.NewBuffer(postBody)

		//Leverage Go's HTTP Post function to make request
   		req, err := http.NewRequest("", "application/json", responseBody)
		if(err != nil){
			fmt.Println("pain")
		}

		resp, err := client.Do(req)
		
		//Handle Error
   		if err != nil {
      		log.Fatalf("An Error Occured %v", err)
   		}

   		defer resp.Body.Close()

		//Read the response body
   		body, err := io.ReadAll(resp.Body)
		
   		if err != nil {
      		log.Fatalln(err)
   		}

   		sb := string(body)
   		fmt.Println(sb)

}