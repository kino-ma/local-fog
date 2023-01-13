package core

import (
	"fmt"
	"local-fog/core/types"
	"strconv"
	"strings"
)

const TXT_V_LOCALFOG = "v=localfog"

var ErrNotLocalFogService error = fmt.Errorf("given txt is not a localfog record")

func ParseTxt(txt string) (*types.NodeInfo, error) {
	info := &types.NodeInfo{}
	words := strings.Split(txt, " ")

	if len(words) < 2 || words[0] != TXT_V_LOCALFOG {
		err := ErrNotLocalFogService
		return info, err
	}

	for _, word := range words {
		if word[:3] == "id=" {
			id, err := strconv.ParseUint(word[3:], 16, 64)

			if err != nil {
				err := fmt.Errorf("failed to parse id: %v", err)
				return info, err
			}

			info.Id = id
		}
	}

	return info, nil
}
