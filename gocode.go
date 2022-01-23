package main


/*
// v3
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fileContent := ""
	data, err := ioutil.ReadFile("/root/test1.txt")
	if err != nil {
		fileContent = "read filed"
	}
	fileContent = string(data)

	w.Write([]byte(fileContent))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":3500", r))
}
*/






/*
// v1
package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3500", nil)
}
*/







/*
// v2
package main

import (
	"fmt"
	"net/http"
)

func Start(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<doctype html> <html> <head> <title>Page Index</title> </head> <body> <p> <a href="/Welcome">Welcome</a> | <a href="/Message">Message</a> </p> </body> </html>`
	fmt.Fprintln(w, html)

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//
	//fileContent := ""
	//data, err := ioutil.ReadFile("/root/test1.txt")
	//if err != nil {
	//	fileContent = "read filed"
	//}
	//fileContent = string(data)
	//
	//w.Write([]byte(fileContent))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<doctype html> <html> <head> <title>Page Index</title> </head> <body> <p> <a href="/Welcome">Welcome</a> | <a href="/Message">Message</a> </p> </body> </html>`
	fmt.Fprintln(w, html)
}

func Message(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<doctype html> <html> <head> <title>Page Message</title> </head> <body> <p> <a href="/Index">Index</a> | <a href="/Welcome">Welcome</a> </p> </body> </html>`
	fmt.Fprintln(w, html)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<doctype html> <html> <head> <title>Page Welcome</title> </head> <body> <p> <a href="/Index">Index</a> | <a href="/Message">Message</a> </p> </body> </html>`
	fmt.Fprintln(w, html)
}

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/Index", Index)
	http.HandleFunc("/Message", Message)
	http.HandleFunc("/Welcome", Welcome)
	err := http.ListenAndServe(":3500", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
*/
