package brightness

import (
	"fmt"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `brightness`,
	Summary:  `change screen brightness`,
	Commands: []*Z.Cmd{help.Cmd, dimCmd, fullCmd},
}

var dimCmd = &Z.Cmd{
	Name:    `80%`,
	Summary: `80% screen brightness`,
	Call: func(_ *Z.Cmd, args ...string) error {
		level := "0.8"
		cmd := exec.Command("xrandr", "--output", "eDP-1", "--brightness", level)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running xrandr command:", err)
			return err

		}
		fmt.Println("brightness:", level)

		return nil
	},
}

var fullCmd = &Z.Cmd{
	Name:    `100%`,
	Summary: `100% screen brightness`,
	Call: func(_ *Z.Cmd, args ...string) error {
		level := "1.00"
		cmd := exec.Command("xrandr", "--output", "eDP-1", "--brightness", level)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running xrandr command:", err)
			return err

		}
		fmt.Println("brightness:", level)

		return nil
	},
}
