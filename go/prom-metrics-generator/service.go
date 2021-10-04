package main

import (
        "prom-metrics-generator/logger"
        "prom-metrics-generator/nifi"
        "flag"
	"os"
        "net/http"
        "time"
)

func infiniteLoop(
        url string,
){
        for {
                nifiClusterUrl := nifi.NifiFlowAbout(url)
                nifiFlowPgRoot := nifi.NifiFlowPgRoot(url , nifiClusterUrl)
                nifi.NifiFlowStatus(url , nifiClusterUrl, nifiFlowPgRoot)
                nifi.NifiFlowClusterSummary(url, nifiClusterUrl)

		// Calling Sleep method
		time.Sleep(20 * time.Second)
        }
}

func main() {
        // Initialize logger
        clog.Logger(os.Stdout, os.Stderr)

        // Accept the command line flags
        nifiUrl := flag.String("nifiUrl", "tmo.wdcd01.uswest2.veritone.com", "Nifi URL to use")
        flag.Parse()

        clog.Info.Println("Nifi url to use is", *nifiUrl)

        go infiniteLoop(*nifiUrl)
        
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, "/tmp/metrics/metrics")
        })
        http.ListenAndServe(":2112", nil)
}
