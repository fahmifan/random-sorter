package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"rebalance-test/utils"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var (
	debug bool
	port  int
)

func init() {
	flag.BoolVar(&debug, "debug", false, "--debug true (default to false)")
	flag.IntVar(&port, "port", 8081, "--port 8082 (default 8081)")
}

func main() {
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fpath := path.Join(dir, "data.out")
	log.Printf("load file from: %s\n", fpath)

	data, err := populateData(fpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("finish populate data")

	ip, err := utils.PrivateIP()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/api/sorts", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		datac := make([]int, len(data))
		copy(datac, data)
		sortNumber(datac)
		since := time.Since(now)
		res := fmt.Sprintf("finished sorting in %.3f second from %s:%d\n", since.Seconds(), ip, port)
		time.Sleep(1 * time.Second) // slow down
		if debug {
			log.Println(res)
		}
		w.Write([]byte(res))
	})

	log.Printf("start http server on at %s:%d", ip, port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), r))
}
