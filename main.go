package main

import (
	"github.com/neovim/go-client/nvim/plugin"
	"github.com/nozo-moto/pomodoro.nvim/command"
)

func main() {
	pomodoro := command.NewPomodoro()
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleCommand(&plugin.CommandOptions{Name: "PmdrStart"}, pomodoro.Start)
		p.HandleCommand(&plugin.CommandOptions{Name: "PmdrStop"}, pomodoro.Stop)
		p.HandleCommand(&plugin.CommandOptions{Name: "PmdrCancel"}, pomodoro.Cancel)
		p.HandleCommand(&plugin.CommandOptions{Name: "PmdrStatus"}, pomodoro.Status)
		return nil
	})
}
