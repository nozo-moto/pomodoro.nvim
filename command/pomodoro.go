package command

import "time"

// pomodoro struct
type Pomodoro struct {
	StartTime time.Time
}

// create instance
func NewPomodoro() *Pomodoro {
	p := &Pomodoro{}
	return p
}

func (p *Pomodoro) Start(args []string) (string, error) {
	return "Start", nil
}

func (p *Pomodoro) Stop(args []string) (string, error) {
	return "Stop", nil
}

func (p *Pomodoro) Cancel(args []string) (string, error) {
	return "Cancel", nil
}
