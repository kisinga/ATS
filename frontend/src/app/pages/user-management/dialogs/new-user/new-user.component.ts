import { Component, OnInit } from "@angular/core";
import { NbDialogRef } from "@nebular/theme";
import { NewUserInput } from "app/models/gql/user.query";

@Component({
  templateUrl: "./new-user.component.html",
  styleUrls: ["./new-user.component.scss"],
})
export class NewUserComponent implements OnInit {
  constructor(protected ref: NbDialogRef<NewUserComponent>) {}
  newUSer: NewUserInput;
  ngOnInit(): void {}

  cancel() {
    this.ref.close();
  }
  submit() {
    this.ref.close(this.newUSer);
  }
}
