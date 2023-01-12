package main

import (
	"fmt"
	"local-fog/core/types"
)

type AppBinary []byte

var Apps map[types.AppId]AppBinary = make(map[types.AppId]AppBinary)

func InsertApp(appId types.AppId, appBin AppBinary) error {
	_, found := Apps[appId]

	if found {
		return fmt.Errorf("app with id %v already exists", appId)
	}

	Apps[appId] = appBin

	return nil
}