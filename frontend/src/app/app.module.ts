import { BrowserModule, Title } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MaterializeDirective } from 'angular2-materialize';
import { RouterModule } from '@angular/router';

import { AppComponent } from './app.component';
import { BlogPostComponent } from './blog-post/blog-post.component';
import { BlogComponent } from './blog/blog.component';
import { PostService } from './post.service';
import { PostDetailComponent } from './post-detail/post-detail.component';
import { AdminIndexComponent } from './admin-index/admin-index.component';
import { AdminNewPostComponent } from './admin-new-post/admin-new-post.component';
import { BlogListComponent } from './blog-list/blog-list.component';
import { AdminEditPostComponent } from './admin-edit-post/admin-edit-post.component';

@NgModule({
  declarations: [
    AppComponent,
    MaterializeDirective,
    BlogPostComponent,
    BlogComponent,
    PostDetailComponent,
    AdminIndexComponent,
    AdminNewPostComponent,
    BlogListComponent,
    AdminEditPostComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    RouterModule.forRoot([
      {
        path: '',
        component: BlogComponent
      },
      {
        path: 'post/:slug',
        component: PostDetailComponent
      },
      {
        path: 'admin',
        component: AdminIndexComponent
      },
      {
        path: 'admin/edit/:slug',
        component: AdminEditPostComponent
      }
    ])
  ],
  providers: [PostService, Title],
  bootstrap: [AppComponent]
})
export class AppModule { }
