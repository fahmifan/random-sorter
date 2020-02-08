package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rebalance-test/utils"
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

	port := getPort()
	ip, err := utils.PrivateIP()
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
		res := fmt.Sprintf("finished sorting in %.3f second from %s%s\n", since.Seconds(), ip, port)

		w.Write([]byte(res))
	})

	log.Printf("start http server on at %s%s", ip, port)
	log.Fatal(http.ListenAndServe(port, r))
}
