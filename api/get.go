package api

import (
	"errors"
	"fmt"
	"github.com.s51ds/ctydb/cty"
)

var records map[string]cty.Dat

func Get(callSign string) (countryData cty.Dat, err error) {
	fmt.Println(records["S5"])
	for i := len(callSign); i >= 0; i-- {
		s := callSign[:i]
		if v, has := records[s]; has {
			return v, nil
		}
	}
	return countryData, errors.New(fmt.Sprintf("no country data for call sign:%s", callSign))
}
