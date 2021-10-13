package cty

import (
	"errors"
	"fmt"
)

var records map[string]Dat

func Get(callSign string) (countryData Dat, err error) {
	if v, has := records[callSign]; has {
		return v, nil
	} else {
		return countryData, errors.New(fmt.Sprintf("no country data for call sign:%s", callSign))
	}
}
