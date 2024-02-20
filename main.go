package main

import (
	sumcmd "cli/cmd/sum"
	"cli/cmdregistry"
	"flag"
	"fmt"
)

func main() {
    registerCommands()
    // cmdregistry.GetInstance().PrintRegisteredCommands()

    commandNamePtr := flag.String("cmd", "",
        "The name of the command to execute")

    flag.Parse()

    command, found := cmdregistry.GetInstance().GetCommand(*commandNamePtr)
    if !found {
        fmt.Println("Error: Invalid command")
        return
    }

    args := flag.Args()
    output := command.Execute(args)
    fmt.Println(output)

    fmt.Println("___done")
}

func registerCommands() {
    cmdregistry.GetInstance().Register("sum", sumcmd.NewSumCommand())
}
