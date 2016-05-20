package logic

import (
	"github.com/inkyblackness/res/data"

	check "gopkg.in/check.v1"
)

type CrossReferenceListSuite struct {
	referencer *TestingTileMapReferencer
}

var _ = check.Suite(&CrossReferenceListSuite{})

func (suite *CrossReferenceListSuite) SetUpTest(c *check.C) {
	suite.referencer = NewTestingTileMapReferencer()
}

func (suite *CrossReferenceListSuite) aListOfSize(size int) *CrossReferenceList {
	references := make([]data.LevelObjectCrossReference, size)
	list := &CrossReferenceList{references: references}

	return list
}

func (suite *CrossReferenceListSuite) aClearListOfSize(size int) *CrossReferenceList {
	list := suite.aListOfSize(size)

	list.Clear()

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

func (suite *CrossReferenceListSuite) TestAddObjectToMapReturnsAnIndex(c *check.C) {
	list := suite.aClearListOfSize(4)
	locations := []TileLocation{AtTile(1, 1)}

	index := list.AddObjectToMap(1, suite.referencer, locations)

	c.Check(index, check.Not(check.Equals), CrossReferenceListIndex(0))
}

func (suite *CrossReferenceListSuite) TestAddObjectToMapRegistersIndicesAtMap(c *check.C) {
	list := suite.aClearListOfSize(4)
	location1 := AtTile(1, 1)
	location2 := AtTile(2, 2)
	locations := []TileLocation{location1, location2}

	list.AddObjectToMap(1, suite.referencer, locations)

	c.Check(suite.referencer.ReferenceIndex(location1), check.Not(check.Equals), CrossReferenceListIndex(0))
	c.Check(suite.referencer.ReferenceIndex(location2), check.Not(check.Equals), CrossReferenceListIndex(0))
	c.Check(suite.referencer.ReferenceIndex(location1), check.Not(check.Equals), suite.referencer.ReferenceIndex(location2))
}

func (suite *CrossReferenceListSuite) TestAddObjectToMapSetsPropertiesOfSingleEntry(c *check.C) {
	list := suite.aClearListOfSize(4)
	location1 := AtTile(7, 8)
	locations := []TileLocation{location1}

	index := list.AddObjectToMap(20, suite.referencer, locations)

	firstEntry := list.entry(index)
	c.Check(firstEntry, check.DeepEquals, &data.LevelObjectCrossReference{
		LevelObjectTableIndex: 20,
		NextObjectIndex:       0,
		NextTileIndex:         uint16(index),
		TileX:                 7,
		TileY:                 8})
}

func (suite *CrossReferenceListSuite) TestAddObjectToMapSetsPropertiesOfMultipleEntries(c *check.C) {
	list := suite.aClearListOfSize(4)
	location1 := AtTile(3, 4)
	location2 := AtTile(5, 6)
	locations := []TileLocation{location1, location2}

	index := list.AddObjectToMap(10, suite.referencer, locations)

	firstEntry := list.entry(index)
	c.Check(firstEntry, check.DeepEquals, &data.LevelObjectCrossReference{
		LevelObjectTableIndex: 10,
		NextObjectIndex:       0,
		NextTileIndex:         uint16(index - 1),
		TileX:                 5,
		TileY:                 6})

	secondEntry := list.entry(index - 1)
	c.Check(secondEntry, check.DeepEquals, &data.LevelObjectCrossReference{
		LevelObjectTableIndex: 10,
		NextObjectIndex:       0,
		NextTileIndex:         uint16(index),
		TileX:                 3,
		TileY:                 4})
}

func (suite *CrossReferenceListSuite) TestAddObjectToMapKeepsReferencesOfObjectsInSameTile(c *check.C) {
	list := suite.aClearListOfSize(10)
	location1 := AtTile(16, 12)
	location2 := AtTile(1, 2)

	list.AddObjectToMap(40, suite.referencer, []TileLocation{AtTile(100, 100)})
	existingIndex := list.AddObjectToMap(50, suite.referencer, []TileLocation{location1})
	index := list.AddObjectToMap(60, suite.referencer, []TileLocation{location1, location2})

	firstEntry := list.entry(index)
	c.Check(firstEntry, check.DeepEquals, &data.LevelObjectCrossReference{
		LevelObjectTableIndex: 60,
		NextObjectIndex:       0,
		NextTileIndex:         uint16(index - 1),
		TileX:                 1,
		TileY:                 2})

	secondEntry := list.entry(index - 1)
	c.Check(secondEntry, check.DeepEquals, &data.LevelObjectCrossReference{
		LevelObjectTableIndex: 60,
		NextObjectIndex:       uint16(existingIndex),
		NextTileIndex:         uint16(index),
		TileX:                 16,
		TileY:                 12})
}
