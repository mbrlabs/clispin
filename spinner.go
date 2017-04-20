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

	dirty            bool
	text             string
	lastSpriteUpdate int64
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
	s.lastSpriteUpdate = time.Now().UnixNano()

	// start user action
	go func() {
		f()
		s.running = false
	}()

	// print spinner
	for s.running {
		updateRequired := false

		// update text if spinner sprite needs to be updated
		now := time.Now().UnixNano()
		if now-s.lastSpriteUpdate >= s.sprite.updateInterval {
			s.sprite.next()
			s.lastSpriteUpdate = now
			updateRequired = true
		}

		// update text if dirty flag set
		if s.dirty {
			updateRequired = true
		}

		if updateRequired {
			s.print()
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func (s *Spinner) Print(text string) {
	s.text = text
	s.dirty = true
}

func (s *Spinner) Printf(format string, p ...interface{}) {
	s.Print(fmt.Sprintf(format, p...))
}

func (s *Spinner) print() {
	fmt.Fprintln(s.writer, s.sprite.current()+" "+s.text)
}
