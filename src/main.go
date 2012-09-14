package main

import (
	"net/http"

	"engine/go.net/websocket"
	. "engine/mango"
	"handlers"
	"wscon"
)

func main() {
	l, r := handlers.LayoutAndRenderer()
	s := new(Stack)
	s.Middleware(l, r)

	http.Handle("/chat", websocket.Handler(wscon.BuildConnection))
	http.HandleFunc("/join", s.HandlerFunc(handlers.Join))
	http.HandleFunc("/", s.HandlerFunc(handlers.Home))
	http.HandleFunc("/public/", assetsHandler)

	go wscon.InitChatRoom()

	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[len("/"):])
}
