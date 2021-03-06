package main

import (
	"log"
	"os/user"

	tl "github.com/JoelOtter/termloop"
)

type Cemetery struct {
	*Shell
	Screen *tl.Screen
}

func BuildCemetery(screen *tl.Screen) *Cemetery {
	return &Cemetery{Shell: NewShell(screen), Screen: screen}
}

func (c *Cemetery) Greet(s string) {
	// get currently logged in user
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	lines := append(
		GetStdOutLinesFromString(linuxLoginStdout, tl.ColorWhite, DarkBG, 0.1, func() float64 { return 1 }),
		NewStdOutLine(usr.Username+"@☠☠☠☠☠☠☠☠:~ $", tl.ColorGreen, DarkBG, 0.3),
	)

	txt := NewStdOut(lines...)
	c.AddEntity(txt)
}

func (c *Cemetery) EnterDreamer() {
	dreamer := NewDreamer(c.Screen)
	c.AddEntity(dreamer)
}
