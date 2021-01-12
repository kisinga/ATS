import { NgModule } from "@angular/core";
import { NbEvaIconsModule } from "@nebular/eva-icons";
import {
  NbButtonModule,
  NbCardModule,
  NbIconModule,
  NbLayoutModule,
} from "@nebular/theme";

import { ThemeModule } from "../../@theme/theme.module";
import { DashboardComponent } from "./dashboard.component";

@NgModule({
  imports: [
    NbCardModule,
    NbButtonModule,
    NbEvaIconsModule,
    NbIconModule,
    ThemeModule,
    NbLayoutModule,
  ],
  declarations: [DashboardComponent],
})
export class DashboardModule {}
