import {NbMenuItem} from '@nebular/theme';

export const MENU_ITEMS: NbMenuItem[] = [
  {
    title: 'Dashboard',
    icon: 'home-outline',
    link: '',
    home: true,
  },
  {
    title: 'Meters',
    icon: 'flash-outline',
    link: 'meters',
    home: false,
  },
  {
    title: 'Users',
    icon: 'people-outline',
    link: 'users',
    home: false,
  },
  {
    title: 'Transactions',
    icon: 'activity-outline',
    link: 'transactions',
    home: false,
  },
  // {
  //   title: 'USERS',
  //   group: true,
  // },
  // {
  //   title: 'Auth',
  //   icon: 'lock-outline',
  //   children: [
  //     {
  //       title: 'Login',
  //       link: '/auth/login',
  //     },
  //     {
  //       title: 'Register',
  //       link: '/auth/register',
  //     },
  //     {
  //       title: 'Request Password',
  //       link: '/auth/request-password',
  //     },
  //     {
  //       title: 'Reset Password',
  //       link: '/auth/reset-password',
  //     },
  //   ],
  // },
];
