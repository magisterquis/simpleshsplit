package simpleshsplit

/*
 * simpleshsplit_test.go
 * Tests for simpleshsplit
 * By J. Stuart McMurray
 * Created 20180803
 * Last Modified 20180803
 */

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	for have, want := range map[string][]string{
		`a b`:               {`a`, `b`},
		`a\ b c\  \ d   e`:  {`a b`, `c `, ` d`, `e`},
		`a\\b c\\\ d e\\ f`: {`a\b`, `c\ d`, `e\`, `f`},
	} {
		if got := Split(have); !reflect.DeepEqual(got, want) {
			t.Errorf("have:%q got:%q want:%q", have, got, want)
		}
	}
}

func ExampleSplit() {
	parts := Split(`arg1 arg2a\ arg2b arg3a\\arg3b`)
	for _, part := range parts {
		fmt.Printf("%v\n", part)
	}

	// Output:
	//
	// arg1
	// arg2a arg2b
	// arg3a\arg3b
}
