package main

import (
	"fmt"
	"net/http"
        "strconv"
        "strings"
)

func sumHandler(w http.ResponseWriter, r *http.Request) {
     xs := strings.Split(r.URL.Path[len("/sum/"):], ",")
     sum := 0.0
     for i := 0; i < len(xs); i++ {         
         if x, err := strconv.ParseFloat(xs[i], 64); err == nil {
            sum = sum + x
         }
     }
     fmt.Fprintln(w, fmt.Sprintf("%f", sum))
     //x := [5]float {0.0, 0.25, 0.5, 0.75, 1.0}
     //y := [5]float {0.0, 1.0, 0.0, 1.0, 0.0}
     //for i := 0; i < 5; i++ {
     //	if
     //fmt.Fprintln(w, fmt.Sprintf("%f", x + 1))
}

func main() {
	http.HandleFunc("/sum/", sumHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
