import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';
import { Router } from '@angular/router';

const COMPONENT_TITLE = "Admin | blog.hakansson.xyz";

@Component({
  selector: 'app-admin-index',
  templateUrl: './admin-index.component.html',
  styleUrls: ['./admin-index.component.css']
})
export class AdminIndexComponent implements OnInit {

  post = new BlogPost();

  preview = false;

  constructor(
    private titleService: Title,
    private postService: PostService,
    private router: Router
  ) { }

  ngOnInit() {
    this.titleService.setTitle(COMPONENT_TITLE);
  }

  createNewPost(): void {
    if (!this.post.title || !this.post.body) {
      return;
    }

    this.postService
      .newPost(this.post)
      .then(newPost => this.processPost(newPost));
  }

   processPost(post: BlogPost): void {
    console.log(post);
    let link = ["/post", post.slug];
    this.router.navigate(link);
  }

}
