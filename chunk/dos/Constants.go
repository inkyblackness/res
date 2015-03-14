package dos

const (
	// HeaderString is expected at the start of a resource file
	HeaderString = "LG Res File v2\r\n"
	// CommentTerminator marks the end of a comment after the header string.
	CommentTerminator = byte(0x1A)
	// ChunkDirectoryFileOffsetPos is the position of the file offset value to the chunk directory
	ChunkDirectoryFileOffsetPos = 0x7C
)
