package main

import "text/template"

var constsTemplate = template.Must(template.New("").Parse(`
// Code generated by yasupGen; DO NOT EDIT.

package {{.FilePackageSimple}}

var (
	ErrIndexOutOfBounds = yasupError{"Index out of bounds"}
	ErrEmptySlice       = yasupError{"The slice is empty but the operation requires at least one element on it"}
)

type yasupError struct {
	msg string
}

func (e yasupError) Error() string {
	return e.msg
}
`))