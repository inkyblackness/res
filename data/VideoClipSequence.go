package data

import (
	"fmt"

	"github.com/inkyblackness/res/serial"
)

// VideoClipSequenceBaseSize is the amount of bytes a sequence needs at least.
const VideoClipSequenceBaseSize = 16

// VideoClipSequenceEndTag is the constant for the last sequence field.
const VideoClipSequenceEndTag = int16(0x010C)

// VideoClipSequence describes a sequence of a low-res video.
type VideoClipSequence struct {
	Width       int16
	Height      int16
	FramesID    uint16
	Unknown0006 [6]byte
	IntroFlag   int16
	EndTag      int16

	Entries []*VideoClipSequenceEntry
}

// DefaultVideoClipSequence returns a sequence instance with an initialized list of entries.
func DefaultVideoClipSequence(entryCount int) *VideoClipSequence {
	sequence := &VideoClipSequence{
		Entries: make([]*VideoClipSequenceEntry, entryCount),
		EndTag:  VideoClipSequenceEndTag}

	for index := range sequence.Entries {
		sequence.Entries[index] = DefaultVideoClipSequenceEntry()
	}

	return sequence
}

// Code serializes the sequence with the given coder.
func (sequence *VideoClipSequence) Code(coder serial.Coder) {
	coder.Code(&sequence.Width)
	coder.Code(&sequence.Height)
	coder.Code(&sequence.FramesID)
	coder.Code(sequence.Unknown0006)
	coder.Code(&sequence.IntroFlag)
	for _, entry := range sequence.Entries {
		coder.Code(entry)
	}
	coder.Code(&sequence.EndTag)
}

func (sequence *VideoClipSequence) String() (result string) {
	result += fmt.Sprintf("%dx%d, Frames: 0x%04X\n", sequence.Width, sequence.Height, sequence.FramesID)
	result += fmt.Sprintf("IntroFlag: %d, Entries: %d\n", sequence.IntroFlag, len(sequence.Entries))
	for index, entry := range sequence.Entries {
		result += fmt.Sprintf("%d: %v\n", index, entry)
	}

	return result
}

// VideoClipSequenceEntrySize is the amount of bytes a sequence entry has.
const VideoClipSequenceEntrySize = 5

// VideoClipSequenceEntryTag is the constant for the first member.
const VideoClipSequenceEntryTag = byte(0x04)

// VideoClipSequenceEntry describes an entry of a video clip sequence.
type VideoClipSequenceEntry struct {
	Tag        byte
	FirstFrame byte
	LastFrame  byte
	FrameTime  uint16
}

// DefaultVideoClipSequenceEntry returns a new instance of an entry
func DefaultVideoClipSequenceEntry() *VideoClipSequenceEntry {
	entry := &VideoClipSequenceEntry{Tag: VideoClipSequenceEntryTag}

	return entry
}

func (entry *VideoClipSequenceEntry) String() (result string) {
	result += fmt.Sprintf("%02d - %02d: frame time %d",
		entry.FirstFrame, entry.LastFrame, entry.FrameTime)

	return
}
