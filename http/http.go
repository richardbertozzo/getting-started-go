package main

import "net/http"

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/hello-world", handlerHelloWorld)

	http.ListenAndServe(":8080", nil)
}
