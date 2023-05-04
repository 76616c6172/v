package main

import (
	"log"

	"github.com/76616c6172/help"
	"github.com/76616c6172/v/bluetooth"
	"github.com/76616c6172/v/brightness"
	"github.com/76616c6172/v/chat"
	"github.com/76616c6172/v/volume"
	"github.com/76616c6172/v/wifi"

	Z "github.com/rwxrob/bonzai/z"
)

func init() {
	Z.Dynamic["uname"] = func() string { return Z.Out("uname", "-a") }
	Z.Dynamic["ls"] = func() string { return Z.Out("ls", "-l", "-h") }
}

func main() {

	// remove log prefixes
	log.SetFlags(0)

	// provide panic trace
	Z.AllowPanic = true

	// can run in multicall, or monolith, not both

	/*
		// MULTICALL (status, afk, etc. linked)
		// (no completion unless set for individual commands)
		// (requires creation of hard/sym links or copies)
		Z.Commands = map[string][]any{
			// "conf": {conf.Cmd}, // bork cuz no multicall mode
			"yq":     {yq.Cmd},
			"y2j":    {y2j.Cmd},
			"status": {tmux, "update"},
			"afk":    {twitch.Cmd, "chat", "!afk"},
		}
		Z.Run()
	*/

	Cmd.Run()
}

var Cmd = &Z.Cmd{

	Name:      `v`,
	Summary:   `vulpix CLI`,
	Copyright: `Copyright 2023 valar`,
	Version:   `v0.0.1`,
	License:   `Apache-2.0`,
	Site:      `76616c6172.com`,
	Source:    `git@github.com:76616c6172/v.git`,
	Issues:    `github.com/76616c6172/v/issues`,

	Description: `
		CLI tool to interact with my laptop (vulpix)
		`,

	Commands: []*Z.Cmd{
		help.Cmd,
		bluetooth.Cmd,
		brightness.Cmd,
		chat.Cmd,
		volume.Cmd,
		wifi.Cmd,
	},

	Shortcuts: Z.ArgMap{
		//`project`:   {`twitch`, `bot`, `commands`, `edit`, `project`},
		//`status`:    {`tmux`, `update`},
		//`offscreen`: {`chat`, `!offscreen`},
		//`info`:      {`twitch`, `bot`, `commands`, `file`, `edit`},
		//`sync`:      {`twitch`, `bot`, `commands`, `sync`},
		//`work`:      {`go`, `work`},
		//`chat`:      {`twitch`, `chat`},
		//`afk`:       {`twitch`, `chat`, `!afk`},
		//`isosec`:    {`uniq`, `isosec`},
		//`isonan`:    {`uniq`, `isonan`},
		//`isodate`:   {`uniq`, `isodate`},
		//`uuid`:      {`uniq`, `uuid`},
		//`epoch`:     {`uniq`, `second`},
		//`path`:      {`env`, `get`, `path`},
		//`ytlink`:    {`filter`, `youtube`, `linkify`},
		//`long version of path`: {`env`, `get`, `path`},
	},
}
