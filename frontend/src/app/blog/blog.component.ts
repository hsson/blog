import { Component, OnInit } from '@angular/core';

import { BlogPost } from '../shared/blogPost';

const POSTS: BlogPost[] = [
  { id: 1, postDate: 1477336203000, title: "Hello world!", body: "I am the *body* of this post!"},
  { id: 2, postDate: 1477336203000, title: "Another post", body: "I am yet another post body!"}
]

@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.css']
})
export class BlogComponent implements OnInit {

  title = "Blog posts";
  posts = POSTS;

  constructor() { }

  ngOnInit() {
  }

}
