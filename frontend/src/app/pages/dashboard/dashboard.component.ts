import { Component, OnInit } from '@angular/core';
import {Apollo, gql} from 'apollo-angular';

@Component({
  selector: 'ngx-dashboard',
  templateUrl: './dashboard.component.html',
})
export class DashboardComponent implements OnInit {
  rates: any[];
  loading = true;
  error: any;

  constructor(private apollo: Apollo) {}

  ngOnInit() {
    this.apollo
      .watchQuery({
        query: gql`
          {
            users {
              ID
              email
            }
          }
        `,
      })
      .valueChanges.subscribe((result: any) => {
        this.rates = result?.data?.rates;
        this.loading = result.loading;
        this.error = result.error;
      });
  }
}
