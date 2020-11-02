package http

import (
	"../client"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	fmt.Println(param)
	client.Send(param)
	//fmt.Fprintf(w, "hello world")
	w.Write([]byte("ok"))
}

//func sendMessage(w http.ResponseWriter, r *http.Request)  {
//	client.Send(r.URL.Query().Get("param"))
//}

func init() {
	fmt.Println("http server init")
	http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/send", sendMessage)
	http.ListenAndServe(":8000", nil)
}
