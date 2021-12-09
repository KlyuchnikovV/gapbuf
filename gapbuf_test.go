package gapbuf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	const (
		Insert = iota
		Delete
	)

	testCases := []struct {
		desc   string
		init   string
		action int
		cursor int
		str    string
		result string
	}{
		{
			desc:   "Insert",
			init:   "Interrption",
			action: Insert,
			cursor: 6,
			str:    "u",
			result: "Interruption",
		},
		{
			desc:   "Delete",
			init:   "Interruption",
			action: Delete,
			cursor: 7,
			result: "Interrption",
		},
	}

	for _, tC := range testCases {
		buffer := New([]byte(tC.init)...)
		t.Run(tC.desc, func(t *testing.T) {
			switch tC.action {
			case Insert:
				buffer.SetCursor(tC.cursor)
				buffer.Insert([]byte(tC.str)...)
			case Delete:
				buffer.SetCursor(tC.cursor)
				buffer.Delete()
			}
			assert.Equal(t, tC.result, string(buffer.Bytes()))

			fmt.Printf("Result %s\n", buffer.Bytes())
		})
	}
}
