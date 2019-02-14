package pomodoro

import "time"

var (
	notifiactionTime = time.Second * 1
	pomodoroTime     = time.Minute * 25
)

// Pomodoro struct
type Pomodoro struct{}

func timer(setTime int) error {
	countTime := 0
	timeNotification := time.NewTicker(time.Second * notifiactionTime)
	for {
		select {
		case <-timeNotification.C:
			if countTime == setTime {
				goto L
			}
			countTime += int(notifiactionTime)
		}
	}
L:
	return nil
}

// Start Timer
func (p *Pomodoro) Start() {
	if err := timer(int(pomodoroTime)); err != nil {
		panic(err)
	}

}
