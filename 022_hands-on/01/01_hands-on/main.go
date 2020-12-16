package main

import (
	"io"
	"net/http"
)

func m(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Me route me route")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog route dog route")
}

func main() {
	http.HandleFunc("/me", m)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/", m)

	http.ListenAndServe(":8080", nil)

}
