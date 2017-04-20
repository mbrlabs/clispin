package clispin

import "time"

var SpriteFrames = map[int][]string{
	0:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	1:  {"←", "↑", "→", "↓"},
	2:  {"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
	3:  {"◢", "◣", "◤", "◥"},
	4:  {"◰", "◳", "◲", "◱"},
	5:  {"◴", "◷", "◶", "◵"},
	6:  {"◐", "◓", "◑", "◒"},
	7:  {"⬒", "⬔", "⬓", "⬕"},
	8:  {"⬖", "⬘", "⬗", "⬙"},
	9:  {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	10: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
}

type Sprite struct {
	frames         []string
	curIndex       int
	updateInterval int64 // update interval in ns
}

func NewSprite(frames []string) *Sprite {
	return &Sprite{
		frames:         frames,
		curIndex:       0,
		updateInterval: (time.Millisecond * 100).Nanoseconds(),
	}
}

func (s *Sprite) current() string {
	return s.frames[s.curIndex]
}

func (s *Sprite) next() {
	if s.curIndex < len(s.frames)-1 {
		s.curIndex++
	} else {
		s.curIndex = 0
	}
}
