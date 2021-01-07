import { Component, OnInit } from "@angular/core";
import { Apollo, gql } from "apollo-angular";

@Component({
  selector: "ngx-dashboard",
  templateUrl: "./dashboard.component.html",
})
export class DashboardComponent implements OnInit {
  constructor(private apollo: Apollo) {}

  ngOnInit() {}
}
