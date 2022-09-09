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
	
	// key := "Wuh1FF+QADZQ4z5dFFffc+Vl1WsTvvhzTURY2Bmg+jsw3FwDNDMY52MdBfWHzb_yJez5fmk+TWh6mwTwWPfyAA=="
	
	// req, err := http.NewRequest("POST", "https://ft-nostress.chili-publish.online/rest-api/v1.2/resources/documents/c6018b06-47a3-448c-b180-194bde802f96/info?extended=false", nil)

		client := &http.Client{}

	   	postBody, _ := json.Marshal(map[string]string{
      		"id":  "c6018b06-47a3-448c-b180-194bde802f96",
   		})

   		responseBody := bytes.NewBuffer(postBody)

		//Leverage Go's HTTP Post function to make request
   		req, err := http.NewRequest("https://ft-nostress.chili-publish.online/rest-api/v1.2/resources/documents/c6018b06-47a3-448c-b180-194bde802f96/info?extended=false", "application/json", responseBody)
		if(err != nil){
			fmt.Println("pain")
		}

		req.Header.Add("api-key", "Wuh1FF+QADZQ4z5dFFffc+Vl1WsTvvhzTURY2Bmg+jsw3FwDNDMY52MdBfWHzb_yJez5fmk+TWh6mwTwWPfyAA==")
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