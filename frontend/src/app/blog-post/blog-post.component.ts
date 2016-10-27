import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';

import { BlogPost } from '../shared/blogPost';

@Component({
  selector: 'app-blog-post',
  templateUrl: './blog-post.component.html',
  styleUrls: ['./blog-post.component.css']
})
export class BlogPostComponent implements OnInit {

  @Input()
  post: BlogPost;

  constructor(
    private router: Router
  ) { }

  ngOnInit() {
  }

  goToDetail(post: BlogPost): void {
    let link = ['/post', post.slug];
    this.router.navigate(link);
  }

}
