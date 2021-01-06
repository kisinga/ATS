import {Component} from '@angular/core';

@Component({
  selector: 'ngx-footer',
  styleUrls: ['./footer.component.scss'],
  template: `
    <span class="created-by">
      Created with ♥ by <b><a href="https://github.com/kisinga" target="_blank">Kisinga</a></b> 2020
    </span>
    <div class="socials">
      <a href="https://github.com/kisinga" target="_blank" class="ion ion-social-github"></a>
      <a href="https://stackoverflow.com/users/5927361/kisinga/" target="_blank" class="ion ion-social-stackoverflow"></a>
      <a href="https://www.linkedin.com/in/stevekisinga/" target="_blank" class="ion ion-social-linkedin"></a>
    </div>
  `,
})
export class FooterComponent {
}
