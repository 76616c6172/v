package volume

// Requires bluez package on debian

import (
	"fmt"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `volume`,
	Summary:  `change volume`,
	Commands: []*Z.Cmd{help.Cmd, volCmd0, volCmd40, volCmd60, volCmd80, volCmd100},
}

var volCmd0 = &Z.Cmd{
	Name:    `0`,
	Summary: `set volume to 0%`,
	Call: func(_ *Z.Cmd, args ...string) error {
		return setVolume("0%")
	},
}

var volCmd40 = &Z.Cmd{
	Name:    `40`,
	Summary: `set volume to 40%`,
	Call: func(_ *Z.Cmd, args ...string) error {
		return setVolume("40%")
	},
}

var volCmd60 = &Z.Cmd{
	Name:    `60`,
	Summary: `set volume to 60%`,
	Call: func(_ *Z.Cmd, args ...string) error {
		return setVolume("60%")
	},
}

var volCmd80 = &Z.Cmd{
	Name:    `80`,
	Summary: `set volume to 80%`,
	Call: func(_ *Z.Cmd, args ...string) error {
		return setVolume("80%")
	},
}

var volCmd100 = &Z.Cmd{
	Name:    `100`,
	Summary: `set volume to 100%`,
	Call: func(_ *Z.Cmd, args ...string) error {
		return setVolume("100%")
	},
}

func setVolume(lvl string) error {
	cmd := exec.Command("amixer", "set", "Master", lvl)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running amixer command:", err)
		return err
	}
	fmt.Printf("volume: %s\n", lvl)
	return nil

}
