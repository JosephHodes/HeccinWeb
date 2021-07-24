import {Component, HostListener, OnInit} from '@angular/core';

@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.css']
})
export class LandingComponent implements OnInit {
  public screenWidth: number = 0;

  constructor() {

  }

  @HostListener('window:resize', ['$event'])
  onResize(event: any) {
    this.screenWidth = event.target.innerWidth;

  }

  ngOnInit(): void {
    this.screenWidth = screen.availWidth;
  }

}
