package main

import (
	"github.com/neovim/go-client/nvim/plugin"
	"github.com/nozo-moto/pomodoro.nvim/command"
)

func main() {
	pomodoro := command.NewPomodoro()
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleCommand(&plugin.CommandOptions{
			Name: "pmdrStart"},
			pomodoro.Start,
		)
		p.HandleCommand(&plugin.CommandOptions{
			Name: "pmdrStop"},
			pomodoro.Stop,
		)
		p.HandleCommand(&plugin.CommandOptions{
			Name: "pmdrCancel"},
			pomodoro.Cancel,
		)

		return nil
	})
}
