package main

import (
	"github.com/tanyoAH/tanyo-server/setup"
	"github.com/tanyoAH/tanyo-server/startup"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	startCmd    = kingpin.Command("start", "Start the server").Default()
	setupCmd    = kingpin.Command("setup", "Setup the data in the database")
	resetDBFlag = setupCmd.Flag("resetdb", "Wipe and reset the database first").Bool()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("철수")
	kingpin.CommandLine.Help = "Tanyo App Server"
	switch kingpin.Parse() {
	case "start":
		startup.StartServer()
	case "setup":
		setup.StartSetup(*resetDBFlag)
	}
}
