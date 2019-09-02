package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"

	"github.com/andreluzz/gfg-search-service/service/routes"
	"github.com/andreluzz/gfg-search-service/service/storage"
)

func main() {
	esHost := os.Getenv("ES_HOST")
	if esHost == "" {
		esHost = "http://localhost:9200"
	}

	r := routes.New(esHost, storage.SearchIndex)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  25 * time.Second,
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		log.Println("Search service started on port 8080")
		if err := server.ListenAndServe(); err != nil {
			if strings.Contains(err.Error(), "bind: address already in use") {
				log.Fatalf("port 8080 already in use\n")
			}
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// execute dataset.sh to index documents to elasticsearch for testing purpose
	// retry every 5 seconds to wait elasticsearch to be ready
	retry := 0
	go func() {
		for {
			cmd := exec.Command("./dataset.sh", esHost)
			err := cmd.Run()
			if err == nil || retry > 6 {
				log.Printf("elasticsearch demo documents indexed")
				break
			}
			if err.Error() == "exit status 7" {
				retry++
				log.Printf("elasticsearch not available. retry: %d", retry)
				time.Sleep(time.Second * 10)
			} else {
				log.Printf("invalid dataset creation script. host: %s | %v", esHost, err)
				break
			}
		}
	}()

	<-stopChan
	fmt.Println("")
	log.Println("Shutting down service...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
	defer cancel()
	log.Println("Service stopped!")
}
