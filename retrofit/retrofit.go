// test.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{user_name=\"%s\", age=%s}", r.FormValue("user_name"), r.FormValue("age"))
}
