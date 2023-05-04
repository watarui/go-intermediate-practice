package controllers_test

import (
	"testing"

	"github.com/watarui/go-intermediate-practice/controllers"
	"github.com/watarui/go-intermediate-practice/controllers/testdata"
)

var ac *controllers.ArticleController

func TestMain(m *testing.M) {
	s := testdata.NewServiceMock()
	ac = controllers.NewArticleController(s)

	m.Run()
}
