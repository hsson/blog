import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';

@Component({
  selector: 'app-admin-new-post',
  templateUrl: './admin-new-post.component.html',
  styleUrls: ['./admin-new-post.component.css']
})
export class AdminNewPostComponent implements OnInit {

  model = new BlogPost();

  constructor(
    private postService: PostService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  createPost(): void {
    if (!this.model.title || !this.model.body) {
      return;
    }

    this.postService
      .newPost(this.model)
      .then(newPost => this.processPost(newPost));
  }

  processPost(post: BlogPost): void {
    console.log(post);
    let link = ["/post", post.slug];
    this.router.navigate(link);
  }
}
