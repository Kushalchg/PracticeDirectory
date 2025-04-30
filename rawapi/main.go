package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

}

func main() {

	http.HandleFunc("/list", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "for bimohit")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil))

	fmt.Println("server started successfully")
	fmt.Printf("Listening on port %v \n ", os.Getenv("PORT"))

}
