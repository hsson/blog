import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';

const COMPONENT_TITLE = "Blog | blog.hakansson.xyz";

@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.css']
})
export class BlogComponent implements OnInit {

  title = "Blog posts";
  posts: BlogPost[];

  constructor(
    private postService: PostService,
    private titleService: Title
  ) {  }

  getPosts(): void {
    this.postService
        .getPosts()
        .then(response => this.posts = response);
  }

  ngOnInit() {
    this.titleService.setTitle(COMPONENT_TITLE);
    this.getPosts();
  }

}
