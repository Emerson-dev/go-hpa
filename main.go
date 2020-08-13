package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := 0.0001
		for i := 0; i <= 1000000; i++ {
			x = x + math.Sqrt(x)
		}

		fmt.Fprintf(w, "Code.education Rocks!")
		fmt.Printf("Code.education Rocks!")

	})
	http.ListenAndServe(":8000", nil)
}
