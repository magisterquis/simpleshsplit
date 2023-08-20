package simpleshsplit

/*
 * goquote_test.go
 * Tests for goquote.go
 * By J. Stuart McMurray
 * Created 20230820
 * Last Modified 20230820
 */

import (
	"fmt"
	"slices"
	"testing"
)

func TestSplitGoUnuote(t *testing.T) {
	for have, want := range map[string][]string{
		"foo":                              {"foo"},
		"foo bar":                          {"foo", "bar"},
		`foo "bar tridge" baaz`:            {"foo", "bar tridge", "baaz"},
		`foo bar"tridge"baaz quux`:         {"foo", "bartridgebaaz", "quux"},
		"foo `bar` tridge":                 {"foo", "bar", "tridge"},
		"foo `bar`":                        {"foo", "bar"},
		"foo `b\\ar`":                      {"foo", `b\ar`},
		`foo "bar"`:                        {"foo", "bar"},
		"foo `A\\x42C` bar":                {"foo", `A\x42C`, "bar"},
		`foo "A\x42C" tridge`:              {"foo", "ABC", "tridge"},
		`\"`:                               {`"`},
		`"foo"`:                            {"foo"},
		`\"foo\"`:                          {`"foo"`},
		`foo \"bar"tridge\"\" "baaz quux`:  {"foo", `"bartridge"" baaz`, "quux"},
		`"and \\'s are hard"`:              {`and \'s are hard`},
		`ls -l "spaces and \\'s are hard"`: {"ls", "-l", "spaces and \\'s are hard"},
		`A \"somewhat\" complex string\ with\ a "double-quoted ` +
			`\"\x70art\"" and ` + "a `backtick-quoted \\part\\` " +
			"as well.": {
			"A", `"somewhat"`, "complex", "string with a",
			"double-quoted \"part\"", "and", "a",
			`backtick-quoted \part\`, "as", "well.",
		},
	} {
		have := have /* :( */
		want := want /* :( */
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got, err := SplitGoUnquote(have)
			if nil != err {
				t.Errorf("Error : %s\nHave: %q", err, have)
				return
			}
			if !slices.Equal(got, want) {
				t.Errorf(
					"Split failed\n"+
						"Have: %q\n"+
						"Got: %q\n"+
						"Want: %q",
					have,
					got,
					want,
				)
			}
		})
	}
}

func ExampleSplitGoUnquote() {
	/* Split a command */
	parts, err := SplitGoUnquote(
		`echo -n "spaces and \\'s are hard" ` +
			"but `backtick quotes are easy`",
	)
	if nil != err {
		panic("SplitGoUnquote: " + err.Error())
	}
	for _, part := range parts {
		fmt.Printf("%s\n", part)
	}

	//Output:
	// echo
	// -n
	// spaces and \'s are hard
	// but
	// backtick quotes are easy
}
