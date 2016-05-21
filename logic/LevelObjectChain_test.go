package logic

import (
	"github.com/inkyblackness/res/data"

	check "gopkg.in/check.v1"
)

type LevelObjectchaInSuite struct {
	chain   *LevelObjectChain
	entries [4]data.LevelObjectPrefix
}

var _ = check.Suite(&LevelObjectchaInSuite{})

func (suite *LevelObjectchaInSuite) SetUpTest(c *check.C) {
	linkGetter := func(index data.LevelObjectChainIndex) LevelObjectChainLink {
		return &suite.entries[index]
	}
	suite.chain = NewLevelObjectChain(&suite.entries[data.LevelObjectChainStartIndex], linkGetter)
	suite.chain.InitializeLevelObjectChain(len(suite.entries) - 1)
}

func (suite *LevelObjectchaInSuite) TestInitializeResetsAllFields(c *check.C) {
	c.Check(suite.entries[0], check.DeepEquals, data.LevelObjectPrefix{
		LevelObjectTableIndex: 0,
		Next:     0,
		Previous: 1})

	c.Check(suite.entries[1].Previous, check.Equals, uint16(2))
	c.Check(suite.entries[2].Previous, check.Equals, uint16(3))
	c.Check(suite.entries[3].Previous, check.Equals, uint16(0))
}

func (suite *LevelObjectchaInSuite) TestAcquireLinkReturnsIndexForNewEntry(c *check.C) {
	index, _ := suite.chain.AcquireLink()

	c.Check(index, check.Not(check.Equals), data.LevelObjectChainIndex(0))
}

func (suite *LevelObjectchaInSuite) TestAcquireLinkUpdatesStartWhenEmpty(c *check.C) {
	index, _ := suite.chain.AcquireLink()

	c.Check(suite.entries[0], check.DeepEquals, data.LevelObjectPrefix{
		LevelObjectTableIndex: uint16(index),
		Next:     uint16(index),
		Previous: 2})
}

func (suite *LevelObjectchaInSuite) TestAcquireLinkUpdatesEntries(c *check.C) {
	index1, _ := suite.chain.AcquireLink()
	index2, _ := suite.chain.AcquireLink()

	c.Check(suite.entries[0], check.DeepEquals, data.LevelObjectPrefix{
		LevelObjectTableIndex: uint16(index2),
		Next:     uint16(index1),
		Previous: 3})

	c.Check(suite.entries[index1], check.DeepEquals, data.LevelObjectPrefix{
		LevelObjectTableIndex: 0,
		Next:     uint16(index2),
		Previous: uint16(data.LevelObjectChainStartIndex)})

	c.Check(suite.entries[index2], check.DeepEquals, data.LevelObjectPrefix{
		LevelObjectTableIndex: 0,
		Next:     uint16(data.LevelObjectChainStartIndex),
		Previous: uint16(index1)})
}

func (suite *LevelObjectchaInSuite) TestAcquireReturnsErrorWhenExhausted(c *check.C) {
	suite.chain.AcquireLink()
	suite.chain.AcquireLink()
	suite.chain.AcquireLink()

	_, err := suite.chain.AcquireLink()

	c.Check(err, check.NotNil)
}
