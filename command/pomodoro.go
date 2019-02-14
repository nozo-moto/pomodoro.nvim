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
	nowTimeChan          chan time.Duration
}

func (p *Pomodoro) runtimer(setTime int) {
	//var isBegin bool
	go func() {
		for {
			select {
			case <-p.timeNotificationChan.C:
				//if isBegin {
				p.nowTime += int(NotifiactionTime.Seconds())
				p.nowTimeChan <- NotifiactionTime
				if p.nowTime == setTime {
					isBegin = false
				}
				//}
			case <-p.startChan:
				p.nowTime = 0
				isBegin = true
			}
		}
	}()
}

// NewPomodoro is to create instance
func NewPomodoro() *Pomodoro {
	return &Pomodoro{
		timeNotificationChan: time.NewTicker(NotifiactionTime),
		startChan:            make(chan int),
		nowTimeChan:          make(chan time.Duration),
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
	return getFormatedNowTime(p.nowTime), nil
}

func getFormatedNowTime(nowTime int) string {
	min := nowTime / 60
	sec := nowTime - min*60
	return fmt.Sprint(min, ":", sec)
}
