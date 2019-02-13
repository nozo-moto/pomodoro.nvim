package command

import "time"

type Pomodoro struct {
	StartTime time.Time
}

func NewPomodoro() *Pomodoro {
	p := &Pomodoro{}
	return p
}

func (p *Pomodoro) Start() (string, error) {
	return "Start", nil
}

func (p *Pomodoro) Stop() (string, error) {
	return "Stop", nil
}

func (p *Pomodoro) Cancel() (string, error) {
	return "Cancel", nil
}
