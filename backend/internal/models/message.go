package models

type Message struct {
    Type string `json:"type"` // e.g., "move" or "join"
    Data string `json:"data"` // e.g., "e2e4" for a move
    Player string `json:"player"` // "white" or "black"
}

type MoveResponse struct {
    Valid bool   `json:"valid"`
    Error string `json:"error,omitempty"`
}
