package main

import (
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
	// in nanosecond
	sleep int = 0
)

// port in format `:portNumber`
func getPort() string {
	port := ":8000"
	if val, ok := os.LookupEnv("PORT"); ok {
		p := stringToInt(val)
		if p >= 3000 {
			port = fmt.Sprintf(":%d", p)
		}
	}

	return port
}

func main() {
	if val, ok := os.LookupEnv("SLEEP"); ok {
		sleep = stringToInt(val)
	}

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

	port := getPort()
	ip, err := utils.PrivateIP()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
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
		sortNumber(data)
		since := time.Since(now)
		res := fmt.Sprintf("finished sorting in %.3f second from %s%s\n", since.Seconds(), ip, port)

		w.Write([]byte(res))
	})

	log.Printf("start http server on at %s%s", ip, port)
	log.Fatal(http.ListenAndServe(port, r))
}
