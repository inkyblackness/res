package base

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/inkyblackness/res/serial"
)

type compressor struct {
	writer *wordWriter

	chain      []byte
	dictionary map[string]word
	bestWord   word
	overtime   int
}

// NewCompressor creates a new compressor instance over an encoder.
func NewCompressor(coder serial.Coder) io.WriteCloser {
	obj := &compressor{
		writer:     newWordWriter(coder),
		dictionary: make(map[string]word),
		bestWord:   reset,
		overtime:   0}

	return obj
}

func (obj *compressor) Close() error {
	if obj.bestWord != reset {
		obj.writer.write(obj.bestWord)
	}
	obj.writer.close()

	return nil
}

func (obj *compressor) Write(p []byte) (n int, err error) {
	n = len(p)

	for _, input := range p {
		obj.chain = append(obj.chain, input)
		key := createKey(obj.chain)

		existingWord, ok := obj.dictionary[key]
		if ok {
			obj.bestWord = existingWord
		} else {
			if obj.bestWord != reset {
				newWord := literalLimit + word(len(obj.dictionary))

				obj.writer.write(obj.bestWord)
				if newWord < reset {
					obj.dictionary[key] = newWord
				} else {
					obj.overtime++
					if obj.overtime >= 1000 {
						obj.writer.write(reset)
						obj.dictionary = make(map[string]word)
					}
				}
			}

			obj.bestWord = word(input)
			obj.chain = []byte{input}
		}
	}

	return
}

func createKey(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
