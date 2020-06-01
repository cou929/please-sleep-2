package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("listening localhost:5000 to serve %s", os.Args[1])
	log.Fatal(http.ListenAndServe(":5000", http.FileServer(http.Dir(os.Args[1]))))
}
