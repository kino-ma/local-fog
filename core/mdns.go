package core

import (
	"fmt"
	"local-fog/core/types"
	"strconv"
	"strings"
)

const TXT_V_LOCALFOG = "v=localfog"

func ParseTxt(txt string) (info types.NodeInfo, err error) {
	words := strings.Split(txt, " ")

	if len(words) < 2 || words[0] != TXT_V_LOCALFOG {
		err = fmt.Errorf("given txt is not a localfog record")
		return info, err
	}

	for _, word := range words {
		if word[:3] == "id=" {
			id, err := strconv.ParseUint(word[3:], 16, 64)

			if err != nil {
				err = fmt.Errorf("failed to parse id: %v", err)
				return info, err
			}

			info.Id = id
		}
	}

	return info, nil
}
