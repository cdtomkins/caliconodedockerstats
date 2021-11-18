package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tidwall/gjson"
)

func getResultFromAPI(theAttr string, theTarget string) {
	go func() {
		// Forever...
		for {
			// Get the response from the API and log + exit on any error
			resp, err := http.Get(theTarget)
			if err != nil {
				log.Fatalln(err)
			}

			// If all okay so far, read the response body and log + exit on any error
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			// Close the HTTP response
			resp.Body.Close()

			// If all okay so far, convert the body to a string and grab the attribute
			stringBody := string(body)
			the_result := gjson.Get(stringBody, theAttr)

			// Update the Prometheus gauge with the float value of the attribute
			calicoNodePullCountGauge.Set(the_result.Float())

			// Wait 5 minutes and do it all again
			time.Sleep(3 * time.Minute)
		}
	}()
}

var (
	calicoNodePullCountGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "caliconodedockerstats_gauge",
		Help: "calico-node's pull_count from Docker Hub",
	})
)

func main() {
	// Grab from the API target details from the env
	envAttr := os.Getenv("CALICONODEDOCKERSTATS_ATTR_NAME")
	envTarget := os.Getenv("CALICONODEDOCKERSTATS_TARGET_NAME")

	// Get the API result and update the Prometheus gauge with the value
	getResultFromAPI(envAttr, envTarget)

	// Wait 10 seconds to make sure the attribute has updated before starting listener
	time.Sleep(10 * time.Second)

	// Start the Prometheus metrics listener
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9088", nil)
}
