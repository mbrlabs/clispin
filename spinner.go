package clispin

import (
	"fmt"
	"time"

	"github.com/mbrlabs/uilive"
)

type Spinner struct {
	writer *uilive.Writer
	sprite *Sprite

	running bool

	dirty bool
	text  string
}

func New(sprite *Sprite) *Spinner {
	if sprite == nil {
		sprite = NewSprite(SpriteFrames[10])
	}

	return &Spinner{
		writer:  uilive.New(),
		sprite:  sprite,
		running: false,
		dirty:   false,
	}
}

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

		time.Sleep(time.Millisecond * 100)
	}

	s.print()
	s.writer.Stop()
}

func (s *Spinner) Print(text string) {
	s.text = text
	s.dirty = true
}

func (s *Spinner) Printf(format string, p ...interface{}) {
	s.Print(fmt.Sprintf(format, p...))
}

func (s *Spinner) print() {
	fmt.Fprintln(s.writer, s.sprite.Frame(), s.text)
}
