package tags

import (
	"fmt"
	"unicode"
)

// DeriveHashTagsFromString returns the list of `Tag` instances derived from "hashtags" (inclusive of regular tags
// and machine tags) contained in 'body'. As of this writing all matches are parsed using the `NewStringTag` method.
func DeriveHashTagsFromString(body string) ([]Tag, error) {

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
