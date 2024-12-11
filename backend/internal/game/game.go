package game

import (
	"sync"

	"github.com/lesismal/nbio/nbhttp/websocket"
	"github.com/notnil/chess"
)

type GameUser struct {
	Socket *websocket.Conn
	Name   string
	Role   string
	UserId uint
}
type Game struct {
	ChessGame *chess.Game
	Mutex     sync.Mutex
	Players   map[uint]*GameUser
	GameId    uint
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



func (g *Game) GetValidMoves() ([]*chess.Move , error) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	moves := g.ChessGame.ValidMoves()

	return moves,nil
}
