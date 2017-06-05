package data

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/text"
)

// ElectronicMessage describes one message.
type ElectronicMessage struct {
	nextMessage  int
	isInterrupt  bool
	colorIndex   int
	leftDisplay  int
	rightDisplay int

	title       string
	sender      string
	subject     string
	verboseText string
	terseText   string
}

// NewElectronicMessage returns a new instance of an electronic message.
func NewElectronicMessage() *ElectronicMessage {
	message := &ElectronicMessage{
		nextMessage:  -1,
		colorIndex:   -1,
		leftDisplay:  -1,
		rightDisplay: -1}

	return message
}

// DecodeElectronicMessage tries to decode a message from given block holder.
func DecodeElectronicMessage(cp text.Codepage, holder chunk.BlockHolder) (message *ElectronicMessage, err error) {
	return
}

// Encode serializes the message into a block holder.
func (message *ElectronicMessage) Encode(cp text.Codepage) chunk.BlockHolder {
	blocks := [][]byte{}
	meta := ""

	blocks = append(blocks, cp.Encode(meta))
	blocks = append(blocks, cp.Encode(message.title))
	blocks = append(blocks, cp.Encode(message.sender))
	blocks = append(blocks, cp.Encode(message.subject))
	blocks = append(blocks, cp.Encode(message.verboseText))
	blocks = append(blocks, []byte{0x00})
	blocks = append(blocks, cp.Encode(message.terseText))
	blocks = append(blocks, []byte{0x00})

	return chunk.NewBlockHolder(chunk.BasicChunkType.WithDirectory(), res.Text, blocks)
}

// NextMessage returns the identifier of an interrupting message. Or -1 if no interrupt.
func (message *ElectronicMessage) NextMessage() int {
	return message.nextMessage
}

// SetNextMessage sets the identifier of the interrupting message. -1 to have no interrupt.
func (message *ElectronicMessage) SetNextMessage(id int) {
	message.nextMessage = id
}

// IsInterrupt returns true if this message is an interrupt of another.
func (message *ElectronicMessage) IsInterrupt() bool {
	return message.isInterrupt
}

// SetInterrupt sets whether the message shall be an interrupting message.
func (message *ElectronicMessage) SetInterrupt(value bool) {
	message.isInterrupt = value
}

// ColorIndex returns the color index for the header text. -1 for default color.
func (message *ElectronicMessage) ColorIndex() int {
	return message.colorIndex
}

// SetColorIndex sets the color index for the header text. -1 for default color.
func (message *ElectronicMessage) SetColorIndex(value int) {
	message.colorIndex = value
}

// LeftDisplay returns the display index for the left side. -1 for no display.
func (message *ElectronicMessage) LeftDisplay() int {
	return message.leftDisplay
}

// SetLeftDisplay sets the display index for the left side. -1 for no display.
func (message *ElectronicMessage) SetLeftDisplay(value int) {
	message.leftDisplay = value
}

// RightDisplay returns the display index for the right side. -1 for no display.
func (message *ElectronicMessage) RightDisplay() int {
	return message.rightDisplay
}

// SetRightDisplay sets the display index for the right side. -1 for no display.
func (message *ElectronicMessage) SetRightDisplay(value int) {
	message.rightDisplay = value
}

// Title returns the title of the message.
func (message *ElectronicMessage) Title() string {
	return message.title
}

// SetTitle sets the title of the message.
func (message *ElectronicMessage) SetTitle(value string) {
	message.title = value
}

// Sender returns the sender of the message.
func (message *ElectronicMessage) Sender() string {
	return message.sender
}

// SetSender sets the sender of the message.
func (message *ElectronicMessage) SetSender(value string) {
	message.sender = value
}

// Subject returns the subject of the message.
func (message *ElectronicMessage) Subject() string {
	return message.subject
}

// SetSubject sets the subject of the message.
func (message *ElectronicMessage) SetSubject(value string) {
	message.subject = value
}

// VerboseText returns the verbose text of the message.
func (message *ElectronicMessage) VerboseText() string {
	return message.verboseText
}

// SetVerboseText sets the verbose text of the message.
func (message *ElectronicMessage) SetVerboseText(value string) {
	message.verboseText = value
}

// TerseText returns the terse text of the message.
func (message *ElectronicMessage) TerseText() string {
	return message.terseText
}

// SetTerseText sets the terse text of the message.
func (message *ElectronicMessage) SetTerseText(value string) {
	message.terseText = value
}
