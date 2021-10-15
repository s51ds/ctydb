package api

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	L := len(cache) + 1
	Get("S59ABC")
	if len(cache) != L {
		t.Errorf(fmt.Sprintf("len(cache) != %d", L))
	}
	Get("S59ABC")
	if len(cache) != L {
		t.Errorf(fmt.Sprintf("len(cache) != %d", L))
	}

	_, err := Get("D1ABC")
	L++
	if len(cache) != L {
		t.Errorf(fmt.Sprintf("len(cache) != %d", L))
	}
	if err == nil {
		t.Errorf("Where is error no country data for call sign:D1ABC ??? ")
	}

}
