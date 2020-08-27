package main

import (
	"./constants"
	"./handlers"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)
	// Format of command line arguments: NYTAPI NEWSAPI TODO: LOGLVL
	// Cmmnd Line Flags to perform particular actions
	flag.StringVar(&constants.NytKeyPtr, "nyt", "",
		"nyt api key")

	flag.StringVar(&constants.NewsKeyPtr, "news", "",
		"NewsAPI key")

	tail := flag.Args()
	flag.Parse()

	log.Infof("NYT Key: %v \n News Key: %v \n",
		constants.NytKeyPtr, constants.NewsKeyPtr)
	fmt.Printf("these are the trailing arguments: %v\n", tail)
	if constants.NytKeyPtr == "" || constants.NewsKeyPtr == "" {
		log.Fatal("Do not have necessary API keys")
		os.Exit(1)
	}

	// TODO: create log file
	log.SetOutput(os.Stdout)
	log.Info("Starting program...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()

	// Declare the static file directory
	fs := http.FileServer(http.Dir("frontendassets"))
	mux.Handle("/frontendassets/", http.StripPrefix("/frontendassets/", fs))

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/search", handlers.SearchHandler)
	log.Info("Listening on port ", port)
	http.ListenAndServe(":"+port, mux)

}
