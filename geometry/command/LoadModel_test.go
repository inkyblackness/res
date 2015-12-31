package command

import (
	"bytes"

	check "gopkg.in/check.v1"
)

type LoadModelSuite struct {
}

var _ = check.Suite(&LoadModelSuite{})

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
