package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

// arg3 -> if you render this data as json set name of json as provied below
// type AppStatus struct {
// 	Status      string `json:"status"`
// 	Environment string `json:"environment"`
// 	Version     string `json:"version"`
// }

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// read value from terminal
	//	* property of object to assign from
	//	* name of flag from the command line
	//	* default value if nil was provided
	//	* description of flag
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development | production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	// fmt.Println("Running...")

	// // declare get request returning "status" on page and routing this
	// // into resource /status
	// http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
	// 	currentStatus := AppStatus{
	// 		Status:      "Available",
	// 		Environment: cfg.env,
	// 		Version:     version,
	// 	}
	// 	// arg2 -> prefix, arg3 -> intends
	// 	js, error := json.MarshalIndent(currentStatus, "", "\t")
	// 	if error != nil {
	// 		log.Println(error)
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(js)

	// 	// write response -> write "status" on page
	// 	//fmt.Fprint(w, "status")
	// })

	// error := http.ListenAndServe(fmt.Sprintf(":%d", cfg.port), nil)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting on port", cfg.port)

	error := server.ListenAndServe()

	if error != nil {
		log.Println(error)
	}
}
