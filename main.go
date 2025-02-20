package main

import (
	"fmt"
	"net/http"
)

func main() {
	tasks := []string{}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if r.Method == "GET" {
			for _, i := range tasks {
				if i == id {
					fmt.Fprintf(w, "task with id=%v found\n", id)
					return
				}
			}
			fmt.Fprintf(w, "task with id=%v not found\n", id)
		} else {
			tasks = append(tasks, id)
			fmt.Fprintf(w, "set task with id=%v\n", id)
		}
	})

	http.ListenAndServe("localhost:3000", mux)
}
