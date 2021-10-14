package api

import (
	"github.com.s51ds/ctydb/ad1c"
	"github.com.s51ds/ctydb/parser"
)

func Load() {

	if r, err := parser.ParseCtyDatRecords(ad1c.CtyDatFile4test); err != nil {
		panic(err)
	} else {
		records = r
	}

}
