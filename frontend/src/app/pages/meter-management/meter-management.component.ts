import { Component, OnInit } from "@angular/core";
import { NbDialogService } from "@nebular/theme";
import { GetMetersQueryInput } from "app/models/gql/meter.model";
import { Meter } from "app/models/meter.model";
import { MeterService } from "../shared/services/meter.service";
import { NewMeterComponent } from "./dialogs/new-meter/new-meter.component";

@Component({
  templateUrl: "./meter-management.component.html",
  styleUrls: ["./meter-management.component.scss"],
})
export class MeterManagementComponent implements OnInit {
  constructor(
    private dialogService: NbDialogService,
    private meterService: MeterService
  ) {
    this.getMeters({});
  }
  loading: Boolean = false;
  loadingMeter = "";
  meters: Meter[];
  displayedColumns: string[] = ["meter_number", "createdby", "date", "delete"];

  ngOnInit(): void {}
  openNewMeterModal() {
    this.dialogService
      .open(NewMeterComponent)
      .onClose.subscribe((refresh: boolean) => {
        if (refresh) {
          this.getMeters({});
        }
      });
  }
  getMeters(params: GetMetersQueryInput) {
    this.loading = true;
    this.meterService.getMeters(params).then((r) => {
      this.meters = r.data.meters.data;
      this.loading = false;
    });
  }
  enableMeter(meterNumber: string) {
    this.loadingMeter = meterNumber;
    this.meterService
      .enableMeter(meterNumber)
      .toPromise()
      .then((t) => {
        if (this.loadingMeter === meterNumber) {
          this.loadingMeter = "";
        }
        this.getMeters({});
      });
  }
  disableMeter(meterNumber: string) {
    this.loadingMeter = meterNumber;
    this.meterService
      .disableMeter(meterNumber)
      .toPromise()
      .then((t) => {
        if (this.loadingMeter === meterNumber) {
          this.loadingMeter = "";
        }
        this.getMeters({});
      });
  }
}
