package tags

import ()

type NotMachineTagError struct{}

func (e NotMachineTagError) Error() string {
	return "Tag is not a machine tag"
}

func IsNotMachineTagError(err error) bool {
	_, ok := err.(NotMachineTagError)

	if ok {
		return true
	}

	return false
}

type Tag interface {
	Raw() string
	Clean() string
	IsMachineTag() bool
	Namespace() (string, error)
	Value() (string, error)
	Predicate() (string, error)
}
