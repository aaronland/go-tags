package tags

import (
	"testing"
)

func TestStringTag(tt *testing.T) {

	clean_tests := map[string]string{
		"Hello, World":    "helloworld",
		"foo:bar=baz":     "foo:bar=baz",
		"foo:bar=Mr. Baz": "foo:bar=mrbaz",
	}

	for raw, clean := range clean_tests {

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

	mt_tests := map[string][]string{
		"foo:bar=Mr. Baz": []string{"foo", "bar", "Mr. Baz"},
	}

	for raw, triple := range mt_tests {

		t, err := NewStringTag(raw)

		if err != nil {
			tt.Fatal(err)
		}

		if !t.IsMachineTag() {
			tt.Fatalf("Not a machine tag '%s'", raw)
		}

		expected_ns := triple[0]
		expected_pred := triple[1]
		expected_value := triple[2]

		ns, _ := t.Namespace()
		pred, _ := t.Predicate()
		value, _ := t.Value()

		if ns != expected_ns {
			tt.Fatalf("Invalid machine tag namespace, '%s'. Expected '%s'.", ns, expected_ns)
		}

		if pred != expected_pred {
			tt.Fatalf("Invalid machine tag predicate, '%s'. Expected '%s'.", pred, expected_pred)
		}

		if value != expected_value {
			tt.Fatalf("Invalid machine tag value, '%s'. Expected '%s'.", value, expected_value)
		}

	}
}
