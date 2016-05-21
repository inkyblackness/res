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
}

func (suite *LevelObjectchaInSuite) TestInitializeResetsAllFields(c *check.C) {
	suite.chain.InitializeLevelObjectChain(len(suite.entries) - 1)

	c.Check(suite.entries[0], check.DeepEquals, data.LevelObjectPrefix{
		LevelObjectTableIndex: 0,
		Next:     0,
		Previous: 1})

	c.Check(suite.entries[1].Previous, check.Equals, uint16(2))
	c.Check(suite.entries[2].Previous, check.Equals, uint16(3))
	c.Check(suite.entries[3].Previous, check.Equals, uint16(0))
}
