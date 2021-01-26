import {Component, OnInit} from "@angular/core";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {NbDialogRef} from "@nebular/theme";
import {UserService} from "app/pages/shared/services/user.service";

@Component({
  templateUrl: "./new-user.component.html",
  styleUrls: ["./new-user.component.scss"],
})
export class NewUserComponent implements OnInit {
  creationError = null;
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

  constructor(
    protected ref: NbDialogRef<NewUserComponent>,
    private userService: UserService
  ) {
  }

  ngOnInit(): void {
  }

  cancel() {
    this.ref.close();
  }

  submit() {
    this.creationError = null;
    if (this.newUserForm.valid) {
      this.userService
        .createUser(this.newUserForm.value)
        .then((r) => {
          // console.log(r);
          this.ref.close(true);
        })
        .catch((e) => {
          this.creationError =
            "Error creating user. Please ensure that the email doesnt already exist";
        });
    } else {
    }
  }
}
