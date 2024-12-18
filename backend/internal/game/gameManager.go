package game

import "sync"

type GameManager struct {
	Games map[uint]*Game
	Mutex sync.Mutex 
}

func NewGameManager() *GameManager {
	return &GameManager{
		Games: make(map[uint]*Game),
	}
}

func (gm *GameManager) GetGame(id uint) *Game{
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	if game, exists := gm.Games[id]; exists {
		return game
	}

	return nil

}

func (gm *GameManager) CreateGame(id uint) *Game{
	gm.Mutex.Lock()
	defer gm.Mutex.Unlock()

	if _, exists := gm.Games[id]; exists {
		return nil
	}
	gm.Games[id] = NewGame()
	return gm.Games[id]

}

