package main

import (
	"flag"
	"github.com/sonemaro/overlap/internal/api"
	"github.com/sonemaro/overlap/internal/rectangle"
	"log"
	"net/http"
	"strconv"
)

func main() {
	addr := flag.String("addr", "127.0.0.1", "server address")
	port := flag.Int("port", 41270, "server port")
	flag.Parse()

	fullAddr := *addr + ":" + strconv.Itoa(*port)
	log.Printf("Listening on %s", fullAddr)
	router := api.SetupRouter()
	rectangle.DB = rectangle.NewStorage()
	log.Fatal(http.ListenAndServe(fullAddr, router))
}
