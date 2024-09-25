package tags

import (
	"fmt"
	"unicode"
)

// DeriveTagsFromString returns the list of `Tag` instances contained in 'body'
func DeriveTagsFromString(body string) ([]Tag, error) {

	tags := make([]Tag, 0)
	str_t := ""

	capture := false

	for _, r := range body {

		if string(r) == "#" {
			capture = true
			continue
		}

		if capture && unicode.IsSpace(r) {
			capture = false

			if str_t != "" {

				t, err := NewStringTag(str_t)

				if err != nil {
					return nil, fmt.Errorf("Failed to derive new string tag from '%s', %w", str_t, err)
				}

				tags = append(tags, t)
				str_t = ""
			}
		}

		if capture {
			str_t = str_t + string(r)
		}
	}

	if str_t != "" {

		t, err := NewStringTag(str_t)

		if err != nil {
			return nil, fmt.Errorf("Failed to derive new string tag from '%s', %w", str_t, err)
		}

		tags = append(tags, t)
	}

	return tags, nil
}
