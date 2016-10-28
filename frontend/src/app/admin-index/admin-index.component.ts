import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';

const COMPONENT_TITLE = "Admin | blog.hakansson.xyz";

@Component({
  selector: 'app-admin-index',
  templateUrl: './admin-index.component.html',
  styleUrls: ['./admin-index.component.css']
})
export class AdminIndexComponent implements OnInit {

  preview = false;

  constructor(
    private titleService: Title
  ) { }

  ngOnInit() {
    this.titleService.setTitle(COMPONENT_TITLE);
  }

}
