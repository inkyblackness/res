package store

import (
	"testing"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"

	check "gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

func emptyBlockHolder() chunk.BlockHolder {
	return chunk.NewBlockHolder(chunk.BasicChunkType, res.Palette, [][]byte{nil})
}
