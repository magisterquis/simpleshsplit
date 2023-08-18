package simpleshsplit

/*
 * simpleshsplit_test.go
 * Tests for simpleshsplit
 * By J. Stuart McMurray
 * Created 20180803
 * Last Modified 20230818
 */

import (
	"fmt"
	"reflect"
	"slices"
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

func TestSplitOn(t *testing.T) {
	for _, c := range []struct {
		have string
		esc  rune
		want []string
	}{{
		have: `a b`,
		esc:  '^',
		want: []string{`a`, `b`},
	}, {
		have: `a^ b c^  ^ d   e`,
		esc:  '^',
		want: []string{`a b`, `c `, ` d`, `e`},
	}, {
		have: `a^^b c^^^ d e^^ f`,
		esc:  '^',
		want: []string{`a^b`, `c^ d`, `e^`, `f`},
	}, {
		have: `foo^ bar\ tridge`,
		esc:  '^',
		want: []string{`foo bar\`, `tridge`},
	}, {
		have: `ls foo\bar^ tridge`,
		esc:  '^',
		want: []string{`ls`, `foo\bar tridge`},
	}} {
		c := c /* :C */
		t.Run(c.have, func(t *testing.T) {
			t.Parallel()
			got := SplitOn(c.have, c.esc)
			if slices.Equal(got, c.want) {
				return
			}
			t.Errorf("got: %q", got)
		})
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
