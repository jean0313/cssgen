package main


import (
	"config"
)

func main() {
	tm := &config.TemplateManager{}
	tm.LoadConfig("config.json")
	tm.DebugDump()
	tm.Gen()
}
