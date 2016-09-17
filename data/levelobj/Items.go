package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
)

var baseItem = interpreters.New()

func initItems() interpreterRetriever {

	class := newInterpreterEntry(baseItem)
	//class.set(0, electronics)

	return class
}
