package main
import (
    "html/template"
    "net/http"
    "fmt"
)

const (
    TEMPLATE_DIR = "../html/"
)

func index(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles(TEMPLATE_DIR + "index.html")
    if err != nil{
        fmt.Println(err.Error())
    }
    s := map[string]string{
        "name": "sd",
        "fuck": "fuck",
    }
    t.Execute(w, s)
}

func main(){
    http.HandleFunc("/", index)
    fmt.Println("http start ! listen on 0.0.0.0:9999")
    http.ListenAndServe(":9999", nil)
}
