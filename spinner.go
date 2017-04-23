package clispin

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/mbrlabs/uilive"
)

type Color int

const (
	ColorNone    = -1
	ColorBlack   = Color(color.FgBlack)
	ColorRed     = Color(color.FgRed)
	ColorGreen   = Color(color.FgGreen)
	ColorYellow  = Color(color.FgYellow)
	ColorBlue    = Color(color.FgBlue)
	ColorMagenta = Color(color.FgMagenta)
	ColorCyan    = Color(color.FgCyan)
	ColorWhite   = Color(color.FgWhite)
)

// Spinner is an infinite ui spinner with status text
type Spinner struct {
	RefreshInterval time.Duration
	LastFrame       string

	writer *uilive.Writer
	sprite *Sprite

	mutex   *sync.Mutex
	running bool
	dirty   bool
	text    string
}

// New creates a new Spinner with the given sprite.
// If sprite nil, a defualt sprite will be used.
func New(sprite *Sprite) *Spinner {
	if sprite == nil {
		sprite = NewSprite(SpriteFrames[10])
	}

	return &Spinner{
		RefreshInterval: time.Millisecond * 100,
		writer:          uilive.New(),
		sprite:          sprite,
		running:         false,
		dirty:           false,
		mutex:           &sync.Mutex{},
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
		s.mutex.Lock()
		s.running = true
		s.writer.Start()
		s.mutex.Unlock()
		defer s.writer.Stop()

		// render loop
		for s.running {
			s.mutex.Lock()
			if s.sprite.Update() || s.dirty {
				s.dirty = false
				s.print(s.sprite.Frame())
			}
			s.mutex.Unlock()

			time.Sleep(s.RefreshInterval)
		}

		// render last frame
		s.mutex.Lock()
		if len(s.LastFrame) > 0 {
			s.print(s.LastFrame)
		} else {
			s.print(s.sprite.Frame())
		}
		s.writer.Flush()
		s.mutex.Unlock()

		// done
		done <- true
	}()

	f() // execute user function on main thread

	s.running = false
	<-done // wait until render loop is done
}

// Color sets the color for the spinner sprite
func (s *Spinner) Color(c Color) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c == -1 {
		s.sprite.color = nil
	} else {
		s.sprite.color = color.New(color.Attribute(c))
	}
}

// Print updates the status text of the spinner
func (s *Spinner) Print(text string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.text = text
	s.dirty = true
}

// Printf updates the status text of the spinner
func (s *Spinner) Printf(format string, p ...interface{}) {
	s.Print(fmt.Sprintf(format, p...))
}

func (s *Spinner) print(frame string) {
	if s.sprite.color != nil {
		s.sprite.color.Fprintf(s.writer, frame) // sprite with color
	} else {
		fmt.Fprintf(s.writer, frame) // sprite without color
	}

	// text
	fmt.Fprintf(s.writer, " "+s.text+"\n")
}
