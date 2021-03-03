package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	cmd := newCollector()


	prometheus.MustRegister(cmd)
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))


}