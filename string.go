package tags

import (
	"fmt"
	"github.com/whosonfirst/go-sanitize"
	"strings"
)

type StringTag struct {
	raw   string
	clean string
}

func NewStringTag(raw string) (Tag, error) {

	opts := sanitize.DefaultOptions()
	sanitized, err := sanitize.SanitizeString(raw, opts)

	if err != nil {
		return nil, err
	}

	clean, err := CleanStringTag(sanitized)

	if err != nil {
		return nil, err
	}

	t := &StringTag{
		raw:   raw,
		clean: clean,
	}

	return t, nil
}

func (t *StringTag) Raw() string {
	return t.raw
}

func (t *StringTag) Clean() string {
	return t.clean
}

func CleanStringTag(raw string) (string, error) {

	clean := make([]string, 0)

	alpha_numeric := [][]int{
		[]int{48, 57},  // (0-9)
		[]int{65, 90},  // (A-Z)
		[]int{97, 122}, // (a-z)
	}

	for _, r := range raw {

		is_alpha_numeric := false

		for _, bookends := range alpha_numeric {

			r_int := int(r)

			if r_int >= bookends[0] && r_int <= bookends[1] {
				is_alpha_numeric = true
				break
			}
		}

		if !is_alpha_numeric {
			continue
		}

		c := fmt.Sprintf("%c", r)
		clean = append(clean, c)
	}

	str_clean := strings.Join(clean, "")
	str_clean = strings.ToLower(str_clean)

	return str_clean, nil
}
