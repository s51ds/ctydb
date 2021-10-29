package api

import (
	"fmt"
	"github.com/s51ds/ctydb/ad1c"
	"github.com/s51ds/ctydb/cty"
	"github.com/s51ds/ctydb/parser"
)

var ctyDatRecords map[string]cty.Dat

func Load() {
	var err error
	fmt.Println("Loading Cty File...")
	if ctyDatRecords, err = parser.ParseCtyDatRecords(ad1c.CtyDatFile4test); err != nil {
		panic(err)
	} else {
		fmt.Println("Cty File loaded, number of ctyDatRecords:", len(ctyDatRecords))
	}
}
