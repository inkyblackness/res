package res

// ObjectClass is a first identification of an object
type ObjectClass byte

// ObjectSubclass is the second identification of an object
type ObjectSubclass byte

// ObjectType is the specific type of a class/subclass combination
type ObjectType byte

// ObjectID completely identifies a specific object
type ObjectID struct {
	Class    ObjectClass
	Subclass ObjectSubclass
	Type     ObjectType
}
