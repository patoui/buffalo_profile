package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/patoui/buffalo_profile/models"
)

// List gets all Posts. This function is mapped to the path
// GET /posts
func PostList(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	posts := &models.Posts{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Posts from the DB
	if err := q.Where("published_at is not null").All(posts); err != nil {
		return err
	}

	c.Set("pagination", q.Paginator)
	c.Set("posts", posts)

	return c.Render(http.StatusOK, r.HTML("/posts/index.plush.html"))
}

// Show gets the data for one Post. This function is mapped to
// the path GET /posts/{post_id}
func PostShow(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Post
	post := &models.Post{}

	// To find the Post the parameter post_slug is used.
	if err := tx.Where("slug = ?", c.Param("post_slug")).First(post); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("post", post)
	c.Set("pageTitle", post.Title)

	return c.Render(http.StatusOK, r.HTML("/posts/show.plush.html"))
}
