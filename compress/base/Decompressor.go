package base

import (
	"io"

	"github.com/inkyblackness/res/serial"
)

type decompressor struct {
	reader *wordReader

	isEndOfStream bool
	dictionary    map[word][]byte
	lastBlock     []byte

	leftover []byte
}

// NewDecompressor creates a new decompressor instance over a decoder.
func NewDecompressor(coder serial.Coder) io.Reader {
	obj := &decompressor{
		reader:     newWordReader(coder),
		dictionary: make(map[word][]byte)}

	return obj
}

func (obj *decompressor) Read(p []byte) (n int, err error) {
	requested := len(p)

	for n < requested && !obj.isEndOfStream {
		n += obj.takeFromLeftover(p, n)
		if n < requested {
			obj.readNextWord()
			n += obj.takeFromLeftover(p, n)
		}
	}

	return
}

func (obj *decompressor) takeFromLeftover(dest []byte, destOffset int) (provided int) {
	requested := len(dest) - destOffset
	available := len(obj.leftover)

	if available > 0 && requested > 0 {
		provided = available
		if provided > requested {
			provided = requested
		}
		copy(dest[destOffset:destOffset+provided], obj.leftover)
		obj.leftover = obj.leftover[provided:]
	}

	return
}

func (obj *decompressor) readNextWord() {
	nextWord := obj.reader.read()

	if nextWord == endOfStream {
		obj.isEndOfStream = true
	} else if nextWord == reset {
		obj.dictionary = make(map[word][]byte)
		obj.lastBlock = nil
	} else if len(obj.lastBlock) == 0 {
		obj.lastBlock = []byte{byte(nextWord)}
		obj.leftover = obj.lastBlock
	} else {
		if nextWord < literalLimit {
			obj.leftover = []byte{byte(nextWord)}
		} else if existingBlock, existing := obj.dictionary[nextWord]; existing {
			obj.leftover = existingBlock
		} else {
			obj.leftover = append(obj.lastBlock, obj.lastBlock[0])
		}

		newWord := literalLimit + word(len(obj.dictionary))
		obj.dictionary[newWord] = append(obj.lastBlock, obj.leftover[0])

		obj.lastBlock = obj.leftover
	}
}
