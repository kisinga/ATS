import { ExtraOptions, RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';
import {
  NbAuthComponent,
  NbLoginComponent,
  NbLogoutComponent,
} from '@nebular/auth';
import { AngularFireAuthGuard, hasCustomClaim, redirectUnauthorizedTo, redirectLoggedInTo } from '@angular/fire/auth-guard';

const redirectUnauthorizedToLogin = () => redirectUnauthorizedTo(['auth/login']);
const redirectLoggedInToDashboard = () => redirectLoggedInTo(['']);


export const routes: Routes = [
  {
    path: '',
    canActivate: [AngularFireAuthGuard], // here we tell Angular to check the access with our AuthGuard
    data: { authGuardPipe: redirectUnauthorizedToLogin },
    loadChildren: () => import('./pages/pages.module')
      .then(m => m.PagesModule),
  },
  {
    path: 'auth',
    component: NbAuthComponent,
    canActivate: [AngularFireAuthGuard], data: { authGuardPipe: redirectLoggedInToDashboard },
    children: [
      {
        path: '',
        redirectTo: 'login',
        pathMatch: 'full',
      },
      {
        path: 'login',
        component: NbLoginComponent,
      },

      {
        path: 'logout',
        component: NbLogoutComponent,
      },
      { path: '**', redirectTo: 'login' },
    ],
  },
  { path: '**', redirectTo: 'dashboard' },
];

const config: ExtraOptions = {
  useHash: false,
};

@NgModule({
  imports: [RouterModule.forRoot(routes, config)],
  exports: [RouterModule],
})
export class AppRoutingModule {
}
