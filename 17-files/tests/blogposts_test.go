package tests

import (
	_ "embed"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/evgenymarkov/learn-golang/17-files/pkg/blogposts"
)

var (
	//go:embed data/post-1.txt
	post1 string
	//go:embed data/post-2.txt
	post2 string
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func (s StubFailingFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(post1)},
		"hello-world2.md": {Data: []byte(post2)},
	}

	posts, postsErr := blogposts.NewPostsFromFS(fs)

	wantPosts := []blogposts.Post{
		{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body:        "Hello\nWorld",
		},
		{
			Title:       "Post 2",
			Description: "Description 2",
			Tags:        []string{"js"},
			Body:        "Nothing special",
		},
	}

	if postsErr != nil {
		t.Fatal(postsErr)
	}

	assertPostsCount(t, len(posts), len(wantPosts))

	for i, post := range posts {
		assertPostEqual(t, post, wantPosts[i])
	}
}

func TestNewBlogPostsFail(t *testing.T) {
	fs := StubFailingFS{}

	posts, postsErr := blogposts.NewPostsFromFS(fs)

	if postsErr == nil {
		t.Fatal("expected an error but didn't get one")
	}

	assertPostsCount(t, len(posts), 0)
}

func assertPostsCount(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Fatalf("got %d posts, but expected %d", got, want)
	}
}

func assertPostEqual(t testing.TB, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v post, want %+v", got, want)
	}
}
