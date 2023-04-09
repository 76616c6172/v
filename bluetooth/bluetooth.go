package bluetooth

// Requires bluez package on debian

import (
	"fmt"
	"os/exec"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `bluetooth`,
	Summary:  `manage bluetooth devices`,
	Commands: []*Z.Cmd{help.Cmd, connectBudsCmd, disconnectBudsCmd},
}

var connectBudsCmd = &Z.Cmd{
	Name:    `connect_pixel_buds`,
	Summary: `connect Pixel Buds Pro`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("bluetoothctl", "connect", "24:29:34:A7:12:1D")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running bluetoothctl command:", err)
			return err
		}
		fmt.Println("conected to Pixel Buds Pro")

		return nil
	},
}

var disconnectBudsCmd = &Z.Cmd{
	Name:    `disconnect_pixel_buds`,
	Summary: `disconnect Pixel Buds Pro`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("bluetoothctl", "disconnect", "24:29:34:A7:12:1D")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running bluetoothctl command:", err)
			return err
		}
		fmt.Println("disconected to Pixel Buds Pro")

		return nil
	},
}
