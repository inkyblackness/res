package resfile

import "io"

type singleBlockChunkReader struct {
	contentType ContentType
	compressed  bool
	source      io.Reader
}

func (reader *singleBlockChunkReader) Fragmented() bool {
	return false
}

func (reader *singleBlockChunkReader) ContentType() ContentType {
	return reader.contentType
}

func (reader *singleBlockChunkReader) Compressed() bool {
	return reader.compressed
}

func (reader *singleBlockChunkReader) BlockCount() int {
	return 1
}

func (reader *singleBlockChunkReader) Block(index int) io.Reader {
	return reader.source
}
