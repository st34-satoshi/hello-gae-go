package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World in GAE" )
}

func main() {
  http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
  port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
