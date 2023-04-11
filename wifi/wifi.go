package wifi

import (
	"fmt"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `wifi`,
	Summary:  `manage wifi connection`,
	Commands: []*Z.Cmd{help.Cmd, statusCmd, connectCmd, onCmd, offCmd},
}

var statusCmd = &Z.Cmd{
	Name:    `status`,
	Summary: `show wifi status`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("nmcli", "device", "status")
		output, outErr := cmd.Output()
		if outErr != nil {
			fmt.Println("Error running nmcli command:", outErr)
			return outErr
		}
		fmt.Println(string(output))

		return nil
	},
}

var onCmd = &Z.Cmd{
	Name:    `on`,
	Summary: `turn wifi on`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("nmcli", "radio", "wifi", "on")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running nmcli command:", err)
			return err
		}
		fmt.Println("wifi: on")

		return nil
	},
}

var offCmd = &Z.Cmd{
	Name:    `off`,
	Summary: `turn wifi off`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("nmcli", "radio", "wifi", "off")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running nmcli command:", err)
			return err
		}
		fmt.Println("wifi: off")

		return nil
	},
}
