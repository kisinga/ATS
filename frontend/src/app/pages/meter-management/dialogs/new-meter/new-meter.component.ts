import { Component, OnInit } from "@angular/core";
import { FormGroup, FormControl, Validators } from "@angular/forms";
import { NbDialogRef } from "@nebular/theme";
import { MeterService } from "app/pages/shared/services/meter.service";
import { UserService } from "app/pages/shared/services/user.service";

@Component({
  templateUrl: "./new-meter.component.html",
  styleUrls: ["./new-meter.component.scss"],
})
export class NewMeterComponent implements OnInit {
  constructor(
    protected ref: NbDialogRef<NewMeterComponent>,
    private meterService: MeterService
  ) {}
  creationError = null;
  newMeterForm = new FormGroup({
    meterNumber: new FormControl(
      "",
      Validators.compose([Validators.required, Validators.min(5)])
    ),
  });

  ngOnInit(): void {}

  cancel() {
    this.ref.close();
  }
  submit() {
    this.creationError = null;
    if (this.newMeterForm.valid) {
      this.meterService
        .createMeter(this.newMeterForm.value)
        .then((r) => {
          // console.log(r);
          this.ref.close(true);
        })
        .catch((e) => {
          console.log(e);
          this.creationError =
            "Error creating Meter. Please ensure that the email doesnt already exist";
        });
    } else {
    }
  }
}
