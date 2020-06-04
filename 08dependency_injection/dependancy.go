package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "Elodie")

	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

func Greet(w io.Writer, name string) (n int, err error) {
	return fmt.Fprintf(w, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
