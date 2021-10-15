package api

import (
	"errors"
	"fmt"
	"github.com.s51ds/ctydb/cty"
	"strings"
)

var records map[string]cty.Dat

func Get(callSign string) (countryData cty.Dat, err error) {
	callSign = strings.TrimSpace(callSign)
	callSign = strings.ToUpper(callSign)

	if strings.Contains(callSign, "/") {
		// handle /P /MM etc
		callSign = strings.TrimSuffix(callSign, "/P")
		callSign = strings.TrimSuffix(callSign, "/MM")
		callSign = strings.TrimSuffix(callSign, "/QRP")
		callSign = strings.TrimSuffix(callSign, "/M")
		callSign = strings.TrimSuffix(callSign, "/AM")
		callSign = strings.TrimSuffix(callSign, "/A")
		callSign = strings.TrimSuffix(callSign, "/LH")

		// handling such cases like AC2AI/KH2 --> KH2/AC2AI
		if strings.Contains(callSign, "/") {
			ss := strings.Split(callSign, "/")
			if len(ss) == 2 && len(ss[0]) > len(ss[1]) {
				callSign = ss[1] + "/" + ss[0]
			}
		}
	}

	for i := len(callSign); i >= 0; i-- {
		s := callSign[:i]
		if v, has := records[s]; has {
			return v, nil
		}
	}
	return countryData, errors.New(fmt.Sprintf("no country data for call sign:%s", callSign))
}
