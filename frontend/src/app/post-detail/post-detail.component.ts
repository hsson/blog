import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { Title } from '@angular/platform-browser';

import { PostService } from '../post.service';
import { BlogPost } from '../shared/blogPost';

@Component({
  selector: 'app-post-detail',
  templateUrl: './post-detail.component.html',
  styleUrls: ['./post-detail.component.css']
})
export class PostDetailComponent implements OnInit {

  post: BlogPost;

  constructor(
    private postService: PostService,
    private route: ActivatedRoute,
    private router: Router,
    private titleService: Title
  ) { }

  ngOnInit() {
    this.route.params.forEach((params: Params) => {
      let slug = params['slug'];
      this.postService.getPost(slug)
        .then(post => this.setPost(post));
    });
  }

  setPost(newPost: BlogPost): void {
    this.post = newPost;
    this.titleService.setTitle(`${newPost.title} | blog.hakansson.xyz`);
  }

  goToHome(): void {
    let link = [''];
    this.router.navigate(link);
  }

}
