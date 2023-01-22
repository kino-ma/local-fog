package apps

import (
	"fmt"
	"local-fog/core/types"
)

var Applications map[types.AppId]types.AppFunction = map[types.AppId]types.AppFunction{
	1: Hello,
	2: RecognizeFace,
}

func RunApp(appId types.AppId, in []byte) ([]byte, error) {
	f, found := Applications[appId]
	if !found {
		err := fmt.Errorf("app with id %v not found", appId)
		return nil, err
	}

	return f(in)
}
