package main

import (
	"fmt"
	"local-fog/core"
	"local-fog/core/types"
	"log"
	"time"
)

func call(c *core.FogConsumer, req *types.CallRequest) (*types.CallReply, time.Duration, error) {
	s := time.Now()

	cr, err := c.Call(req)
	if err != nil {
		err = fmt.Errorf("call: failed to Call: %w", err)
		log.Print(err)
		return nil, 0, err
	}

	log.Printf("Call success: %v", cr)
	e := time.Since(s)
	log.Printf("took %s", e)

	return cr, e, nil
}
