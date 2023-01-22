package main

import (
	"fmt"
	"io/ioutil"
)

const faceFile string = "../data/sample.jpeg"

func Face() ([]byte, error) {
	raw, err := ioutil.ReadFile(faceFile)
	if err != nil {
		err = fmt.Errorf("failed to read face image file: %w", err)
		return nil, err
	}

	return raw, nil
}
