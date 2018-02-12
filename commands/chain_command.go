package commands

type ChainCommand struct {
	SingleCommand
	nextCommand Command
}

func (c *ChainCommand) Execute() {
	c.executeFunc()

	if c.nextCommand != nil {
		c.nextCommand.Execute()
	}
}

func NewChainCommand(executeFunc func(), nextCommand Command) *ChainCommand {
	cmd := new(ChainCommand)
	cmd.executeFunc = executeFunc
	cmd.nextCommand = nextCommand
	return cmd
}
