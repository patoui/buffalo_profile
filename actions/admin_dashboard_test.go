package actions

import (
	"github.com/patoui/buffalo_profile/models"
)

func (as *ActionSuite) Test_AdminDashboard_List() {
	as.LoadFixture("one published post, one not published post")
	as.LoadFixture("admin user")

	u := &models.User{}
	as.DB.First(u)

	as.Session.Set("current_user_id", u.ID)

	res := as.HTML("/admin/dashboard").Get()

	as.Equal(200, res.Code)

	rb := res.Body.String()
	as.Contains(rb, "This is published post title #1")
	as.NotContains(rb, "This is published post title #2")

	posts := []models.Post{}
	as.DB.Order("created_at desc").All(&posts)

	as.Contains(rb, "Published on "+posts[0].ShortPublishedAt())
	as.Contains(rb, "Not published")

	// Update user email to non-admin
	u.Email = "johndoe@gmail.com"
	as.DB.Update(u)

	res = as.HTML("/admin/dashboard").Get()

	// ensure non-admin users are redirected
	as.Equal(302, res.Code)
}
