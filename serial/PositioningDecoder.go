package serial

import "io"

type positioningDecoder struct {
	Decoder

	seeker io.Seeker
}

// NewPositioningDecoder creates a new Decoder from given reader.
func NewPositioningDecoder(source io.ReadSeeker) PositioningCoder {
	return &positioningDecoder{Decoder: Decoder{source: source, offset: 0}, seeker: source}
}

func (coder *positioningDecoder) CurPos() uint32 {
	return coder.offset
}

func (coder *positioningDecoder) SetCurPos(offset uint32) {
	if coder.firstError != nil {
		return
	}
	_, coder.firstError = coder.seeker.Seek(int64(offset), io.SeekStart)
	if coder.firstError != nil {
		return
	}
	coder.offset = offset
}
