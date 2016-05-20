package logic

import (
	"github.com/inkyblackness/res/data"

	check "gopkg.in/check.v1"
)

type CrossReferenceListSuite struct {
}

var _ = check.Suite(&CrossReferenceListSuite{})

func (suite *CrossReferenceListSuite) aListOfSize(size int) *CrossReferenceList {
	references := make([]data.LevelObjectCrossReference, size)
	list := &CrossReferenceList{references: references}

	return list
}

func (suite *CrossReferenceListSuite) TestNewCrossReferenceListReturnsListWithASizeOf1600(c *check.C) {
	list := NewCrossReferenceList()

	c.Check(list.size(), check.Equals, 1600)
}

func (suite *CrossReferenceListSuite) TestEncodeReturnsExpectedAmountOfBytes(c *check.C) {
	list := suite.aListOfSize(5)

	bytes := list.Encode()

	c.Check(len(bytes), check.Equals, data.LevelObjectCrossReferenceSize*5)
}

func (suite *CrossReferenceListSuite) TestEncodeSerializesAccordingToFormat(c *check.C) {
	list := suite.aListOfSize(1)

	entry0 := list.entry(0)
	entry0.TileX = 0x0123
	entry0.TileY = 0x4567
	entry0.LevelObjectTableIndex = 0x89AB
	entry0.NextObjectIndex = 0xCDEF
	entry0.NextTileIndex = 0x0011

	bytes := list.Encode()

	c.Check(bytes, check.DeepEquals, []byte{0x23, 0x01, 0x67, 0x45, 0xAB, 0x89, 0xEF, 0xCD, 0x11, 0x00})
}

func (suite *CrossReferenceListSuite) TestClearResetsTheList(c *check.C) {
	list := suite.aListOfSize(3)

	list.Clear()
	bytes := list.Encode()

	c.Check(bytes, check.DeepEquals, []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0,
		0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0,
		0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
}
