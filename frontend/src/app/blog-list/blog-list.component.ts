import { Component, OnInit } from '@angular/core';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';

@Component({
  selector: 'app-blog-list',
  templateUrl: './blog-list.component.html',
  styleUrls: ['./blog-list.component.css']
})
export class BlogListComponent implements OnInit {

  posts: BlogPost[];

  constructor(
    private postService: PostService
  ) { }

  ngOnInit() {
    this.getPosts();
  }

  getPosts(): void {
    this.postService
        .getPosts()
        .then(response => this.posts = response);
  }

  editPost(post: BlogPost): void {
    // TODO: Implement
  }
}
