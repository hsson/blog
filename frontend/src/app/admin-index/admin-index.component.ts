import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';

import { BlogPost } from '../shared/blogPost';
import { PostService } from '../post.service';
import { EditType } from '../admin-new-post/admin-new-post.component';

const COMPONENT_TITLE = "Admin | blog.hakansson.xyz";

@Component({
  selector: 'app-admin-index',
  templateUrl: './admin-index.component.html',
  styleUrls: ['./admin-index.component.css']
})
export class AdminIndexComponent implements OnInit {

  post = new BlogPost();

  editType = EditType.NewPost;

  preview = false;

  constructor(
    private titleService: Title,
    private postService: PostService,
    private router: Router
  ) { }

  ngOnInit() {
    this.titleService.setTitle(COMPONENT_TITLE);
  }

}
