package apps

import "local-fog/core/types"

var Applications map[types.AppId]types.AppFunction = map[types.AppId]types.AppFunction{
	1: Hello,
}
