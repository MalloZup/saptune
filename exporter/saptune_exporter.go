package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", landing)
	http.Handle("/metrics", promhttp.Handler())

	log.Infof("Serving metrics on port 9758")
	log.Fatal(http.ListenAndServe(":9758", nil))
}

func landing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<html>
<head>
	<title>SUSE  Saptune Exporter</title>
</head>
<body>
	<h1>SUSE Saptune expoter</h1>
	<h2>Prometheus exporter for Saptune</h2>
	<ul>
		<li><a href="metrics">Metrics</a></li>
		<li><a href="https://github.com/SUSE/saptune" target="_blank">GitHub</a></li>
	</ul>
</body>
</html>
`))
}
