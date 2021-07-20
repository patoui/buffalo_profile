package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/patoui/buffalo_profile/models"
	"github.com/pkg/errors"
)

// AdminPostsNew renders the post form
func AdminPostsNew(c buffalo.Context) error {
	p := models.Post{}
	c.Set("post", p)
	c.Set("pageTitle", "Create")
	return c.Render(http.StatusOK, r.HTML("admin/posts/new.plush.html"))
}

// AdminPostsCreate registers a new user with the application.
func AdminPostsCreate(c buffalo.Context) error {
	p := &models.Post{}
	if err := c.Bind(p); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := tx.ValidateAndCreate(p)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("post", p)
		c.Set("errors", verrs)
		return c.Render(200, r.HTML("admin/posts/new.plush.html"))
	}

	return c.Redirect(302, "/admin/dashboard")
}
