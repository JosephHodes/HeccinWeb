import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {BlogpageComponent} from "../blogpage/blogpage.component";

const routes: Routes = [{path: ':id', component: BlogpageComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BlogRoutingModule {
}
