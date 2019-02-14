package main

import (
	"time"

	"github.com/neovim/go-client/nvim/plugin"
	"github.com/nozo-moto/pomodoro.nvim/command"
)

func main() {
	startChan := make(chan int)
	pomodoro := command.NewPomodoro(
		time.NewTicker(time.Second*command.NotifiactionTime),
		startChan,
	)
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrStart"}, pomodoro.Start)
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrStop"}, pomodoro.Stop)
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrCancel"}, pomodoro.Cancel)
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrStatus"}, pomodoro.Status)
		return nil
	})

	pomodoro.Timer(int(command.PomodoroTime))
}
