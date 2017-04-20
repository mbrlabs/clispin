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
	frames       []string
	currentFrame int

	lastUpdate int64
	interval   int64
}

func NewSprite(frames []string) *Sprite {
	return &Sprite{
		frames:       frames,
		currentFrame: 0,
		lastUpdate:   -1,
		interval:     (time.Millisecond * 100).Nanoseconds(),
	}
}

func (s *Sprite) Frame() string {
	return s.frames[s.currentFrame]
}

func (s *Sprite) Update() bool {
	if (time.Now().UnixNano()-s.lastUpdate) >= s.interval || s.lastUpdate < 0 {
		s.lastUpdate = time.Now().UnixNano()
		if s.currentFrame < len(s.frames)-1 {
			s.currentFrame++
		} else {
			s.currentFrame = 0
		}

		return true
	}

	return false
}
