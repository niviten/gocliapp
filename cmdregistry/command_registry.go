package cmdregistry

import (
	"cli/cmd"
	"fmt"
)

type CommandRegistry struct {
    commands map[string]cmd.Command
}

func newCommandRegistry() *CommandRegistry {
    return &CommandRegistry{
        commands: make(map[string]cmd.Command),
    }
}

var instance *CommandRegistry

func GetInstance() *CommandRegistry {
    if instance == nil {
        instance = newCommandRegistry()
    }
    return instance
}

func (r *CommandRegistry) Register(name string, command cmd.Command) {
    if _, exists := r.commands[name]; exists {
        fmt.Printf("Command already registered: %s\n", name)
        return
    }
    r.commands[name] = command
}

func (r *CommandRegistry) UnregisterByName(name string) {
    delete(r.commands, name)
}

func (r *CommandRegistry) UnregisterByObject(command cmd.Command) {
    for name, registeredCommand := range r.commands {
        if command == registeredCommand {
            r.UnregisterByName(name)
            return
        }
    }
}

func (r *CommandRegistry) GetCommand(name string) (cmd.Command, bool) {
    cmd, exists := r.commands[name]
    return cmd, exists
}

func (r *CommandRegistry) PrintRegisteredCommands() {
    for name := range r.commands {
        fmt.Println(name)
    }
}
