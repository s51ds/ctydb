package api

import (
	"errors"
	"fmt"
	"github.com.s51ds/ctydb/cty"
)

var records map[string]cty.Dat

func Get(callSign string) (countryData cty.Dat, err error) {
	if v, has := records[callSign]; has {
		return v, nil
	} else {
		return countryData, errors.New(fmt.Sprintf("no country data for call sign:%s", callSign))
	}
}
