package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct {
}

func main() {
	app := Config{}
	log.Printf("Starting broker-service on port %s\n", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
		return
	}
}

// func logRequestHandler(h http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println((*r).Header)
// 		fmt.Println(w)
// 	}
// 	return http.HandlerFunc(fn)
// }
