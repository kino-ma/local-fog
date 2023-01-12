package main

import (
	"fmt"
	"local-fog/core/types"

	"golang.org/x/exp/maps"
)

type AppBinary []byte

var Apps map[types.AppId]AppBinary = make(map[types.AppId]AppBinary)

func InsertApp(app *types.Application) error {
	id := types.AppId(app.AppId)
	bin := app.Binary

	_, found := GetAppBinary(id)

	if found {
		return fmt.Errorf("app with id %v already exists", id)
	}

	Apps[id] = bin

	return nil
}

func GetAppBinary(appId types.AppId) (AppBinary, bool) {
	bin, found := Apps[appId]
	return bin, found
}

func GetApps(appIds []types.AppId) ([]types.Application, error) {
	apps := make([]types.Application, len(appIds))

	for _, id := range appIds {
		bin, got := GetAppBinary(id)

		if !got {
			return nil, fmt.Errorf("app with id %v not found", id)
		}

		apps[id] = types.Application{
			AppId:  uint64(id),
			Binary: bin,
		}
	}

	return apps, nil
}

func GetAppIds() []types.AppId {
	return maps.Keys(Apps)
}
