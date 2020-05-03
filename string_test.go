package tags

import (
	"testing"
)

func TestStringTag(tt *testing.T) {

	tests := map[string]string{
		"Hello, World": "helloworld",
		"foo:bar=baz":  "foo:bar=baz",
	}

	for raw, clean := range tests {

		t, err := NewStringTag(raw)

		if err != nil {
			tt.Fatal(err)
		}

		if t.Raw() != raw {
			tt.Fatalf("Invalid raw tag, '%s'. Expected '%s'.", t.Raw(), raw)
		}

		if t.Clean() != clean {
			tt.Fatalf("Invalid clean tag, '%s'. Expected '%s'.", t.Clean(), clean)
		}
	}
}
