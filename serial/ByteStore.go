package serial

import (
	"io"
)

// TODO: write tests
// TODO: remove Close functionality
// TODO: remove unused code (e.g.: Len function)

const bufferCapacityIncrement = 1024 * 1024

// ByteStore is implementing a ReadWriteSeeker storing data in memory.
type ByteStore struct {
	data    []byte
	offset  int
	onClose func([]byte)
}

// NewByteStore returns a new byte store instance
func NewByteStore() *ByteStore {
	return NewByteStoreFromData(make([]byte, 0, bufferCapacityIncrement), func([]byte) {})
}

// NewByteStoreFromData returns a new ByteStore instance wrapping the provided data.
func NewByteStoreFromData(data []byte, onClose func([]byte)) *ByteStore {
	store := &ByteStore{
		data:    data,
		offset:  0,
		onClose: onClose}

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

	return int64(store.offset), nil
}

// Read implements the Reader interface
func (store *ByteStore) Read(p []byte) (n int, err error) {
	size := len(p)
	n = len(store.data) - store.offset
	if n > size {
		n = size
	}
	if n < size && size > 0 {
		err = io.EOF
	}
	copy(p, store.data[store.offset:store.offset+n])
	store.offset += n

	return
}

// Write implements the Writer interface
func (store *ByteStore) Write(p []byte) (n int, err error) {
	size := len(p)
	store.ensureAvailable(size)
	copy(store.data[store.offset:], p)
	store.offset += size

	return size, nil
}

// Close implements the Closer interface.
func (store *ByteStore) Close() error {
	store.onClose(store.data)
	return nil
}

func (store *ByteStore) ensureAvailable(size int) {
	expected := store.offset + size
	available := len(store.data)

	if expected > available {
		if expected > cap(store.data) {
			old := store.data

			store.data = make([]byte, expected, expected+bufferCapacityIncrement)
			copy(store.data, old)
		} else {
			store.data = append(store.data, make([]byte, expected-available)...)
		}
	}
}
