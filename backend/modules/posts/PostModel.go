package posts

import (
	"time"

	"github.com/Machiel/slugify"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Post struct {
	Slug      string    `json:"slug" datastore:"-"`
	PostDate  time.Time `json:"postDate" datastore:"post_date"`
	Title     string    `json:"title" datastore:"post_title"`
	Body      string    `json:"body" datastore:"post_body"`
	Published bool      `json:"published" datastore:"post_published"`
}

func (p *Post) key(c context.Context) *datastore.Key {
	if p.Slug == "" {
		p.Slug = slugify.Slugify(p.Title)
		return datastore.NewKey(c, "blog_post", p.Slug, 0, nil)
	}

	return datastore.NewKey(c, "blog_post", p.Slug, 0, nil)
}

func keyToString(c context.Context, s string) *datastore.Key {
	return datastore.NewKey(c, "blog_post", s, 0, nil)
}

func (p *Post) save(c context.Context) error {
	key, err := datastore.Put(c, p.key(c), p)
	if err != nil {
		return err
	}
	p.Slug = key.StringID()
	return nil
}

func getPost(c context.Context, slug string) (*Post, error) {
	var post Post
	post.Slug = slug

	key := post.key(c)
	err := datastore.Get(c, key, &post)
	if err != nil {
		return nil, err
	}

	post.Slug = key.StringID()
	return &post, nil
}

func getPosts(c context.Context) ([]Post, error) {
	q := datastore.NewQuery("blog_post").Order("-post_date")

	var posts []Post
	keys, err := q.GetAll(c, &posts)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(posts); i++ {
		posts[i].Slug = keys[i].StringID()
	}
	return posts, nil
}
