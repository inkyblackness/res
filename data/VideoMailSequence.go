package data

import (
	"fmt"

	"github.com/inkyblackness/res"
)

// VideoMailSequenceBaseSize is the amount of bytes a sequence needs at least.
const VideoMailSequenceBaseSize = 16

// VideoMailSequenceEndTag is the constant for the last sequence field.
const VideoMailSequenceEndTag = int16(0x010C)

// VideoMailSequence describes a sequence of a video mail.
type VideoMailSequence struct {
	Width       int16
	Height      int16
	FramesID    res.ResourceID
	Unknown0006 [6]byte
	IntroFlag   int16
	Entries     []*VideoMailSequenceEntry
	EndTag      int16
}

// DefaultVideoMailSequence returns a sequence instance with an initialized list of entries.
func DefaultVideoMailSequence(entryCount int) *VideoMailSequence {
	sequence := &VideoMailSequence{
		Entries: make([]*VideoMailSequenceEntry, entryCount),
		EndTag:  VideoMailSequenceEndTag}

	for index := range sequence.Entries {
		sequence.Entries[index] = DefaultVideoMailSequenceEntry()
	}

	return sequence
}

func (sequence *VideoMailSequence) String() (result string) {
	result += fmt.Sprintf("%dx%d, Frames: 0x%04X\n", sequence.Width, sequence.Height, uint16(sequence.FramesID))
	result += fmt.Sprintf("IntroFlag: %d, Entries: %d\n", sequence.IntroFlag, len(sequence.Entries))
	for index, entry := range sequence.Entries {
		result += fmt.Sprintf("%d: %v\n", index, entry)
	}

	return result
}

// VideoMailSequenceEntrySize is the amount of bytes a sequence entry has.
const VideoMailSequenceEntrySize = 5

// VideoMailSequenceEntryTag is the constant for the first member.
const VideoMailSequenceEntryTag = byte(0x04)

// VideoMailSequenceEntry describes an entry of a video mail sequence.
type VideoMailSequenceEntry struct {
	Tag         byte
	FirstFrame  byte
	LastFrame   byte
	Unknown0003 [2]byte
}

// DefaultVideoMailSequenceEntry returns a new instance of an entry
func DefaultVideoMailSequenceEntry() *VideoMailSequenceEntry {
	entry := &VideoMailSequenceEntry{Tag: VideoMailSequenceEntryTag}

	return entry
}

func (entry *VideoMailSequenceEntry) String() (result string) {
	result += fmt.Sprintf("%02d - %02d: 0x%02X 0x%02X (%d)",
		entry.FirstFrame, entry.LastFrame, entry.Unknown0003[0], entry.Unknown0003[1],
		int(entry.Unknown0003[1])<<8+int(entry.Unknown0003[0]))

	return
}
