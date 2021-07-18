package models_test

import (
	"github.com/gobuffalo/suite"
	"github.com/patoui/buffalo_profile/models"
)

type ModelSuite struct {
	*suite.Model
}

func (ms *ModelSuite) Test_Post_ShortBody() {
	p := &models.Post{Body: `### Intro
	Test body !@#$%^&*() it should remove special characters

	AND new line characters!
	`}

	ms.EqualValues("Intro Test body it should remove special characters AND new line characters", p.ShortBody())
}
