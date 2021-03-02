import {Component, OnInit} from "@angular/core";
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {NbDialogRef} from "@nebular/theme";
import {MeterService} from "app/pages/shared/services/meter.service";

@Component({
  templateUrl: "./new-meter.component.html",
  styleUrls: ["./new-meter.component.scss"],
})
export class NewMeterComponent implements OnInit {
  creationError = null;
  newMeterForm = new FormGroup({
    meterNumber: new FormControl(
      "",
      Validators.compose([Validators.required, Validators.min(5)])
    ),
    phone: new FormControl(
      "",
      Validators.compose([Validators.required, Validators.min(10)])
    ),
  });

  constructor(
    protected ref: NbDialogRef<NewMeterComponent>,
    private meterService: MeterService
  ) {}

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
          // console.log(e);
          this.creationError =
            "Error creating Meter. Please ensure that the Meter-Number doesnt already exist";
        });
    } else {
    }
  }
}
