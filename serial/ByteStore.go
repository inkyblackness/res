package serial

// ByteStore is implementing a WriteSeeker storing in memory
type ByteStore struct {
	data   []byte
	offset int
}

// NewByteStore returns a new byte store instance
func NewByteStore() *ByteStore {
	store := &ByteStore{
		data:   make([]byte, 0),
		offset: 0}

	return store
}

// Len returns the current length of the buffer
func (store *ByteStore) Len() int {
	return len(store.data)
}

// Data returns the current data buffer
func (store *ByteStore) Data() []byte {
	return store.data
}

// Seek implements the Seeker interface
func (store *ByteStore) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case 0:
		store.offset = int(offset)
	case 1:
		store.offset += int(offset)
	case 2:
		store.offset = len(store.data) + int(offset)
	}

	return offset, nil
}

// Write implements the Writer interface
func (store *ByteStore) Write(p []byte) (n int, err error) {
	size := len(p)
	store.ensureAvailable(size)
	copy(store.data[store.offset:], p)
	store.offset += size

	return size, nil
}

func (store *ByteStore) Close() error {
	// ignored
	return nil
}

func (store *ByteStore) ensureAvailable(size int) {
	expected := store.offset + size
	available := len(store.data)

	if expected > available {
		old := store.data

		store.data = make([]byte, expected)
		copy(store.data, old)
	}
}
