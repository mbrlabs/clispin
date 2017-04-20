package clispin

import "time"

// SpriteFrames are a premade sets of sprite frames.
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

// Sprite is a unicode sprite implementation.
type Sprite struct {
	Interval int64

	frames       []string
	currentFrame int

	lastUpdate int64
}

// NewSprite creates a new sprite given the specified frames.
// TODO: add update interval param
func NewSprite(frames []string, interval int64) *Sprite {
	return &Sprite{
		frames:       frames,
		currentFrame: 0,
		lastUpdate:   -1,
		Interval:     interval,
	}
}

// Frame returns the current frame
func (s *Sprite) Frame() string {
	return s.frames[s.currentFrame]
}

// Update updates the animation.
// If the function returns true the current frame changed.
func (s *Sprite) Update() bool {
	// handle first update
	if s.lastUpdate == -1 {
		s.lastUpdate = time.Now().UnixNano()
		return false
	}

	delta := time.Now().UnixNano() - s.lastUpdate
	if delta >= s.Interval {
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
