package game

import (
    "github.com/notnil/chess"
)

type Game struct {
    ChessGame *chess.Game
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
