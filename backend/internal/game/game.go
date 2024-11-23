package game

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/notnil/chess"
)

type Game struct {
	ChessGame *chess.Game
	Mutex     sync.Mutex
	Players   map[string]*websocket.Conn
}

func NewGame() *Game {
	return &Game{
		ChessGame: chess.NewGame(),
	}
}

func (g *Game) ValidateAndMakeMove(move string) (bool, error) {
	err := g.ChessGame.MoveStr(move)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (g *Game) BroadcastMove(move string) error {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	for role, conn := range g.Players {
		if err := conn.WriteJSON(move); err != nil {
			log.Println("Error broadcasting move to", role, ":", err)

			// Close the stale connection
			conn.Close() 
			delete(g.Players, role)

			return err
		}
	}

	return nil
}
