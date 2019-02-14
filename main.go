package main

import (
	"github.com/neovim/go-client/nvim/plugin"
	"github.com/nozo-moto/pomodoro.nvim/command"
)

func main() {
	pomodoro := command.NewPomodoro()
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrStart"}, pomodoro.Start)
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrStop"}, pomodoro.Stop)
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrCancel"}, pomodoro.Cancel)
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrStatus"}, pomodoro.Status)
		p.HandleFunction(&plugin.FunctionOptions{Name: "PmdrInit"}, pomodoro.Init)
		return nil
	})

}
