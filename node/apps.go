package main

import (
	"fmt"
	"local-fog/core/types"

	"golang.org/x/exp/maps"
)

type AppBinary []byte

var Apps map[types.AppId]AppBinary = make(map[types.AppId]AppBinary)

func InsertApp(appId types.AppId, appBin AppBinary) error {
	_, found := GetApp(appId)

	if found {
		return fmt.Errorf("app with id %v already exists", appId)
	}

	Apps[appId] = appBin

	return nil
}

func GetApp(appId types.AppId) (AppBinary, bool) {
	bin, found := Apps[appId]
	return bin, found
}

func GetApps(appIds []types.AppId) (map[types.AppId]AppBinary, error) {
	bins := make(map[types.AppId]AppBinary)

	for _, id := range appIds {
		bin, got := GetApp(id)

		if !got {
			return nil, fmt.Errorf("app with id %v not found", id)
		}

		bins[id] = bin
	}

	return bins, nil
}

func GetAppIds() []types.AppId {
	return maps.Keys(Apps)
}
