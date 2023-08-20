Simpleshsplit
=============
[![Go Reference](https://pkg.go.dev/badge/github.com/magisterquis/simpleshsplit.svg)](https://pkg.go.dev/github.com/magisterquis/simpleshsplit)

This is a small library for splitting strings on whitespace, such as might be
used for a simple shell.

With [`Split`](https://pkg.go.dev/github.com/magisterquis/simpleshsplit#Split),
the only special character is `\` which is used to escape a
space or another backslash.
[`SplitOn`](https://pkg.go.dev/github.com/magisterquis/simpleshsplit#SplitOn)
allows for changing the escape character.

For splitting but still allowing quoted strings,
[`SplitGoUnquote`](https://pkg.go.dev/github.com/magisterquis/simpleshsplit)
unquotes quoted substrings, similar to shell unquoting but uses
[strconv.Unquote](https://pkg.go.dev/strconv#Unquote) to unquote quoted
portions.

Example
--------
### Simple Splitting
```go
parts := simpleshsplit.Split(`arg1 arg2a\ arg2b arg3a\\arg3b`)
for _, part := range parts {
        fmt.Printf("%v\n", part)
}   
```
Produces
```
arg1
arg2a arg2b
arg3a\arg3b
```

### Go-ish Unquoting
Calling
[`SplitGoUnquote`](https://pkg.go.dev/github.com/magisterquis/simpleshsplit)
on
```
A \"somewhat\" complex string\ with\ a "double-quoted \"\x70art\"" and a
`backtick-quoted \part\` as well.
```
returns as, one split string per line,
```
A
"somewhat"
complex
string with a
double-quoted \"part\"
and
a
backtick-quoted \part\
as
well.
```
