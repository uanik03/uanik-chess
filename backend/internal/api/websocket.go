package api

import (
	"chess-backend/internal/game"
	"fmt"
	"log"
	"net/http"

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
type WsContent struct {
	event string
}
func (ws *WebSocketServer) HandleWebsocketConnection(c *gin.Context){

	var content WsContent
    if err := c.ShouldBindJSON(&content); err != nil {
        c.Error(err)
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

	if content.event == "create"{

		// ws.GameManager.CreateGame()

	}else{

	}

	// ws.GameManager

}

func (ws *WebSocketServer) BroadcastValidMoves(gameId uint, userId uint) error {

	currentGame := ws.GameManager.Games[gameId]
	
	//get all the valid moves of a particular game instance through the function GetValidMoves I Wrote in game struct
	moves, err := currentGame.GetValidMoves()
	if err != nil {
		return err
	}
	
	user, exists := currentGame.Players[userId]
	if !exists {
		return  fmt.Errorf("user with ID %d not found in game", userId)
	}

	if err := user.Socket.WriteJSON(moves); err != nil {
		log.Println("Error broadcasting move to",user.Name,":", err)

		// Close the stale connection
		user.Socket.Close()
		delete(currentGame.Players, userId)

		return  err
	}

	return nil
}


func (ws *WebSocketServer) BroadcastMove(move string, gameId uint) error {

	currentGame := ws.GameManager.Games[gameId]
	for id, user := range currentGame.Players {
		if err := user.Socket.WriteJSON(move); err != nil {
			log.Println("Error broadcasting move to", user.Name, ":", err)

			// Close the stale connection
			user.Socket.Close()
			delete(currentGame.Players, id)

			return err
		}
	}

	return nil
}