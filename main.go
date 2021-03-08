package main

import(
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func postJson(url string, json []byte) (error) {
    fmt.Println("url:", url)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    return err
}

func defaultHandler(writer http.ResponseWriter, request *http.Request) {

    var url = os.Getenv("WEBHOOK_URL")
    var json = []byte(`{"text":"Hi"}`)

    err := postJson(url, json)
    if err != nil {
        fmt.Fprintf(writer, "Error: %s", err.Error())
        return
    }

    fmt.Fprintf(writer, "OK")
}

func main() {
    var server = fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT"))
    fmt.Printf("Server listening on %s", server)
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(server, nil)
}
