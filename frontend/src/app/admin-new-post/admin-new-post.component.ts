import { Component, OnInit } from '@angular/core';

import { BlogPost } from '../shared/blogPost';

@Component({
  selector: 'app-admin-new-post',
  templateUrl: './admin-new-post.component.html',
  styleUrls: ['./admin-new-post.component.css']
})
export class AdminNewPostComponent implements OnInit {

  model = new BlogPost();

  constructor() { }

  ngOnInit() {
  }
}
