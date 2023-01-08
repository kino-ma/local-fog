package main

import (
	"fmt"

	t "local-fog/core/types"
)

type Node struct {
}

func (n Node) HandlePing(p *t.Ping) {
	fmt.Printf("ping: %+v\n", p)
}

func (n Node) HandleSync(s *t.Sync) {
	fmt.Printf("sync: %+v\n", s)

}

func (n Node) HandleCall(c *t.Call) {
	fmt.Printf("call: %+v\n", c)
}

func (n Node) HandleGetProgram(g *t.GetProgram) {
	fmt.Printf("get program: %+v\n", g)
}
