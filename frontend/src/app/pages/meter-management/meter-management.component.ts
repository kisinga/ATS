import { Component, OnInit } from "@angular/core";
import { FetchPolicy } from "@apollo/client/core";
import { NbDialogService } from "@nebular/theme";
import { GetMetersQueryInput } from "app/models/gql/meter.model";
import { Meter } from "app/models/meter.model";
import { MeterService } from "../shared/services/meter.service";
import { TokensComponent } from "../tokens/tokens.component";
import { EditMeterComponent } from "./dialogs/edit-meter/edit-meter.component";
import { NewMeterComponent } from "./dialogs/new-meter/new-meter.component";

@Component({
  templateUrl: "./meter-management.component.html",
  styleUrls: ["./meter-management.component.scss"],
})
export class MeterManagementComponent implements OnInit {
  loading: Boolean = false;
  loadingMeter = "";
  meters: Meter[];
  displayedColumns: string[] = [
    "meter_number",
    "phone",
    "createdby",
    "date",
    "action",
  ];

  constructor(
    private dialogService: NbDialogService,
    private meterService: MeterService
  ) {
    this.getMeters({}, "cache-first");
  }

  ngOnInit(): void {}

  openNewMeterModal() {
    this.dialogService
      .open(NewMeterComponent)
      .onClose.subscribe((refresh: boolean) => {
        if (refresh) {
          this.getMeters({}, "network-only");
        }
      });
  }

  getMeters(params: GetMetersQueryInput, fetchPolicy: FetchPolicy) {
    this.loading = true;
    this.meterService.getMeters(params, fetchPolicy).then((r) => {
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
        this.getMeters({}, "network-only");
      });
  }

  editMeter(meterNumber: string) {
    this.dialogService
      .open(EditMeterComponent, {
        context: { meterNumber },
      })
      .onClose.subscribe((refresh: boolean) => {
        if (refresh) {
          this.getMeters({}, "network-only");
        }
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
        this.getMeters({}, "network-only");
      });
  }
  showTokens(meterNumber: string) {
    this.dialogService.open(TokensComponent, {
      dialogClass: "model-full",

      context: { meterNumber },
    });
  }
}
