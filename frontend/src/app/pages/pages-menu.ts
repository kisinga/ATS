import { NbMenuItem } from "@nebular/theme";

export const MENU_ITEMS: NbMenuItem[] = [
  {
    title: "Dashboard",
    icon: "home-outline",
    link: "/dashboard",
    home: true,
  },
  {
    title: "Meters",
    icon: "flash-outline",
    link: "/meters",
    home: false,
  },
  {
    title: "Users",
    icon: "people-outline",
    link: "/users",
    home: false,
  },
  {
    title: "Transactions",
    icon: "activity-outline",
    link: "/transactions",
    home: false,
  },
];
