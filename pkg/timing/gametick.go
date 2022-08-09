package timing

import "time"

type Tick struct {
	t     int
	f     float64
	timer *Timer
	tick  bool
}

func NewTick(t int) *Tick {
	f := float64(time.Second / time.Duration(t))
	return &Tick{
		t:     t,
		timer: New(f),
	}
}

func (t *Tick) Update() {
	t.tick = false
	if t.timer.UpdateDone() {
		t.tick = true
		t.timer = New(t.f)
	}
}

func (t *Tick) Tick() bool {
	return t.tick
}