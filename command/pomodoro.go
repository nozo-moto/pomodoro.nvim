package command

import (
	"fmt"
	"time"

	"github.com/neovim/go-client/nvim"
)

var (
	// NotifiactionTime is hogehoge
	NotifiactionTime = time.Second * 1
	// PomodoroTime is setTime
	PomodoroTime = time.Minute * 25
)

// Pomodoro struct
type Pomodoro struct {
	StartTime            time.Time
	nowTime              int
	timeNotificationChan *time.Ticker
	startChan            chan int
}

// Timer is Pomodoro Timer
func (p *Pomodoro) Timer(setTime int) error {
	var isBegin bool
	for {
		select {
		case <-p.timeNotificationChan.C:
			if isBegin {
				p.nowTime += int(NotifiactionTime)
				if p.nowTime == setTime {
					isBegin = false
					break
				}
			}
		case <-p.startChan:
			p.nowTime = 0
			isBegin = true
		}
	}
}

// NewPomodoro is to create instance
func NewPomodoro(timeNotificationChan *time.Ticker, startChan chan int) *Pomodoro {
	p := &Pomodoro{
		timeNotificationChan: timeNotificationChan,
		startChan:            startChan,
	}

	return p
}

// Start Pomodoro
func (p *Pomodoro) Start(v *nvim.Nvim, args []string) (string, error) {
	p.startChan <- 1
	return "Start", nil
}

// Stop Pomodoro
func (p *Pomodoro) Stop(v *nvim.Nvim, args []string) (string, error) {
	return "Stop", nil
}

// Cancel Pomodoro
func (p *Pomodoro) Cancel(v *nvim.Nvim, args []string) (string, error) {
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
