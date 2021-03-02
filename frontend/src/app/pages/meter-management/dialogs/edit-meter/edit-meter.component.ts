import { Component, Input, OnInit } from "@angular/core";
import { FormGroup, FormControl, Validators } from "@angular/forms";
import { NbDialogRef } from "@nebular/theme";
import { MeterService } from "app/pages/shared/services/meter.service";
import { NewMeterComponent } from "../new-meter/new-meter.component";

@Component({
  selector: "edit-meter",
  templateUrl: "./edit-meter.component.html",
  styleUrls: ["./edit-meter.component.scss"],
})
export class EditMeterComponent implements OnInit {
  updateError = null;
  newMeterForm = new FormGroup({
    meterNumber: new FormControl({ value: "", disabled: true }),
    phone: new FormControl(
      "",
      Validators.compose([Validators.required, Validators.min(10)])
    ),
  });
  @Input() meterNumber: string;

  constructor(
    protected ref: NbDialogRef<NewMeterComponent>,
    private meterService: MeterService
  ) {}

  ngOnInit(): void {
    this.newMeterForm.controls["meterNumber"].setValue(this.meterNumber);
  }

  cancel() {
    this.ref.close();
  }

  submit() {
    this.updateError = null;
    if (this.newMeterForm.valid) {
      this.meterService
        .updateMeter({
          meterNumber: this.meterNumber,
          phone: this.newMeterForm.controls.phone.value,
          location: "",
        })
        .toPromise()
        .then((r) => {
          // console.log(r);
          this.ref.close(true);
        })
        .catch((e) => {
          // console.log(e);
          this.updateError =
            "Error updating Meter. Please ensure that the check your internet connection";
        });
    } else {
    }
  }
}
