package resfile

import "io"

type singleBlockChunkReader struct {
}

func (reader *singleBlockChunkReader) Fragmented() bool {
	return false
}

func (reader *singleBlockChunkReader) ContentType() ContentType {
	return ContentType(0xFF)
}

func (reader *singleBlockChunkReader) Compressed() bool {
	return false
}

func (reader *singleBlockChunkReader) BlockCount() int {
	return 1
}

func (reader *singleBlockChunkReader) Block(index int) io.Reader {
	return nil
}
