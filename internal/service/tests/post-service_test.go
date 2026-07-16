package tests

import (
	"blog/internal/models"
	"blog/internal/service"
	"net/http"
	"testing"
)

func TestArticleValidation(t *testing.T) {
	tests := []struct {
		test_name string
		article models.Article
		want_code int
	}{{
		test_name: "valid article",
		article: models.Article{Title: "Matadora linux", Content: "Matadora linux is the best!!!!!!!"},
		want_code: http.StatusOK,
	},
	{
		test_name: "short title",
		article: models.Article{Title: "Ma", Content: "Matadora linux!!!"},
		want_code: http.StatusBadRequest,
	},
	{
		test_name: "short content",
		article: models.Article{Title: "Simbios", Content: "11"},
		want_code: http.StatusUnprocessableEntity,
	},
	}

	ps := &service.PostService{}

	for _, tt := range tests{
		t.Run(tt.test_name, func(t *testing.T) {
			got := ps.ValidateArticle(tt.article)
			if got != tt.want_code{
				t.Errorf("ValidateUserData() = %d; want = %d", got, tt.want_code)
			}
		})
	}
}
