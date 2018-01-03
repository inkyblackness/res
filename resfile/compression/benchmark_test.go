package compression

import (
	"math/rand"
	"testing"

	"github.com/inkyblackness/res/serial"
	"io"
)

func rawData() []byte {
	data := make([]byte, 1024*1024)
	rand.Read(data)
	return data
}

func BenchmarkRawDataStorage(b *testing.B) {
	data := rawData()
	b.ResetTimer()
	for run := 0; run < b.N; run++ {
		encoder := serial.NewEncoder(serial.NewByteStore())
		encoder.Write(data)
	}
}

func BenchmarkEncoding(b *testing.B) {
	data := rawData()
	b.ResetTimer()
	for run := 0; run < b.N; run++ {
		compressor := NewCompressor(serial.NewByteStore())
		compressor.Write(data)
		compressor.Close()
	}
}

func BenchmarkEncodingDecoding(b *testing.B) {
	data := rawData()
	output := make([]byte, len(data))
	b.ResetTimer()
	for run := 0; run < b.N; run++ {
		store := serial.NewByteStore()
		compressor := NewCompressor(store)
		compressor.Write(data)
		compressor.Close()
		store.Seek(0, io.SeekStart)
		decompressor := NewDecompressor(store)
		decompressor.Read(output)
	}
}
