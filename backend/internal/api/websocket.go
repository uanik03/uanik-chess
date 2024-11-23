package api

import (
	"chess-backend/internal/game"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin: func(r *gin.Context) bool { return true },
   }

type WebSocketServer struct {
    GameManager *game.GameManager
}

func NewWebSocketServer() *WebSocketServer {
    return &WebSocketServer{
        GameManager: game.NewGameManager(),
    }
}

func (ws *WebSocketServer) HandleWebsocketConnection(c *gin.Context){

}