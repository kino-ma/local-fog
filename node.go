package main

type Node struct {
}

func (n Node) HandlePing(p Ping)             {}
func (n Node) HandleSync(s Sync)             {}
func (n Node) HandleCall(c Call)             {}
func (n Node) HandleGetProgram(g GetProgram) {}
