package api

import (
	"chess-backend/config"
	"chess-backend/internal/game"
	"chess-backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/lesismal/nbio/nbhttp/websocket"
)

var upgrader = websocket.NewUpgrader()

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

func (ws *WebSocketServer) HandleWebsocketConnection(c *gin.Context) {

	w := c.Writer
	r := c.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	log.Println("OnOpen:", conn.RemoteAddr().String())

	var requestBody struct {
		Join   bool `json:"join"`
		GameID uint `json:"gameId"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("line 50: websocket.go", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid request body"))
		conn.Close()
		return
	}

	log.Printf("Join=%t, GameID=%d", requestBody.Join, requestBody.GameID)
	userId, exists := c.Get("userId")
	userName, exists := c.Get("userName")

	if !exists {
		log.Println("err-2")
		conn.WriteMessage(websocket.TextMessage, []byte("Authentication required"))
		conn.Close()
		return
	}

	if requestBody.Join {
		currentGame := ws.GameManager.GetGame(requestBody.GameID)
		if currentGame == nil {
			log.Println("err-3 ")
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid game"))
			conn.Close()
			return
		}
		if len(currentGame.Players) == 2 {
			log.Println("err-4")
			conn.WriteMessage(websocket.TextMessage, []byte("Game is full"))
			conn.Close()
			return
		}

		currentGame.Players[userId.(uint)] = &game.GameUser{
			Socket: conn,
			Name:   userName.(string),
			Role:   "b",
			UserId: userId.(uint),
		}

		conn.WriteMessage(websocket.TextMessage, []byte("joined successfully"))

	} else {
		newGame := models.Game{
			PlayerWhiteID: userId.(uint),
			Status:        "pending", 
			Result:        "nil",  
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}


	if err := config.DB.Create(&newGame).Error; err != nil {
		log.Println("err-5", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}
		currentGame := ws.GameManager.CreateGame(newGame.ID)
		currentGame.GameId = newGame.ID
		currentGame.Players[userId.(uint)] = &game.GameUser{
			Socket: conn,
			Name:   userName.(string),
			Role:   "w",
			UserId: userId.(uint),
		}

		response := map[string]interface{}{
			"message": "joined successfully",
			"gameId":  newGame.ID,
		}
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Println("err-6:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("An error occurred"))
			conn.Close()
			return
		}
	
		// Send the response message to the user
		conn.WriteMessage(websocket.TextMessage, responseBytes)

	}

}

func (ws *WebSocketServer) BroadcastValidMoves(gameId uint, userId uint) error {

	currentGame, exists := ws.GameManager.Games[gameId]
	if !exists {
		return fmt.Errorf("game with ID %d not found", gameId)
	}

	//get all the valid moves of a particular game instance through the function GetValidMoves I Wrote in game struct
	moves, err := currentGame.GetValidMoves()
	if err != nil {
		return err
	}

	user, exists := currentGame.Players[userId]
	if !exists {
		return fmt.Errorf("user with ID %d not found in game", userId)
	}

	moveData, err := json.Marshal(moves)
	if err != nil {
		log.Println("err-7:", err)
		return fmt.Errorf("failed to serialize moves: %v", err)
	}

	if err := user.Socket.WriteMessage(websocket.TextMessage, moveData); err != nil {
		log.Println("err-8:", err)

		// 	// Close the stale connection
		user.Socket.Close()
		delete(currentGame.Players, userId)

		return err
	}

	return nil
}

func (ws *WebSocketServer) BroadcastMove(move string, gameId uint) error {

	// currentGame := ws.GameManager.Games[gameId]
	// for id, user := range currentGame.Players {
	// 	if err := user.Socket.WriteJSON(move); err != nil {
	// 		log.Println("Error broadcasting move to", user.Name, ":", err)

	// 		// Close the stale connection
	// 		user.Socket.Close()
	// 		delete(currentGame.Players, id)

	// 		return err
	// 	}
	// }

	return nil
}
