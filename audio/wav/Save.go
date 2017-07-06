package wav

import (
	"encoding/binary"
	"io"
)

// Save encodes the provided samples into the given writer
func Save(writer io.Writer, sampleRate float32, samples []byte) {
	dataSize := uint32(len(samples))
	var fmt formatHeader
	fmtSize := fmt.size()
	contentType := riffContentTypeWave
	contentTypeSize := uint32(4)
	tagSize := uint32(8)
	tagSizes := uint32(tagSize * 2)

	fmt.base.formatType = waveFormatTypePcm
	fmt.base.channels = 1
	fmt.base.samplesPerSec = uint32(sampleRate)
	fmt.base.avgBytesPerSec = fmt.base.samplesPerSec
	fmt.base.blockAlign = 1
	fmt.extension.bitsPerSample = 8

	riffTag := riffChunkTag{riffChunkTypeRiff, tagSizes + contentTypeSize + fmtSize + dataSize}

	binary.Write(writer, binary.LittleEndian, &riffTag)
	binary.Write(writer, binary.LittleEndian, &contentType)
	binary.Write(writer, binary.LittleEndian, &riffChunkTag{riffChunkTypeFmt, fmtSize})
	binary.Write(writer, binary.LittleEndian, &fmt.base)
	binary.Write(writer, binary.LittleEndian, &fmt.extension)
	binary.Write(writer, binary.LittleEndian, &riffChunkTag{riffChunkTypeData, dataSize})
	writer.Write(samples)
}
