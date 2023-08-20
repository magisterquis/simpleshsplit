package simpleshsplit

/*
 * goquote.go
 * Split into go-quoted strings
 * By J. Stuart McMurray
 * Created 20230820
 * Last Modified 20230820
 */

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitGoUnquote splits s into substrings on spaces.  Strings quoted with
// double-quotes or backticks are unquoted using strconv.Unquote.
// Outside of quoted portions, DefaultEscape may be used to escape
// double-quotes, backticks, spaces, and itself but will cause an error if used
// to escape anything else.
func SplitGoUnquote(s string) ([]string, error) {
	var (
		ret  []string
		cur  strings.Builder /* Current split string. */
		qCur strings.Builder /* Current quoted portion. */
		inDQ bool            /* In a "" section. */
		inBT bool            /* In a `` section. */
		esc  bool            /* Previous character was \. */
	)
	/* wer writes DefaultEscape to be if wesc is true, then writes r. */
	wer := func(b *strings.Builder, wesc bool, r rune) {
		if wesc {
			b.WriteRune(DefaultEscape)
		}
		b.WriteRune(r)
	}
	/* saveSplitString appends the contents of cur to ret and clears cur.
	If cur is empty, saveSplitString is a no-op. */
	saveSplitString := func() {
		if 0 == cur.Len() {
			return
		}
		ret = append(ret, cur.String())
		cur.Reset()
	}
	/* endQuote unquotes and appends the contents of qCur to cur and clears
	qCur.  r is appended to qCur first. */
	endQuote := func(r rune) error {
		qCur.WriteRune(r)
		uq, err := strconv.Unquote(qCur.String())
		if nil != err {
			return fmt.Errorf(
				"unquoting %q: %s",
				qCur.String(),
				err,
			)
		}
		cur.WriteString(uq)
		qCur.Reset()
		return nil
	}
	for _, r := range s {
		switch r {
		case '"':
			switch {
			case inBT:
				/* In `` " isn't special. */
				qCur.WriteRune(r)
			case esc && inDQ:
				/* In quotes, \" means \". */
				wer(&qCur, true, r)
			case esc:
				/* Not in a quote, \" means ". */
				cur.WriteRune(r)
			case inDQ:
				/* End the "". */
				if err := endQuote(r); nil != err {
					return nil, err
				}
				inDQ = false
			default:
				/* Start a "" portion. */
				qCur.WriteRune(r)
				inDQ = true
			}
		case '`':
			switch {
			case inBT:
				/* End the `` portion. */
				if err := endQuote(r); nil != err {
					return nil, err
				}
				inBT = false
			case inDQ:
				/* There's nothing interesting about us in
				double-quotes. */
				wer(&qCur, esc, r)
			case esc:
				/* Outside of quotes, we're not special if
				we're escaped. */
				cur.WriteRune(r)
			default:
				/* Start a `` portion. */
				qCur.WriteRune(r)
				inBT = true
			}
		case DefaultEscape:
			switch {
			case inBT:
				/* \ isn't special in ``. */
				qCur.WriteRune(r)
			case inDQ && esc:
				/* If we're already escaped, we really did mean
				a //. */
				qCur.WriteRune(r)
				qCur.WriteRune(r)
			default:
				esc = true
				continue /* Don't esc = false. */
			}
		case ' ':
			switch {
			case inBT || inDQ:
				/* Inside quotes, we're not special. */
				wer(&qCur, esc, r)
			case esc:
				/* If we're escaped but not in a quote, we're
				just a regular character. */
				cur.WriteRune(r)
			default:
				/* End of a split substring, if it's not a run
				of spaces. */
				saveSplitString()
			}
		default:
			/* If we're in a quote, we always keep the escape
			character. */
			if inBT || inDQ {
				wer(&qCur, esc, r)
				break
			}
			/* If we're escaped and we're here, we don't know this
			escaped character. */
			if esc {
				return nil, fmt.Errorf(
					"Unknown escape %c%c",
					DefaultEscape, r,
				)
			}
			cur.WriteRune(r)
		}
		esc = false
	}

	/* Get the last split string. */
	if 0 != qCur.Len() {
		if inBT {
			return nil, fmt.Errorf("missing terminating `")
		} else if inDQ {
			return nil, fmt.Errorf("missing terminating \"")
		} else {
			return nil, fmt.Errorf(
				"simpleshsplit BUG: missing unknown terminator",
			)
		}
	}
	saveSplitString()

	return ret, nil
}
