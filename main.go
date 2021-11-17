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

func getResultFromAPI(the_attr string, the_target string) {
	go func() {
		// Forever...
		for {
			// Get the response from the API and log + exit on any error
			resp, err := http.Get(the_target)
			if err != nil {
				log.Fatalln(err)
			}

			// If all okay so far, read the response body and log + exit on any error
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			// If all okay so far, convert the body to a string and grab the attribute
			string_body := string(body)
			the_result := gjson.Get(string_body, the_attr)

			// Update the Prometheus gauge with the float value of the attribute
			calico_node_pull_count_gauge.Set(the_result.Float())

			// Wait 5 minutes and do it all again
			time.Sleep(3 * time.Minute)
		}
	}()
}

var (
	calico_node_pull_count_gauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "caliconodedockerstats_gauge",
		Help: "calico-node's pull_count from Docker Hub",
	})
)

func main() {
	// Grab from the API target details from the env
	env_attr := os.Getenv("CALICONODEDOCKERSTATS_ATTR_NAME")
	env_target := os.Getenv("CALICONODEDOCKERSTATS_TARGET_NAME")

	// Get the API result and update the Prometheus gauge with the value
	getResultFromAPI(env_attr, env_target)

	// Wait 10 seconds to make sure the attribute has updated before starting listener
	time.Sleep(10 * time.Second)

	// Start the Prometheus metrics listener
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9088", nil)
}
