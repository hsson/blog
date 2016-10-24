import { Component, OnInit } from '@angular/core';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';

@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.css']
})
export class BlogComponent implements OnInit {

  title = "Blog posts";
  posts: BlogPost[];

  constructor(
    private postService: PostService
  ) {  }

  debugPosts(): void {
    console.log(this.posts);
  }

  getPosts(): void {
    this.postService
        .getPosts()
        .then(response => this.posts = response);
  }

  ngOnInit() {
    this.getPosts();
  }

}
