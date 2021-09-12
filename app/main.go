package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/app/models"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
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
	models models.Models
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
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://root:P%40ssw0rd@localhost:15432/go_movies?sslmode=disable", "postgres connection string")
	//flag.StringVar(&cfg.db.dsn, "dsn", "host=localhost port=5432 user=root password= dbname=%s sslmode=disable", "postgres connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, error := openDB(cfg)
	if error != nil {
		logger.Fatal(error)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
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

	error = server.ListenAndServe()

	if error != nil {
		log.Println(error)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, error := sql.Open("postgres", cfg.db.dsn)
	if error != nil {
		return nil, error
	}

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	error = db.PingContext(context)
	if error != nil {
		return nil, error
	}
	return db, nil
}
