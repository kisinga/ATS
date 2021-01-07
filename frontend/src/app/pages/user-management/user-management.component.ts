import { Component, OnInit } from '@angular/core';
import {Apollo, gql} from 'apollo-angular';

@Component({
  templateUrl: './user-management.component.html',
  styleUrls: ['./user-management.component.scss']
})
export class UserManagementComponent implements OnInit {

  constructor(private apollo: Apollo) { }

  ngOnInit(): void {
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
   
    });
  }

}
