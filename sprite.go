package clispin

import "time"

var Sprites = map[int][]string{
	0:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	1:  {"←", "↑", "→", "↓"},
	2:  {"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
	3:  {"◢", "◣", "◤", "◥"},
	4:  {"◰", "◳", "◲", "◱"},
	5:  {"◴", "◷", "◶", "◵"},
	6:  {"◐", "◓", "◑", "◒"},
	7:  {"🌍", "🌎", "🌏"},
	8:  {"⬒", "⬔", "⬓", "⬕"},
	9:  {"⬖", "⬘", "⬗", "⬙"},
	10: {"◡◡", "⊙⊙", "◠◠"},
	11: {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	12: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
}

type Sprite struct {
	sprite         []string
	curIndex       int
	updateInterval int64 // update interval in ns
}

func NewSprite(sprite []string) *Sprite {
	return &Sprite{
		sprite:         sprite,
		curIndex:       0,
		updateInterval: (time.Millisecond * 100).Nanoseconds(),
	}
}

func (s *Sprite) current() string {
	return s.sprite[s.curIndex]
}

func (s *Sprite) next() {
	if s.curIndex < len(s.sprite)-1 {
		s.curIndex++
	} else {
		s.curIndex = 0
	}
}
