package tags

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/whosonfirst/go-sanitize"
)

// type StringTag implements the Tag interface for string-based tags.
type StringTag struct {
	raw           string
	clean         string
	is_machinetag bool
	triple        [3]string
}

var re_machinetag *regexp.Regexp
var valid_clean [][]int

func init() {

	valid_clean = [][]int{
		[]int{48, 57},  // (0-9)
		[]int{65, 90},  // (A-Z)
		[]int{97, 122}, // (a-z)
		[]int{45},      // -
		[]int{58},      // :
		[]int{61},      // =
		[]int{95},      // _
	}

	re_machinetag = regexp.MustCompile(`^([a-zA-Z][a-zA-Z0-9\-\_]+):([a-zA-Z][a-zA-Z0-9\-\_]+)=(.*)$`)
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

	if len(clean) < 3 {
		return nil, errors.New("Invalid tag")
	}

	t := &StringTag{
		raw:   raw,
		clean: clean,
	}

	m := re_machinetag.FindStringSubmatch(sanitized)

	if len(m) == 4 {

		t.is_machinetag = true
		t.triple = [3]string{m[1], m[2], m[3]}
	}

	return t, nil
}

func (t *StringTag) Raw() string {
	return t.raw
}

func (t *StringTag) Clean() string {
	return t.clean
}

func (t *StringTag) IsMachineTag() bool {
	return t.is_machinetag
}

func (t *StringTag) Namespace() (string, error) {

	if !t.IsMachineTag() {
		return "", NotMachineTagError{}
	}

	return t.triple[0], nil
}

func (t *StringTag) Predicate() (string, error) {

	if !t.IsMachineTag() {
		return "", NotMachineTagError{}
	}

	return t.triple[1], nil
}

func (t *StringTag) Value() (string, error) {

	if !t.IsMachineTag() {
		return "", NotMachineTagError{}
	}

	return t.triple[2], nil
}

func CleanStringTag(raw string) (string, error) {

	clean := make([]string, 0)

	for _, r := range raw {

		is_valid := false

		for _, bookends := range valid_clean {

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
