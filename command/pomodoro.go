package command

import (
	"fmt"
	"time"

	"github.com/neovim/go-client/nvim"
)

const (
	first = iota
	working
	stop
)

var (
	notifiactionTime = time.Second * 1
	pomodoroTime     = time.Minute * 25
)

// Pomodoro struct
type Pomodoro struct {
	StartTime time.Time
	state     int
	nowTime   int
}

func (p *Pomodoro) timer(nowCh chan int, setTime int) error {
	countTime := 0
	timeNotification := time.NewTicker(time.Second * notifiactionTime)
	for {
		select {
		case <-timeNotification.C:
			nowCh <- int(countTime)
			if countTime == setTime {
				goto L
			}
			countTime += int(notifiactionTime)
		}
	}
L:

	return nil
}

// NewPomodoro is to create instance
func NewPomodoro() *Pomodoro {
	p := &Pomodoro{}
	return p
}

// Start Pomodoro
func (p *Pomodoro) Start(v *nvim.Nvim, args []string) (string, error) {
	if p.state != stop {
		p.StartTime = time.Now()
	}
	nowCh := make(chan int)
	go p.timer(nowCh, int(pomodoroTime))

	p.state = working
	return "Start", nil
}

// Stop Pomodoro
func (p *Pomodoro) Stop(v *nvim.Nvim, args []string) (string, error) {
	p.state = stop
	return "Stop", nil
}

// Cancel Pomodoro
func (p *Pomodoro) Cancel(v *nvim.Nvim, args []string) (string, error) {
	p.state = first
	return "Cancel", nil
}

// Status Pomodoro
func (p *Pomodoro) Status(v *nvim.Nvim, args []string) (string, error) {
	return getFormatedNowTime(p.nowTime), nil
}

func getFormatedNowTime(nowTime int) string {
	min := nowTime % 60
	sec := nowTime - min
	return fmt.Sprint(min, ":", sec)
}
