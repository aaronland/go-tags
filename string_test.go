package tags

import (
	"testing"
)

func TestStringTag(tt *testing.T) {

	raw := "Hello, World"
	clean := "helloworld"

	t, err := NewStringTag(raw)

	if err != nil {
		tt.Fatal(err)
	}

	if t.Raw() != raw {
		tt.Fatal("Invalid raw tag")
	}

	if t.Clean() != clean {
		tt.Fatal("Invalid clean tag")
	}

}
