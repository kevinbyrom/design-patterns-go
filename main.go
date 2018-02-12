package main

import "fmt"
import "github.com/kevinbyrom/design-patterns-go/commands"
import "github.com/kevinbyrom/design-patterns-go/builders"

func main() {

	CommandExample()
	BuilderExample()

}

func CommandExample() {
	printThreeCmd := commands.NewSingleCommand(func() { fmt.Println("CommandThree") })
	printTwoCmd := commands.NewChainCommand(func() { fmt.Println("CommandTwo") }, printThreeCmd)
	printOneCmd := commands.NewChainCommand(func() { fmt.Println("CommandOne") }, printTwoCmd)

	ExecuteCommand(printOneCmd)
}

func ExecuteCommand(cmd commands.Command) {
	cmd.Execute()
}

func BuilderExample() {
	b := builders.NewCommandBuilder()

	b.Add(func() { fmt.Println("BuilderOne") })
	b.Add(func() { fmt.Println("BuilderTwo") })
	b.Add(func() { fmt.Println("BuilderThree") })

	b.GetCommand().Execute()
}
