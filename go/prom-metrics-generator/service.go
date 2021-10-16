package main

import (
        "tmomon/logger"
        "tmomon/nifi"
        "tmomon/redis"
        "flag"
	"os"
        "net/http"
        "time"
        "strings"
)

func infiniteLoop(
        sleepSec time.Duration,
){
        // Accept the command line flags
        nifiUrl := flag.String("nifiUrl", "none", "Nifi URL to use")
        redisUrl := flag.String("redisUrl", "none", "Redis URL to use")
        redisTmoApps := flag.String("redisTmoApps", "none", "Comma seperated list of TMO apps. Ex, auditor,preflighter")
        flag.Parse()

        var nifiEnabled, redisEnabled bool = false, false


        if *nifiUrl != "none" {
                clog.Info.Println("Nifi url to use is", *nifiUrl)
                nifiEnabled = true
        }

        if *redisUrl != "none" && *redisTmoApps != "none" && *redisTmoApps != "," {
                clog.Info.Println("Redis url to use is", *redisUrl)
                redisEnabled = true
        }

        for {
                // NIFI metrics
                if nifiEnabled {
                        nifiClusterUrl := nifi.NifiFlowAbout(*nifiUrl)
                        nifiFlowPgRoot := nifi.NifiFlowPgRoot(*nifiUrl , nifiClusterUrl)
                        nifi.NifiFlowStatus(*nifiUrl , nifiClusterUrl, nifiFlowPgRoot)
                        nifi.NifiFlowClusterSummary(*nifiUrl, nifiClusterUrl)
                }


                // Redis metrics
                if redisEnabled {
                        redis.RedisConn(*redisUrl)
                        redisTmoAppsSliced := strings.Split(*redisTmoApps, ",")
                        for _, app := range redisTmoAppsSliced {
                                appStream := app+"-stream"
                                appGroup := app+"-group"
        
                                if exists := redis.RedisIfExists(appStream); exists == 1 {
                                        redis.RedisStreamLength(*redisUrl,appStream)
                                        redis.RedisStreamPending(*redisUrl,appStream,appGroup)
                                }
                        }
                }

		// Calling Sleep method
		time.Sleep(sleepSec * time.Second)
        }
}

func main() {
        // Initialize logger
        clog.Logger(os.Stdout, os.Stderr)

        // Run the infinitr function in Go routine
        go infiniteLoop(20)
        
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, "/tmp/metrics/metrics")
        })
        http.ListenAndServe(":2112", nil)
}
