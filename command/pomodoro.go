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
	StartTime   time.Time
	nowTime     int
	startChan   chan int
	nowTimeChan chan time.Duration
	isBegin     bool
}

func (p *Pomodoro) runtimer(setTime int) {
	timer := time.NewTimer(time.Second * 1)
	p.nowTime = 0
	go func() {
		for {
			select {
			case <-p.startChan:
				p.nowTime = 0
				p.isBegin = true

			case <-timer.C:
				if p.isBegin {
					p.nowTime++
					p.nowTimeChan <- NotifiactionTime
					if p.nowTime == setTime {
						p.isBegin = false
					}
				}
			}
		}
	}()
}

// NewPomodoro is to create instance
func NewPomodoro() *Pomodoro {
	return &Pomodoro{
		startChan:   make(chan int),
		nowTimeChan: make(chan time.Duration),
	}
}

// Init Pomodoro
func (p *Pomodoro) Init() error {
	p.runtimer(int(PomodoroTime.Seconds()))
	return nil
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
	if p.nowTime == 0 {
		return fmt.Sprint("Pomodoro"), nil
	}
	return fmt.Sprint(p.nowTime), nil
}

func getFormatedNowTime(nowTime int) string {
	min := nowTime / 60
	sec := nowTime - min*60
	return fmt.Sprint(min, ":", sec)
}
