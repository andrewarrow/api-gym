package screen

import (
	"os"

	"golang.org/x/term"
)

func Run() {
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		if c == 3 { // control-c
			term.Restore(int(os.Stdin.Fd()), oldState)
			break
		}
	}
}
