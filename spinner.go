package clispin

import (
	"fmt"
	"time"

	"github.com/mbrlabs/uilive"
)

// Spinner is an infinite ui spinner with status text
type Spinner struct {
	RefreshInterval time.Duration
	LastFrame       string

	writer *uilive.Writer
	sprite *Sprite

	running bool
	dirty   bool
	text    string
}

// New creates a new Spinner with the given sprite.
// If sprite nil, a defualt sprite will be used.
func New(sprite *Sprite) *Spinner {
	if sprite == nil {
		sprite = NewSprite(SpriteFrames[10], (time.Millisecond * 100).Nanoseconds())
	}

	return &Spinner{
		RefreshInterval: time.Millisecond * 100,
		writer:          uilive.New(),
		sprite:          sprite,
		running:         false,
		dirty:           false,
	}
}

// Start starts the spinner while executing the provided function.
// This functions blocks until the end of the function, which runs in
// it's own goroutine.
func (s *Spinner) Start(f func()) {
	s.running = true
	s.writer.Start()

	// start user action
	go func() {
		f()
		s.running = false
	}()

	// print spinner
	for s.running {
		if s.sprite.Update() || s.dirty {
			s.print()
		}

		time.Sleep(s.RefreshInterval)
	}

	// last frame
	if len(s.LastFrame) > 0 {
		fmt.Fprintln(s.writer, s.LastFrame, s.text)
	} else {
		s.print()
	}
	s.writer.Stop()
}

// Print updates the status text of the spinner
func (s *Spinner) Print(text string) {
	s.text = text
	s.dirty = true
}

// Printf updates the status text of the spinner
func (s *Spinner) Printf(format string, p ...interface{}) {
	s.Print(fmt.Sprintf(format, p...))
}

func (s *Spinner) print() {
	fmt.Fprintln(s.writer, s.sprite.Frame(), s.text)
}
