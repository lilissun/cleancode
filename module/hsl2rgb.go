package main

import (
	"encoding/json"

	"bitbucket.org/lilissun/cleancode/test/color"
)

func hsl2rgb(req []byte) ([]byte, error) {
	var hsl color.HSLColor
	err := json.Unmarshal([]byte(req), &hsl)
	if err != nil {
		return nil, err
	}
	resp, err := json.Marshal(hsl.ToRGB())
	if err != nil {
		return nil, err
	}
	return resp, nil
}
