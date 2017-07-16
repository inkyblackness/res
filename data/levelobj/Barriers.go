package levelobj

import (
	"fmt"

	"github.com/inkyblackness/res/data/interpreters"
)

var baseBarrier = interpreters.New().
	With("LockVariableIndex", 0, 2).As(interpreters.RangedValue(0, 0x1FF)).
	With("LockMessageIndex", 2, 1).
	With("ForceDoorColor", 3, 1).
	With("RequiredAccessLevel", 4, 1).
	With("AutoCloseTime", 5, 1).As(interpreters.FormattedRangedValue(0, 255,
	func(value int64) string {
		return fmt.Sprintf("%.2f sec  - raw: %d", float64(value)*0.5, value)
	})).
	With("OtherObjectIndex", 6, 2).As(interpreters.ObjectIndex())

func initBarriers() interpreterRetriever {
	return newInterpreterLeaf(baseBarrier)
}
