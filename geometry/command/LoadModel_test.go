package command

import (
	"bytes"

	"github.com/inkyblackness/res/geometry"

	check "gopkg.in/check.v1"
)

type LoadModelSuite struct {
	nodeAnchors []geometry.NodeAnchor
	nodes       []geometry.Node
}

var _ = check.Suite(&LoadModelSuite{})

func (suite *LoadModelSuite) SetUpTest(c *check.C) {
	suite.nodeAnchors = nil
	suite.nodes = nil
}

func (suite *LoadModelSuite) TestLoadModelReturnsErrorOnNil(c *check.C) {
	_, err := LoadModel(nil)

	c.Check(err, check.ErrorMatches, "source is nil")
}

func (suite *LoadModelSuite) TestLoadModelReturnsModelInstanceOnValidData(c *check.C) {
	source := bytes.NewReader(suite.aSimpleList())
	model, err := LoadModel(source)

	c.Assert(err, check.IsNil)
	c.Check(model, check.NotNil)
}

func (suite *LoadModelSuite) TestModelContainsSingleVertex(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
	}))
	model, err := LoadModel(source)

	c.Assert(err, check.IsNil)
	c.Check(model.VertexCount(), check.Equals, 1)
}

func (suite *LoadModelSuite) TestModelContainsMultipleVertices(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertices([]Vector{Vector{1, 0, 0}, Vector{0, 2, 0}, Vector{0, 0, 3}})
	}))
	model, err := LoadModel(source)

	c.Assert(err, check.IsNil)
	c.Check(model.VertexCount(), check.Equals, 3)
}

func (suite *LoadModelSuite) TestModelContainsSingleOffsetVertices(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
		writer.WriteDefineOneOffsetVertex(CmdDefineOffsetVertexX, 1, 0, 1.0)
		writer.WriteDefineOneOffsetVertex(CmdDefineOffsetVertexY, 2, 0, 1.0)
		writer.WriteDefineOneOffsetVertex(CmdDefineOffsetVertexZ, 3, 0, 1.0)
	}))
	model, err := LoadModel(source)

	c.Assert(err, check.IsNil)
	c.Check(model.VertexCount(), check.Equals, 4)
}

func (suite *LoadModelSuite) TestLoadModelReturnsErrorForSingleOffsetVertexWhenNewIndexIsNotEqualCurrentCount(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
		writer.WriteDefineOneOffsetVertex(CmdDefineOffsetVertexX, 2, 0, 1.0)
	}))
	_, err := LoadModel(source)

	c.Check(err, check.ErrorMatches, "Offset vertex uses invalid newIndex.*")
}

func (suite *LoadModelSuite) TestModelContainsDoubleOffsetVertices(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
		writer.WriteDefineTwoOffsetVertex(CmdDefineOffsetVertexXY, 1, 0, 1.0, 2.0)
		writer.WriteDefineTwoOffsetVertex(CmdDefineOffsetVertexXZ, 2, 0, 1.0, 2.0)
		writer.WriteDefineTwoOffsetVertex(CmdDefineOffsetVertexYZ, 3, 0, 1.0, 2.0)
	}))
	model, err := LoadModel(source)

	c.Assert(err, check.IsNil)
	c.Check(model.VertexCount(), check.Equals, 4)
}

func (suite *LoadModelSuite) TestLoadModelReturnsErrorForDoubleOffsetVertexWhenNewIndexIsNotEqualCurrentCount(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
		writer.WriteDefineTwoOffsetVertex(CmdDefineOffsetVertexXY, 2, 0, 1.0, 2.0)
	}))
	_, err := LoadModel(source)

	c.Check(err, check.ErrorMatches, "Offset vertex uses invalid newIndex.*")
}

func (suite *LoadModelSuite) TestModelCanHaveANodeAnchor(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
		writer.WriteNodeAnchor(Vector{1, 0, 0}, Vector{0, 0, 0}, 2, 4)
		writer.WriteEndOfNode() // end of root node
		writer.WriteEndOfNode() // end of left node
	}))
	model, err := LoadModel(source)

	c.Assert(err, check.IsNil)

	model.WalkAnchors(suite)
	c.Check(len(suite.nodeAnchors), check.Equals, 1)
}

func (suite *LoadModelSuite) TestLoadModelReturnsErrorForInvalidNodeAnchorOffsets(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
		writer.WriteNodeAnchor(Vector{1, 0, 0}, Vector{0, 0, 0}, 2, 5)
		writer.WriteEndOfNode() // end of root node
		writer.WriteEndOfNode() // end of left node
	}))
	_, err := LoadModel(source)

	c.Check(err, check.ErrorMatches, "Wrong offset values for node anchor")
}

func (suite *LoadModelSuite) TestModelCanHaveNestedNodeAnchors(c *check.C) {
	source := bytes.NewReader(suite.anEmptyModelWith(func(writer *Writer) {
		writer.WriteDefineVertex(Vector{0, 0, 0})
		// root node
		writer.WriteNodeAnchor(Vector{1, 0, 0}, Vector{0, 0, 0}, 2, 2+cmdDefineNodeAnchorSize+2+2+2)
		writer.WriteEndOfNode() // end of root node
		// root.left node
		writer.WriteNodeAnchor(Vector{1, 0, 0}, Vector{0, 0, 0}, 4, 2)
		writer.WriteEndOfNode() // end of root.left node
		// root.left.right node
		writer.WriteEndOfNode() // end of root.left.right node
		// root.left.left node
		writer.WriteEndOfNode() // end of root.left.left node
		// root.right node (end provided by helper function)
	}))
	model, err := LoadModel(source)

	c.Assert(err, check.IsNil)

	model.WalkAnchors(suite)
	c.Check(len(suite.nodeAnchors), check.Equals, 2)
	c.Check(len(suite.nodes), check.Equals, 4)
}

func (suite *LoadModelSuite) Nodes(anchor geometry.NodeAnchor) {
	suite.nodeAnchors = append(suite.nodeAnchors, anchor)
	suite.nodes = append(suite.nodes, anchor.Left(), anchor.Right())
	anchor.Left().WalkAnchors(suite)
	anchor.Right().WalkAnchors(suite)
}

func (suite *LoadModelSuite) Faces(anchor geometry.FaceAnchor) {

}

func (suite *LoadModelSuite) aSimpleList() []byte {
	writer := NewWriter()

	writer.WriteHeader(0)
	writer.WriteEndOfNode()

	return writer.Bytes()
}

func (suite *LoadModelSuite) anEmptyModelWith(vertices func(writer *Writer)) []byte {
	writer := NewWriter()

	writer.WriteHeader(0)
	vertices(writer)
	writer.WriteEndOfNode()

	return writer.Bytes()
}
