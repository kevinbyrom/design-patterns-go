package commands

type SingleCommand struct {
	executeFunc func()
}

func (c *SingleCommand) Execute() {
	c.executeFunc()
}

func NewSingleCommand(executeFunc func()) *SingleCommand {
	cmd := new(SingleCommand)
	cmd.executeFunc = executeFunc
	return cmd
}
