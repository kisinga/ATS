import { Component, OnInit } from "@angular/core";
import { FormGroup, FormControl, Validators } from "@angular/forms";
import { NbDialogRef } from "@nebular/theme";
import { UserService } from "app/pages/shared/services/user.service";

@Component({
  templateUrl: "./new-user.component.html",
  styleUrls: ["./new-user.component.scss"],
})
export class NewUserComponent implements OnInit {
  constructor(
    protected ref: NbDialogRef<NewUserComponent>,
    private userService: UserService
  ) {}
  newUserForm = new FormGroup({
    name: new FormControl(
      "",
      Validators.compose([Validators.required, Validators.min(5)])
    ),
    email: new FormControl(
      "",
      Validators.compose([Validators.required, Validators.email])
    ),
  });

  ngOnInit(): void {}

  cancel() {
    this.ref.close();
  }
  submit() {
    if (this.newUserForm.valid) {
      this.userService.createUser(this.newUserForm.value).then((r) => {
        console.log(r);
        this.ref.close();
      });
    } else {
    }
  }
}
