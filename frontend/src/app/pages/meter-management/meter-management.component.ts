import { Component, OnInit } from "@angular/core";
import { Meter } from "app/models/meter.model";

@Component({
  templateUrl: "./meter-management.component.html",
  styleUrls: ["./meter-management.component.scss"],
})
export class MeterManagementComponent implements OnInit {
  constructor() {}
  loading: Boolean = false;
  loadingMeter = "";
  meters: Meter[];
  displayedColumns: string[] = [
    "meter_number",
    "owner",
    "phone",
    "date",
    "delete",
  ];

  ngOnInit(): void {}
  openNewMeterModal() {}
  enableMeter() {}
  disableMeter() {}
}
