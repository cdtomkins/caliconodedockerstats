package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
	//	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Grab from the API target details from the env
	the_attr := os.Getenv("CALICONODEDOCKERSTATS_ATTR_NAME")
	the_target := os.Getenv("CALICONODEDOCKERSTATS_TARGET_NAME")

	// Get the response from the API and log any error
	resp, err := http.Get(the_target)
	if err != nil {
		log.Fatalln(err)
	}

	// Read the response body and log any error
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the body and target attribute to a string and print it
	string_body := string(body)
	the_result := gjson.Get(string_body, the_attr)
	fmt.Println(string(body))
	fmt.Println(the_result.Str)

	//	http.Handle("/metrics", promhttp.Handler())
	//	http.ListenAndServe(":2112", nil)
}
