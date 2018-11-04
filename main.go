package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	)
type Response struct{
	Date string  `json:"date"`
	Base string `json:"base"`
	TimeStamps int `json:"timestamp"`
	Rates interface{} `json:"rates"` 
}

func main() {
	response, err := http.Get("https://bankersalgo.com/apirates2/<apikey>/<secretkey>/EUR")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var responseObject Response
	json.Unmarshal(data, &responseObject)

	 // fmt.Println(responseObject.Rates)
	for key, value := range responseObject.Rates.(map[string]interface{}) {
	 fmt.Println(key,value)

	}

	}



}
