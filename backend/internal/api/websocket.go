package api

import (
    "github.com/uanik03/uanik-chess/internal/game"
    "github.com/uanik03/uanik-chess/internal/models"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketServer struct {
    GameManager *game.GameManager
}

func NewWebSocketServer() *WebSocketServer {
    return &WebSocketServer{
        GameManager: game.NewGameManager(),
    }
}

func (wss *WebSocketServer) HandleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer conn.Close()

    // Example Game ID and Player Role
    gameID := "game-1" // For simplicity, this can be dynamic in the future.
    // playerRole := r.URL.Query().Get("role") // "white" or "black"
    game := wss.GameManager.GetOrCreateGame(gameID)

    for {
        var msg models.Message
        if err := conn.ReadJSON(&msg); err != nil {
            log.Println("Error reading message:", err)
            break
        }

        if msg.Type == "move" {
            valid, err := game.ValidateAndMakeMove(msg.Data)
            response := models.MoveResponse{Valid: valid}
            if err != nil {
                response.Error = err.Error()
            }

            if err := conn.WriteJSON(response); err != nil {
                log.Println("Error sending response:", err)
                break
            }
        }
    }
}
