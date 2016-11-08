import { Component, OnInit, Input, ViewChild } from '@angular/core';
import { Router } from '@angular/router';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';

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
  postFunc: Function;

  constructor(
    private postService: PostService,
    private router: Router
  ) { }

  ngOnInit() {
  }

  getPost(): BlogPost {
    return this.post;
  }
}
