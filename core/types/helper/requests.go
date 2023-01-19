package helper

import (
	"fmt"
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
)

func RequestForAllNode(nodes []*types.NodeInfoWrapper, request func(n *types.NodeInfoWrapper, consumer core.FogConsumer) error) error {
	errs := make([]error, 0, len(nodes))

	for _, n := range nodes {
		addr := utils.Uint32ToIp(n.AddrV4)
		consumer, err := core.Connect(addr.String(), core.DEFAULT_PORT)
		if err != nil {
			err = fmt.Errorf("RequestForAllNode: failed to connect to node [%v]: %w", n.Id, err)
			errs = append(errs, err)
			continue
		}

		err = request(n, consumer)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	if len(errs) == 0 {
		return nil
	}

	// if there is any errors

	errString := ""
	for _, err := range errs {
		errString += err.Error() + ", "
	}
	return fmt.Errorf("RequestForAllNode: 1 ore more errors occured while syncing: %v", errString)
}
