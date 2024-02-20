package cmd

type Command interface {
    Execute([]string) string
}
