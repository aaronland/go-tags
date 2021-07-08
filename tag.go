package tags

import ()

// type Tag defines a common interface for working with tag values.
type Tag interface {
	// The raw user-defined string tag value.
	Raw() string
	// The URI-safe value of the raw tag value.
	Clean() string
	// A boolean flag indicating whether or not the tag can be parsed as a machine tag.
	IsMachineTag() bool
	// The namespace value of a valid machine tag triple.
	Namespace() (string, error)
	// The predicate value of a valid machine tag triple.
	Predicate() (string, error)
	// The value of a valid machine tag triple.
	Value() (string, error)
}
