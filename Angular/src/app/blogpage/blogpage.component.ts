import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";


@Component({
  templateUrl: './blogpage.component.html',
  styleUrls: ['./blogpage.component.scss']
})
export class BlogpageComponent implements OnInit {
  public articleName: string | null = null;

  constructor(private ActivatedRoute: ActivatedRoute, private routes: Router) {

  }

  ngOnInit(): void {
    this.articleName = this.ActivatedRoute.snapshot.params.id
    if (this.articleName == null) {
      this.routes.navigate(['blog']).catch(err => console.log(err))
    }
  }

}
