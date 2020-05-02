package tags

import ()

type Tag interface {
	Raw() string
	Clean() string
}
