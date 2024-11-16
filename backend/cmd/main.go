package main

import (
    "github.com/uanik03/uanik-chess/internal/api"
    "log"
    "net/http"
)

func main() {
    wss := api.NewWebSocketServer()

    http.HandleFunc("/ws", wss.HandleConnection)

    log.Println("WebSocket server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
