package main

import(
       "fmt"
       "net/http"
       "net/http/httputil"
)

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
       fmt.Fprintf(writer, "OK")

        dump, err := httputil.DumpRequest(request, true)
        if err != nil {
            http.Error(writer, fmt.Sprint(err), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(writer, "%q", dump)
}

func main() {
       http.HandleFunc("/", defaultHandler)
       fmt.Printf("Listening at 0.0.0.0:80")
       http.ListenAndServe("0.0.0.0:80", nil)
}
