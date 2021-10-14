package parser

import (
	"fmt"
	"github.com.s51ds/ctydb/cty"
	"strings"
)

// ParseCtyDatRecords parses ctyDatRecords and returns map, key is Primary or Alias DXCC Prefix.
// ParseCtyDatRecords panics if ctyDatRecords is wrong formatted.
func ParseCtyDatRecords(ctyDatRecords string) (m map[string]cty.Dat, err error) {
	m = make(map[string]cty.Dat)
	ctyDatRecords = removeComments(ctyDatRecords)
	records := strings.Split(ctyDatRecords, ";")

	numberOfRecords := len(records)
	ch := make(chan []cty.Dat, numberOfRecords)

	for _, rec := range records {
		if len(rec) > 10 {
			go func(r string) {
				v, e := parseCtyDatRecord(r)
				if e != nil {
					fmt.Println(r)
					fmt.Println(e)
					panic(e.Error())
				}
				ch <- v
			}(rec)
		} else {
			ch <- []cty.Dat{}
		}
	}

	j := numberOfRecords
	for j > 0 {
		data := <-ch
		for _, item := range data {
			m[item.AliasPrefix] = item
		}
		j--
	}

	return m, nil
}
