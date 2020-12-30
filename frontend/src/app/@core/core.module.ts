import { ModuleWithProviders, NgModule, Optional, SkipSelf } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NbAuthJWTToken, NbAuthModule, NbDummyAuthStrategy, NbPasswordAuthStrategy } from '@nebular/auth';
// import { NbSecurityModule, NbRoleProvider } from '@nebular/security';
import { of as observableOf } from 'rxjs';

import { throwIfAlreadyLoaded } from './module-import-guard';
import { AnalyticsService, SeoService } from './utils';
import { UserData } from './data/users';
import { UserService } from './mock/users.service';
import { MockDataModule } from './mock/mock-data.module';

const socialLinks = [

];

const DATA_SERVICES = [
  { provide: UserData, useClass: UserService },
];

// export class NbSimpleRoleProvider extends NbRoleProvider {
//   getRole() {
//     // here you could provide any role based on any auth flow
//     return observableOf('guest');
//   }
// }

export const NB_CORE_PROVIDERS = [
  ...MockDataModule.forRoot().providers,
  ...DATA_SERVICES,
  ...NbAuthModule.forRoot({
    // strategies: [
    //   NbDummyAuthStrategy.setup({
    //     name: 'email',
    //     delay: 3000,
    //   }),
    // ],
    // forms: {
    //   login: {
    //     socialLinks: socialLinks,
    //   },
    //   register: {
    //     socialLinks: socialLinks,
    //   },
    // },
    strategies: [
      NbPasswordAuthStrategy.setup({
        name: 'email',
        token: {
          class: NbAuthJWTToken,
        },
        baseEndpoint: '',
        login: {
          alwaysFail: false,
          method: 'post',
          requireValidToken: true,
          redirect: {
            success: '/',
            failure: null,
          },
          defaultErrors: ['Login/Email combination is not correct, please try again.'],
          defaultMessages: ['You have been successfully logged in.'],
          endpoint: '/api/auth/login',
        },
        logout: {
          alwaysFail: false,
          method: 'delete',
          redirect: {
            success: '/',
            failure: null,
          },
          defaultErrors: ['Something went wrong, please try again.'],
          defaultMessages: ['You have been successfully logged out.'],
          endpoint: '/api/auth/sign-out',
        },
        refreshToken: {
          endpoint: 'refresh-token',
          method: 'post',
          requireValidToken: true,
          redirect: {
            success: null,
            failure: null,
          },
          defaultErrors: ['Something went wrong, please try again.'],
          defaultMessages: ['Your token has been successfully refreshed.'],
        },

      }),
    ],
    forms: {
      redirectDelay: 500, // delay before redirect after a successful login, while success message is shown to the user
      strategy: 'email',  // strategy id key.
      rememberMe: false,   // whether to show or not the `rememberMe` checkbox
      showMessages: {     // show/not show success/error messages
        success: true,
        error: true,
      },
      socialLinks: socialLinks, // social links at the bottom of a page
      logout: {
        redirectDelay: 500,
        strategy: 'email',
      },
      validation: {
        password: {
          required: true,
          minLength: 5,
          maxLength: 50,
        },
        email: {
          required: true,
        },
      },
    },
  }).providers,

  // NbSecurityModule.forRoot({
  //   accessControl: {
  //     guest: {
  //       view: '*',
  //     },
  //     user: {
  //       parent: 'guest',
  //       create: '*',
  //       edit: '*',
  //       remove: '*',
  //     },
  //   },
  // }).providers,

  // {
  //   provide: NbRoleProvider, useClass: NbSimpleRoleProvider,
  // },
  AnalyticsService,
  SeoService,
];

@NgModule({
  imports: [
    CommonModule,
  ],
  exports: [
    NbAuthModule,
  ],
  declarations: [],
})
export class CoreModule {
  constructor(@Optional() @SkipSelf() parentModule: CoreModule) {
    throwIfAlreadyLoaded(parentModule, 'CoreModule');
  }

  static forRoot(): ModuleWithProviders<CoreModule> {
    return {
      ngModule: CoreModule,
      providers: [
        ...NB_CORE_PROVIDERS,
      ],
    };
  }
}
