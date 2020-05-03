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

	valid := [][]int{
		[]int{48, 57},  // (0-9)
		[]int{65, 90},  // (A-Z)
		[]int{97, 122}, // (a-z)
		[]int{45},      // -
		[]int{58},      // :
		[]int{61},      // =
		[]int{95},      // _
	}

	for _, r := range raw {

		is_valid := false

		for _, bookends := range valid {

			r_int := int(r)

			switch len(bookends) {
			case 1:

				if r_int == bookends[0] {
					is_valid = true
				}
			default:
				if r_int >= bookends[0] && r_int <= bookends[1] {
					is_valid = true
				}
			}

			if is_valid {
				break
			}
		}

		if !is_valid {
			continue
		}

		c := fmt.Sprintf("%c", r)
		clean = append(clean, c)
	}

	str_clean := strings.Join(clean, "")
	str_clean = strings.ToLower(str_clean)

	return str_clean, nil
}
