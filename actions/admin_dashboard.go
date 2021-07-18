package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/patoui/buffalo_profile/models"
)

// AdminDashboard display the admin dashboard
func AdminDashboard(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	posts := []models.Post{}

	if err := tx.Order("created_at desc").All(&posts); err != nil {
		return err
	}

	postPublishedCount, _ := tx.Where("published_at is not null").Count(models.Post{})

	c.Set("posts", posts)
	c.Set("postCount", len(posts))
	c.Set("postPublishedCount", postPublishedCount)

	c.Set("pageTitle", "Admin Dashboard")
	return c.Render(http.StatusOK, r.HTML("admin/dashboard.plush.html"))
}
