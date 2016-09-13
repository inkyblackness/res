package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
)

var baseScenery = interpreters.New()

var words = interpreters.New().
	With("TextIndex", 0, 2).
	With("FontAndSize", 2, 1).
	With("Color", 4, 1)
