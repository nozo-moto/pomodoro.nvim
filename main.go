package main

import (
	"strings"

	"github.com/neovim/go-client/nvim/plugin"
	"github.com/nozo-moto/pomodoro.nvim/command"
)

func hello(args []string) (string, error) {
	return "Hello " + strings.Join(args, " "), nil
}

func main() {
	pomodoro := command.NewPomodoro()
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleCommand(&plugin.CommandOptions{
			Name: "PomodoroStart"},
			pomodoro.Start,
		)
		p.HandleCommand(&plugin.CommandOptions{
			Name: "PomodoroStop"},
			pomodoro.Stop,
		)
		p.HandleCommand(&plugin.CommandOptions{
			Name: "PomodoroCancel"},
			pomodoro.Cancel,
		)

		p.HandleFunction(&plugin.FunctionOptions{Name: "Hello"}, hello)
		return nil
	})
}
