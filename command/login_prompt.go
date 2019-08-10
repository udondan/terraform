package command

import (
	"fmt"
	"io"

	"github.com/chzyer/readline"
	"github.com/hashicorp/terraform/helper/wrappedreadline"
)

func (c *LoginCommand) prompt(prompt string, sensitive bool) (string, error) {
	l, err := readline.NewEx(wrappedreadline.Override(&readline.Config{
		Prompt:            prompt + ": ",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
		EnableMask:        sensitive,
		MaskRune:          '*',
	}))
	if err != nil {
		return "", fmt.Errorf("cannot prompt for input: %s", err)
	}
	defer l.Close()

	line, err := l.Readline()
	if err == readline.ErrInterrupt || err == io.EOF {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to read input: %s", err)
	}

	return line, nil
}
