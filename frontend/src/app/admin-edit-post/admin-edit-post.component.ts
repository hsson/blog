import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';

import { PostService } from '../post.service';
import { BlogPost } from '../shared/blogPost';
import { EditType } from '../admin-new-post/admin-new-post.component';

@Component({
  selector: 'app-admin-edit-post',
  templateUrl: './admin-edit-post.component.html',
  styleUrls: ['./admin-edit-post.component.css']
})
export class AdminEditPostComponent implements OnInit {

  post: BlogPost;

  editType = EditType.EditPost;

  preview = false;

  constructor(
    private route: ActivatedRoute,
    private postService: PostService
    ) { }

  ngOnInit() {
    this.route.params.forEach((params: Params) => {
      let slug = params['slug'];
      this.postService.getPost(slug)
        .then(post => this.post = post);
    });
  }

}
