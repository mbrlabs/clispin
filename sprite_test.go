package clispin

import (
	"testing"
	"time"
)

func TestSprite(t *testing.T) {
	sprite := NewSprite([]string{"1", "2", "3"})
	sprite.Interval = time.Millisecond.Nanoseconds() * 500 // 500 ms

	// no update
	if sprite.Update() || sprite.Frame() != "1" {
		t.Error("Frame: " + sprite.Frame())
	}

	// first update - no pause
	if sprite.Update() || sprite.Frame() != "1" {
		t.Error("Frame: " + sprite.Frame())
	}

	// wait long enough to reach frame 2
	time.Sleep(time.Millisecond * 510)
	if !sprite.Update() || sprite.Frame() != "2" {
		t.Error("Frame: " + sprite.Frame())
	}

	// frame 3
	time.Sleep(time.Millisecond * 510)
	if !sprite.Update() || sprite.Frame() != "3" {
		t.Error("Frame: " + sprite.Frame())
	}

	//  back to first frame
	time.Sleep(time.Millisecond * 510)
	if !sprite.Update() || sprite.Frame() != "1" {
		t.Error("Frame: " + sprite.Frame())
	}
}
