package game

import "sync"

type GameManager struct {
    Games map[string]*Game
    Mutex sync.Mutex
}

func NewGameManager() *GameManager {
    return &GameManager{
        Games: make(map[string]*Game),
    }
}

func (gm *GameManager) GetOrCreateGame(id string) *Game {
    gm.Mutex.Lock()
    defer gm.Mutex.Unlock()

    if game, exists := gm.Games[id]; exists {
        return game
    }
    newGame := NewGame()
    gm.Games[id] = newGame
    return newGame
}
