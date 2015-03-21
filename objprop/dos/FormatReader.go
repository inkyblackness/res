package dos

import (
	"fmt"
	"io"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/objprop"
	"github.com/inkyblackness/res/serial"
)

type typeEntry struct {
	genericOffset  uint32
	genericLength  uint32
	specificOffset uint32
	specificLength uint32
	commonOffset   uint32
}

type formatReader struct {
	coder   serial.PositioningCoder
	entries map[res.ObjectID]*typeEntry
}

var errFormatMismatch = fmt.Errorf("Format mismatch")

// NewProvider wraps the provided ReadSeeker in a provider for object properties.
func NewProvider(source io.ReadSeeker, descriptors []objprop.ClassDescriptor) (provider objprop.Provider, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	if source == nil {
		panic(fmt.Errorf("source is nil"))
	}
	verifySourceLength(source, descriptors)

	provider = &formatReader{
		coder:   serial.NewPositioningDecoder(source),
		entries: calculateEntryValues(descriptors)}

	return
}

func (provider *formatReader) Provide(id res.ObjectID) objprop.ObjectData {
	entry := provider.entries[id]
	result := objprop.ObjectData{
		Generic:  provider.readData(entry.genericOffset, entry.genericLength),
		Specific: provider.readData(entry.specificOffset, entry.specificLength),
		Common:   provider.readData(entry.commonOffset, objprop.CommonPropertiesLength)}

	return result
}

func (provider *formatReader) readData(offset uint32, length uint32) []byte {
	data := make([]byte, length)

	provider.coder.SetCurPos(offset)
	provider.coder.CodeBytes(data)

	return data
}

func verifySourceLength(source io.Seeker, descriptors []objprop.ClassDescriptor) {
	sourceLength := getSeekerSize(source)

	expectedLength := uint32(0)
	expectedLength += uint32(4)
	for _, classDesc := range descriptors {
		expectedLength += classDesc.TotalDataLength()
	}
	if expectedLength != sourceLength {
		panic(errFormatMismatch)
	}
}

func getSeekerSize(seeker io.Seeker) uint32 {
	length, err := seeker.Seek(0, 2)

	if err != nil {
		panic(err)
	}

	return uint32(length)
}

func calculateEntryValues(descriptors []objprop.ClassDescriptor) map[res.ObjectID]*typeEntry {
	startOffset := uint32(4)
	entries := make(map[res.ObjectID]*typeEntry)
	var entryList []*typeEntry

	for classIndex, classDesc := range descriptors {
		genericOffset := startOffset
		specificOffset := startOffset + classDesc.GenericDataLength*classDesc.TotalTypeCount()

		for subclassIndex, subclassDesc := range classDesc.Subclasses {
			for typeIndex := uint32(0); typeIndex < subclassDesc.TypeCount; typeIndex++ {
				entry := &typeEntry{
					genericOffset:  genericOffset,
					genericLength:  classDesc.GenericDataLength,
					specificOffset: specificOffset,
					specificLength: subclassDesc.SpecificDataLength}
				entryKey := res.MakeObjectID(res.ObjectClass(classIndex), res.ObjectSubclass(subclassIndex), res.ObjectType(typeIndex))

				entries[entryKey] = entry
				entryList = append(entryList, entry)
				specificOffset += subclassDesc.SpecificDataLength
				genericOffset += classDesc.GenericDataLength
			}
			startOffset = specificOffset
		}
	}
	for _, entry := range entryList {
		entry.commonOffset = startOffset
		startOffset += objprop.CommonPropertiesLength
	}

	return entries
}
