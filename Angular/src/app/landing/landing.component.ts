import {Component, HostListener, OnInit} from '@angular/core';


@Component({
  selector: 'app-landing',
  templateUrl: './landing.component.html',
  styleUrls: ['./landing.component.scss']
})
export class LandingComponent implements OnInit {


  constructor() {
  }

  @HostListener('window:scroll')
  checkScroll() {

    // Both window.pageYOffset and document.documentElement.scrollTop returns the same result in all the cases. window.pageYOffset is not supported below IE 9.

    const scrollPosition = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop || 0;

    console.log('[scroll]', scrollPosition);
  }

  scrollToTop = () => {
    window.scroll({
      top: -1000,
      left: 0,
      behavior: 'smooth'
    });
  }

  ngOnInit(): void {
  }

}
