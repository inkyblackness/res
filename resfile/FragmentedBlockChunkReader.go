package resfile

import "io"

type fragmentedBlockChunkReader struct {
	contentType ContentType
	compressed  bool
}

func (reader *fragmentedBlockChunkReader) Fragmented() bool {
	return true
}

func (reader *fragmentedBlockChunkReader) ContentType() ContentType {
	return reader.contentType
}

func (reader *fragmentedBlockChunkReader) Compressed() bool {
	return reader.compressed
}

func (reader *fragmentedBlockChunkReader) BlockCount() int {
	return 1
}

func (reader *fragmentedBlockChunkReader) Block(index int) io.Reader {
	return nil
}
