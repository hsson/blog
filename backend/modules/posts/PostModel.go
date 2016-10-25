package posts

import (
  "golang.org/x/net/context"
  "google.golang.org/appengine/datastore"
  "time"
)

type Post struct {
  Id        int64         `json:"id" datastore:"-"`
  PostDate  time.Time     `json:"postDate" datastore:"post_date"`
  Title     string        `json:"title" datastore:"post_title"`
  Body      string        `json:"body" datastore:"post_body"`
}

func (p *Post) key(c context.Context) *datastore.Key {
  if p.Id == 0 {
    return datastore.NewIncompleteKey(c, "blog_post", nil)
  }

  return datastore.NewKey(c, "blog_post", "", p.Id, nil)
}

func (p *Post) save(c context.Context) error {
  key, err := datastore.Put(c, p.key(c), p)
  if err != nil {
    return err
  }

  p.Id = key.IntID()
  return nil
}

func getPosts(c context.Context) ([]Post, error) {
  q := datastore.NewQuery("blog_post").Order("-post_date")

  var posts []Post
  keys, err := q.GetAll(c, &posts)
  if err != nil {
    return nil, err
  }

  for i := 0; i < len(posts); i++ {
    posts[i].Id = keys[i].IntID()
  }
  return posts, nil
}
