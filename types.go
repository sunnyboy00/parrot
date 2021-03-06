package main

type Hand int

const (
	Up Hand = iota
	Right
	Down
	Left
	Quit
)

type Player interface {
	NextHand(func() uint32) Hand
	SetState(GameState)
}

type GameState struct {
	Grid       [4][4]int
	Won        bool
	Moved      bool
	Over       bool
	Score      int
	Points     int
	Zen        string
	Session_id string
}

type Simulator interface {
	Move(Hand) bool
	AddRandomCell() bool
	GetAvailableCells() int
	GetMaxTile() int
	Score() int
	GetBoard() [4][4]int
}

func (s *GameState) MaxTile() (ret int) {
	ret = 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			ret = Max(ret, s.Grid[y][x])
		}
	}
	return
}
