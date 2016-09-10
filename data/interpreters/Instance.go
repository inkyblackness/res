package interpreters

// Instance is one instantiated interpreter for a data block,
// based on a predetermined description.
type Instance struct {
	desc *Description
	data []byte
}

// Raw returns the raw array.
func (inst *Instance) Raw() []byte {
	return inst.data
}

// Keys returns an array of all registered keys, sorted by start index
func (inst *Instance) Keys() []string {
	return sortKeys(inst.desc.fields)
}

// ActiveRefinements returns an array of all keys where the corresponding refinement
// is active (The corresponding predicate returns true).
func (inst *Instance) ActiveRefinements() (keys []string) {
	entries := make(map[string]*entry)
	for key, r := range inst.desc.refinements {
		entries[key] = &r.entry
	}
	sortedKeys := sortKeys(entries)
	for _, key := range sortedKeys {
		if inst.desc.refinements[key].predicate(inst) {
			keys = append(keys, key)
		}
	}

	return
}

// Get returns the value associated with the given key. Should there be no
// value for the requested key, the function returns 0.
func (inst *Instance) Get(key string) uint32 {
	e := inst.desc.fields[key]
	value := uint32(0)

	if e != nil && inst.isValidRange(e) {
		for i := 0; i < e.count; i++ {
			value = (value << 8) | uint32(inst.data[e.start+e.count-1-i])
		}
	}

	return value
}

// Set stores the provided value with the given key. Should there be no
// registration for the key, the function does nothing.
func (inst *Instance) Set(key string, value uint32) {
	e := inst.desc.fields[key]

	if e != nil && inst.isValidRange(e) {
		for i := 0; i < e.count; i++ {
			inst.data[e.start+i] = byte(value >> uint32(i*8))
		}
	}
}

// Refined returns an instance that was nested according to the description.
// This method returns an instance for registered refinements even if their predicate
// would not specify them being active.
// Should the refinement not exist, an instance without any fields will be returned.
func (inst *Instance) Refined(key string) (refined *Instance) {
	r := inst.desc.refinements[key]

	if r != nil && inst.isValidRange(&r.entry) {
		refined = &Instance{
			desc: r.desc,
			data: inst.data[r.start : r.start+r.count]}
	} else {
		refined = New().For(nil)
	}

	return
}

func (inst *Instance) isValidRange(e *entry) bool {
	available := len(inst.data)

	return (e.start + e.count) <= available
}
