package parser

import (
	"errors"
	"fmt"
)

func errorWrongFormattedRecord(description string, ctyDataRecord string) error {
	s := "wrong formatted cty.dat record: " + ctyDataRecord
	if len(description) > 0 {
		s = fmt.Sprintf("%s\n%s", description, s)
	}
	return errors.New(s)
}
