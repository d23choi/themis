package local

import (
	"github.com/infobloxopen/themis/pdp"
)

func init() {
	pdp.RegisterSelector("local", &Selector{})
}
