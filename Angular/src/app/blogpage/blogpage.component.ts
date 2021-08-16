import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";


@Component({
  templateUrl: './blogpage.component.html',
  styleUrls: ['./blogpage.component.scss']
})
export class BlogpageComponent implements OnInit {

  constructor(private ActivatedRoute: ActivatedRoute) {

  }

  ngOnInit(): void {
    console.log(this.ActivatedRoute.snapshot.params.id)
  }

}
