package wifi

import (
	"fmt"
	"os/exec"
	"strings"

	bubbletea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	Z "github.com/rwxrob/bonzai/z"
)

type model struct {
	choices  []string
	cursor   int
	selected bool
}

var cursor int
var selected bool = false

var connectCmd = &Z.Cmd{
	Name:    `connect`,
	Summary: `connect to network (interative)`,
	Call: func(_ *Z.Cmd, args ...string) error {
		cmd := exec.Command("nmcli", "device", "wifi", "list")
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error running nmcli command:", err)
			return err
		}

		availNetworks := parseNetworks(string(output))
		choices := make([]string, len(availNetworks))
		for i, network := range availNetworks {
			choices[i] = network[1]
		}

		m := model{choices: choices}
		p := bubbletea.NewProgram(m)
		if _, err = p.Run(); err != nil {
			fmt.Printf("Error running bubble tea program: %v\n", err)
			return err
		}

		if selected {
			selectedNetwork := availNetworks[cursor]
			cmd := exec.Command("nmcli", "device", "wifi", "connect", selectedNetwork[0])
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error running nmcli command:", err)
				return err
			}
			fmt.Printf("Connected to %s at %s\n", availNetworks[cursor][1], selectedNetwork[0])
		} else {
			fmt.Println("No network selected.")
		}

		return nil
	},
}

// expects input from "nmcli d wifi list"
// returns output in the form off:
// [ [bssid, ssid] ]
// [ [bssid, ssid] ]
func parseNetworks(input string) [][]string {
	lines := strings.Split(input, "\n")
	availNetworks := make([][]string, 0)

	for i, line := range lines {
		if i == 0 || len(strings.TrimSpace(line)) == 0 {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 4 {
			bssid := fields[0]
			ssid := fields[1]
			if bssid == "*" {
				bssid = ssid
				ssid = fields[2]
			}

			availNetworks = append(availNetworks, []string{bssid, ssid})
		}
	}

	return availNetworks
}

func (m model) Init() bubbletea.Cmd {
	return nil
}

func (m model) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	switch msg := msg.(type) {
	case bubbletea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			m.selected = false
			return m, bubbletea.Quit
		case "enter", "y":
			m.selected = true
			selected = true
			cursor = m.cursor
			return m, bubbletea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	list := ""

	for i, choice := range m.choices {
		style := lipgloss.NewStyle()
		style.PaddingBottom(0)
		if i == m.cursor {
			style = style.Foreground(lipgloss.Color("205"))
		}
		list += style.Render(choice) + "\n"
	}

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(1).
		MarginBottom(1).
		Width(30)

	return boxStyle.Render(list)
}
