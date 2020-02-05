package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
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
	log.Println("load data from file")
	data, err := populateData()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/api/sorts", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		sortNumber(data)
		since := time.Since(now)
		res := fmt.Sprintf("finished sorting in %f second\n", since.Seconds())

		w.Write([]byte(res))
	})

	port := getPort()

	log.Printf("start http server on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
