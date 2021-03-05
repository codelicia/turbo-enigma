package main

import(
       "fmt"
       "net/http"
)

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
       fmt.Fprintf(writer, "OK")
}

func main() {
       http.HandleFunc("/", defaultHandler)
       fmt.Printf("Listening at 0.0.0.0:8080")
       http.ListenAndServe("0.0.0.0:8080", nil)
}
