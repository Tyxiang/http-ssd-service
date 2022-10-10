package main
 
import (
    "fmt"
    "net/http"
)
 
type server int
 
func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    w.Write([]byte("Hello World!"))
}
 
func main() {
    var s server
    http.ListenAndServe("localhost:8000", &s)
}