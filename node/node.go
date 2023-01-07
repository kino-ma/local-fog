package main

import (
	"fmt"

	"local-fog/core"
)

type Node struct {
}

func (n Node) HandlePing(p core.Ping) {
	fmt.Printf("ping: %+v\n", p)
}

func (n Node) HandleSync(s core.Sync) {
	fmt.Printf("sync: %+v\n", s)

}

func (n Node) HandleCall(c core.Call) {
	fmt.Printf("call: %+v\n", c)
}

func (n Node) HandleGetProgram(g core.GetProgram) {
	fmt.Printf("get program: %+v\n", g)
}
