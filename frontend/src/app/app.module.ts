/**
 * @license
 * Copyright Akveo. All Rights Reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */
import { BrowserModule } from "@angular/platform-browser";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { NgModule } from "@angular/core";
import { HttpClientModule } from "@angular/common/http";
// import { CoreModule } from './@core/core.module';
import { ThemeModule } from "./@theme/theme.module";
import { AppComponent } from "./app.component";
import { AppRoutingModule } from "./app-routing.module";
import { AngularFireAuthGuard } from "@angular/fire/auth-guard";
import {
  NbButtonModule,
  NbCardModule,
  NbDatepickerModule,
  NbDialogModule,
  NbIconModule,
  NbLayoutModule,
  NbMenuModule,
  NbSidebarModule,
  NbToastrModule,
  NbWindowModule,
} from "@nebular/theme";
import { environment } from "../environments/environment";
import { AngularFireModule } from "@angular/fire";
import { LoginComponent } from "./auth/login/login.component";
import { UnauthorisedComponent } from "./auth/unauthorised/unauthorised.component";
import { NbEvaIconsModule } from "@nebular/eva-icons";
import { GraphQLModule } from "./graphql.module";
import { APIKeySubscription } from "./models/gql/api-key.model";

@NgModule({
  declarations: [AppComponent, LoginComponent, UnauthorisedComponent],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    GraphQLModule,
    HttpClientModule,
    AppRoutingModule,
    NbSidebarModule.forRoot(),
    NbMenuModule.forRoot(),
    NbDatepickerModule.forRoot(),
    NbDialogModule.forRoot(),
    NbWindowModule.forRoot(),
    NbToastrModule.forRoot(),
    NbLayoutModule,
    NbCardModule,
    NbButtonModule,
    NbEvaIconsModule,
    NbIconModule,
    AngularFireModule.initializeApp(environment.firebase),
    ThemeModule.forRoot(),
    NbIconModule,
  ],
  exports: [GraphQLModule],
  providers: [
    // ...
    AngularFireAuthGuard,
  ],
  bootstrap: [AppComponent],
})
export class AppModule {
  constructor() {}
}
