package main

import (
   "io"
   "fmt"
   "github.com/tidwall/gjson"
   "log"
   "net/http"
//   "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
//   http.Handle("/metrics", promhttp.Handler())
//   http.ListenAndServe(":2112", nil)
   resp, err := http.Get("https://hub.docker.com/v2/repositories/calico/node/")
   if err != nil {
      log.Fatalln(err)
   }

//We Read the response body on the line below.
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }

//Convert the body to type string
   fmt.Println(string(body))
   sb := string(body)
   theresult := gjson.Get(sb,"pull_count")
   log.Printf(theresult.Str)
   log.Printf(sb)
}
