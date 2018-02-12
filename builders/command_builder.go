package builders

import "github.com/kevinbyrom/design-patterns-go/commands"

type CommandBuilder struct {
	cmds []func()
}

func NewCommandBuilder() *CommandBuilder {
	b := new(CommandBuilder)
	return b
}

func (b *CommandBuilder) Add(f func()) {
	b.cmds = append(b.cmds, f)
}

func (b *CommandBuilder) GetCommand() commands.Command {
	var cmd *commands.ChainCommand

	for i := len(b.cmds) - 1; i >= 0; i-- {
		f := b.cmds[i]

		if cmd == nil {
			cmd = commands.NewChainCommand(f, nil)

		} else {
			cmd = commands.NewChainCommand(f, cmd)
		}
	}

	return cmd
}
