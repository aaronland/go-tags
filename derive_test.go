package tags

import (
	"testing"
)

func TestDeriveTagsFromString(t *testing.T) {

	t1, err := NewStringTag("test")

	if err != nil {
		t.Fatal(err)
	}

	t2, err := NewStringTag("test2")

	if err != nil {
		t.Fatal(err)
	}

	mt1, err := NewStringTag("hello:world=test")

	if err != nil {
		t.Fatal(err)
	}

	mt2, err := NewStringTag("omg:wtf=bbq")

	if err != nil {
		t.Fatal(err)
	}

	tests := map[string][]Tag{
		"Hello world": []Tag{},
		"Text with tag #test": []Tag{
			t1,
		},
		"Text with machine tag #hello:world=test": []Tag{
			mt1,
		},
		"Test with mixed tags #test #omg:wtf=bbq #test2": []Tag{
			t1,
			mt2,
			t2,
		},
		`Multiline test

Some text
#test #test2
#omg:wtf=bbq

#hello:world=test`: []Tag{
			t1,
			t2,
			mt2,
			mt1,
		},
	}

	for str, expected_tags := range tests {

		tags, err := DeriveTagsFromString(str)

		if err != nil {
			t.Fatalf("Failed to derive tags from string '%s', %v", str, err)
		}

		if len(tags) != len(expected_tags) {
			t.Fatalf("Failed to derive expected tag count from string '%s'. Expected %d but got %d.", str, len(expected_tags), len(tags))
		}

		for idx, tags_t := range tags {

			expected_t := expected_tags[idx]

			if tags_t.Raw() != expected_t.Raw() {
				t.Fatalf("Failed to match tag derived from string '%s' at offset %d. Expected '%s' but got '%s'.", str, idx, expected_t.Raw(), tags_t.Raw())
			}

			if expected_t.IsMachineTag() && !tags_t.IsMachineTag() {
				t.Fatalf("Expected tag derived from string '%s' at offset %d to be a machine tag.", str, idx)
			}
		}
	}
}
