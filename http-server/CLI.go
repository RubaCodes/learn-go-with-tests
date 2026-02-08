package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

func (c *CLI) PlayPoker() {
	c.playerStore.RecordWin(extractWinner(c.readLine()))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
