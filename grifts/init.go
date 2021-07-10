package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/patoui/buffalo_profile/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
