import { Component, OnInit, Input, ViewChild } from '@angular/core';
import { Router } from '@angular/router';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';

export enum EditType {NewPost, EditPost}

@Component({
  selector: 'app-admin-new-post',
  templateUrl: './admin-new-post.component.html',
  styleUrls: ['./admin-new-post.component.css']
})
export class AdminNewPostComponent implements OnInit {

  @Input()
  preview: boolean;

  @Input()
  post: BlogPost;

  @Input()
  editType: EditType;

  constructor(
    private postService: PostService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  submitPost(): void {
    if (this.editType == EditType.NewPost) {
      this.createNewPost();
    }
  }

  createNewPost(): void {
    if (!this.post.title || !this.post.body) {
      return;
    }

    this.postService
      .newPost(this.post)
      .then(newPost => this.processNewPost(newPost));
  }

   processNewPost(post: BlogPost): void {
    console.log(post);
    let link = ["/post", post.slug];
    this.router.navigate(link);
  }
}
