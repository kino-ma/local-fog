package main

import "fmt"

type Node struct {
}

func (n Node) HandlePing(p Ping) {
	fmt.Printf("ping: %+v\n", p)
}

func (n Node) HandleSync(s Sync) {
	fmt.Printf("sync: %+v\n", s)

}

func (n Node) HandleCall(c Call) {
	fmt.Printf("call: %+v\n", c)
}

func (n Node) HandleGetProgram(g GetProgram) {
	fmt.Printf("get program: %+v\n", g)
}
