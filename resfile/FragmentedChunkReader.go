package resfile

import "io"

type fragmentedChunkReader struct {
	contentType ContentType
	compressed  bool
}

func (reader *fragmentedChunkReader) Fragmented() bool {
	return true
}

func (reader *fragmentedChunkReader) ContentType() ContentType {
	return reader.contentType
}

func (reader *fragmentedChunkReader) Compressed() bool {
	return reader.compressed
}

func (reader *fragmentedChunkReader) BlockCount() int {
	return 1
}

func (reader *fragmentedChunkReader) Block(index int) io.Reader {
	return nil
}
