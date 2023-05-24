// Package simpleshsplit splits strings on whitespace, allowing escaped spaces
package simpleshsplit

/*
 * simpleshsplit.go
 * Split shell commands, very simply
 * By J. Stuart McMurray
 * Created 20180803
 * Last Modified 20230524
 */

// Split splits b into space-separated substrings.  A space may be escaped by
// preceeding it with a backslash.  A backslash may be escaped in the same way.
func Split(s string) []string {
	var (
		p  bool /* Previous character was a backslash */
		w  []rune
		ss []string
	)
	for _, r := range s {
		switch r {
		case '\\':
			/* If this is the first backslash in a pair note it */
			if !p {
				p = true
				continue
			}
			/* This one was escaped */
			w = append(w, r)
		case ' ':
			/* If this was escaped, add it to the word */
			if p {
				w = append(w, r)
				break
			}
			/* If not, we're not in a word */
			if 0 != len(w) {
				ss = append(ss, string(w))
				w = w[:0]
			}
		default:
			/* Other characters get appended to the current word */
			if r != ' ' && r != '\\' {
				w = append(w, r)
			}
		}
		p = false
	}
	/* Add the final word, if we have one */
	if 0 != len(w) {
		ss = append(ss, string(w))
	}
	return ss
}
