Simpleshsplit
=============
[![GoDoc](https://godoc.org/github.com/magisterquis/simpleshsplit?status.svg)](https://godoc.org/github.com/magisterquis/simpleshsplit)

This is a small library for splitting strings on whitespace, such as might be
used for a simple shell.  By default, the only special character is `\`, which
is used to escape a space or another backslash.

Splitting on a character other than `\` is supported as well.

Example
--------
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
