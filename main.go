package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("./view/index.html")
		if err != nil {
			http.Error(w, "Could not open requested file", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", content)
	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
