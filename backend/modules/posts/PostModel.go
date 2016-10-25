package posts

type Post struct {
  Id        int     `json:"id"`
  PostDate  int     `json:"postDate"`
  Title     string  `json:"title"`
  Body      string  `json:"body"`
}
