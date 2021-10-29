package api

import (
	"errors"
	"fmt"
	"github.com/s51ds/ctydb/cty"
	"strings"
	"sync"
)

var (
	cache = make(map[string]cty.Dat) // Key is callSign
	mutex = &sync.RWMutex{}
)

func Get(callSign string) (countryData cty.Dat, err error) {
	callSign = strings.TrimSpace(callSign)
	callSign = strings.ToUpper(callSign)

	// check cache
	has := false
	mutex.RLock()
	countryData, has = cache[callSign]
	mutex.RUnlock()
	if has {
		return
	}
	// not in cache, check if complete callSign is in DB
	if countryData, has = ctyDatRecords[callSign]; has {
		mutex.Lock()
		cache[callSign] = countryData
		mutex.Unlock()
		return
	}

	// following algorithm try to find cty.Dat...
	origCallSign := callSign

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
		if countryData, has = ctyDatRecords[s]; has {
			mutex.Lock()
			cache[origCallSign] = countryData
			mutex.Unlock()
			return
		}
	}

	countryData.CountryName = "Unknown"
	mutex.Lock()
	cache[origCallSign] = countryData
	mutex.Unlock()
	return countryData, errors.New(fmt.Sprintf("no country data for call sign:%s", callSign))
}
