package http

import (
	"../client"
	"fmt"
	"net/http"
)

func conn(w http.ResponseWriter, r *http.Request) {
	client.RequireConn()
	w.Write([]byte("ok"))
}
func disConn(w http.ResponseWriter, r *http.Request) {
	client.DisConn()
	w.Write([]byte("ok"))
}

func init() {
	fmt.Println("http server init")
	http.HandleFunc("/conn", conn)
	http.HandleFunc("/dis", disConn)
	//http.HandleFunc("/send", sendMessage)
	http.ListenAndServe(":8000", nil)
}
