package main

import (
	"github.com/neovim/go-client/nvim/plugin"
	"github.com/nozo-moto/pomodoro.nvim/command"
)

func main() {
	pomodoro := command.NewPomodoro()
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleCommand(&plugin.CommandOptions{
			Name: "pmdr#Start"},
			pomodoro.Start,
		)
		p.HandleCommand(&plugin.CommandOptions{
			Name: "pmdr#Stop"},
			pomodoro.Stop,
		)
		p.HandleCommand(&plugin.CommandOptions{
			Name: "pmdr#Cancel"},
			pomodoro.Cancel,
		)

		return nil
	})
}
