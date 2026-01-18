package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"

	blogrenderer "example.com/renderer"
)

func TestRenderer(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello, world",
			Body:        "This is a Post",
			Description: "This is a Description",
			Tags:        []string{"TDD", "Go"},
		}
	)

	pr, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into html", func(t *testing.T) {

		t.Run("it converts a single post into HTML", func(t *testing.T) {
			buf := bytes.Buffer{}

			if err := pr.Render(&buf, aPost); err != nil {
				t.Fatal(err)
			}

			approvals.VerifyString(t, buf.String())
		})

	})
	t.Run("it renders an index of posts", func(t *testing.T) {

		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := pr.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})
}
func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}

}
