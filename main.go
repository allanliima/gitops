package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello ArgoCD!!"))
	})

	http.ListenAndServe(":8081", nil)
}
