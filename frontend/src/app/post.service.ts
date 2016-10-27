import { Injectable } from '@angular/core';
import { Headers, Http, RequestOptions, Response } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { BlogPost } from './shared/blogPost';

@Injectable()
export class PostService {

  private postsUrl = 'api/posts';  // URL to web api
  private newPostUrl = 'api/admin/posts'; // URL to create new posts

  constructor(private http: Http) { }

  getPosts(): Promise<BlogPost[]> {
    return this.http.get(this.postsUrl)
      .toPromise()
      .then(response => response.json() as BlogPost[])
      .catch(this.handleError);
  }

  getPost(slug: string): Promise<BlogPost> {
    return this.http.get(`${this.postsUrl}/${slug}`)
      .toPromise()
      .then(response => response.json() as BlogPost)
      .catch(this.handleError)
  }

  newPost(post: BlogPost): Promise<BlogPost> {
    let headers = new Headers({ 'Content-Type': 'application/json' });
    let options = new RequestOptions({ headers: headers });

    return this.http.post(this.newPostUrl, { "title":post.title, "body":post.body }, options)
      .toPromise()
      .then(response => response.json() as BlogPost)
      .catch(this.handleError);
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error);
    return Promise.reject(error.message || error);
  }
}
