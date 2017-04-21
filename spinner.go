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
// This functions blocks until the end of the supplied function. The rendering of
// the spinner + text is done on a different goroutine. The user function will be
// executed on the main thread.
func (s *Spinner) Start(f func()) {
	done := make(chan bool)

	// render in own go routine
	go func() {
		s.running = true
		s.writer.Start()
		defer s.writer.Stop()

		// render loop
		for s.running {
			if s.sprite.Update() || s.dirty {
				s.print()
			}

			time.Sleep(s.RefreshInterval)
		}

		// render last frame
		if len(s.LastFrame) > 0 {
			fmt.Fprintln(s.writer, s.LastFrame, s.text)
		} else {
			s.print()
		}

		// done
		done <- true
	}()

	// execute user function on main thread
	f()

	// wait until render loop exists
	s.running = false
	<-done
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
