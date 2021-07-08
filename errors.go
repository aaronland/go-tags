package tags

// type NotMachineTagError defines an implementation of the error interface signaling
// that a tag is not a valid machine tag.
type NotMachineTagError struct{}

// The string representation of a NotMachineTagError error.
func (e NotMachineTagError) Error() string {
	return "Tag is not a machine tag"
}

// Return a boolean value indicating whether or not err is of type NotMachineTagError
func IsNotMachineTagError(err error) bool {

	_, ok := err.(NotMachineTagError)

	if ok {
		return true
	}

	return false
}
