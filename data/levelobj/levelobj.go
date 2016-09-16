package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
)

var realWorldEntries *interpreterEntry
var cyberspaceEntries *interpreterEntry

func init() {

	projectiles := newInterpreterEntry(baseProjectile)

	hardware := newInterpreterEntry(baseHardware)

	software := newInterpreterEntry(baseSoftware)

	scenery := newInterpreterEntry(baseScenery)

	items := newInterpreterEntry(baseItem)

	panels := newInterpreterEntry(basePanel)

	barriers := newInterpreterEntry(baseBarrier)

	animations := newInterpreterEntry(baseAnimation)

	markers := newInterpreterEntry(baseMarker)

	containers := newInterpreterEntry(baseContainer)

	critters := newInterpreterEntry(baseCritter)

	realWorldEntries = newInterpreterEntry(interpreters.New())
	realWorldEntries.set(0, initWeapons())
	realWorldEntries.set(1, newInterpreterEntry(interpreters.New())) // have no data
	realWorldEntries.set(2, projectiles)
	realWorldEntries.set(3, initExplosives())
	realWorldEntries.set(4, newInterpreterEntry(interpreters.New())) // have no data
	realWorldEntries.set(5, hardware)
	realWorldEntries.set(6, software)
	realWorldEntries.set(7, scenery)
	realWorldEntries.set(8, items)
	realWorldEntries.set(9, panels)
	realWorldEntries.set(10, barriers)
	realWorldEntries.set(11, animations)
	realWorldEntries.set(12, markers)
	realWorldEntries.set(13, containers)
	realWorldEntries.set(14, critters)

	cyberspaceEntries = newInterpreterEntry(interpreters.New())
}
