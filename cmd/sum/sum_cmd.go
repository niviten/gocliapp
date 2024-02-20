package sumcmd

import (
	"fmt"
	"strconv"
)

type SumCommand struct{}

func NewSumCommand() *SumCommand {
    return &SumCommand{}
}

func (cmd *SumCommand) Execute(inputs []string) string {
    var sum float64 = 0
    for _, input := range inputs {
        if len(input) == 0 {
            continue
        }
        num, err := strconv.ParseFloat(input, 64)
        if err != nil {
            return fmt.Sprintf("Error: Invalid number: %s", input)
        }
        sum = sum + num
    }
    return fmt.Sprintf("Sum: %f\n", sum)
}
