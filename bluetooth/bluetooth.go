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
	Commands: []*Z.Cmd{help.Cmd, connectBudsProCmd, disconnectBudsProCmd},
}

var connectBudsProCmd = &Z.Cmd{
	Name:    `connect_pixel_buds_pro`,
	Summary: `connect Pixel Buds Pro`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("bluetoothctl", "connect", "24:29:34:A7:12:1D")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running bluetoothctl command:", err)
			return err
		}
		fmt.Println("Conected: Pixel Buds Pro")

		return nil
	},
}

var disconnectBudsProCmd = &Z.Cmd{
	Name:    `disconnect_pixel_buds_pro`,
	Summary: `disconnect Pixel Buds Pro`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("bluetoothctl", "disconnect", "24:29:34:A7:12:1D")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running bluetoothctl command:", err)
			return err
		}
		fmt.Println("Disconected: Pixel Buds Pro")

		return nil
	},
}
