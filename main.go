package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port int
	var directory string

	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.IntVar(&port, "port", 8080, "Port to start the server on")
	fs.StringVar(&directory, "dir", ".", "Directory to serve. Defaults to the current directory.")
	fs.Parse(os.Args[1:])

	fmt.Printf("Listening on %d, serving %s", port, directory)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(directory))))
}
