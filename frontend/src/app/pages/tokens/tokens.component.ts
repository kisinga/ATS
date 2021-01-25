import {
  AfterViewInit,
  Component,
  OnDestroy,
  OnInit,
  ViewChild,
} from "@angular/core";
import { MatPaginator } from "@angular/material/paginator";
import {
  GetTokensQueryInput,
  TokensQueryModel,
} from "app/models/gql/token.model";
import { FetchPolicy } from "@apollo/client/core";
import { TokenService } from "../shared/services/token.service";
import { ReplaySubject } from "rxjs";
import { takeUntil } from "rxjs/operators";

@Component({
  templateUrl: "./tokens.component.html",
  styleUrls: ["./tokens.component.scss"],
})
export class TokensComponent implements OnInit, AfterViewInit, OnDestroy {
  tokensPage: TokensQueryModel;
  loading: Boolean = false;
  comopnentDestroyed: ReplaySubject<boolean> = new ReplaySubject<boolean>();

  displayedColumns: string[] = [
    "meter_number",
    "tokenNumber",
    "apiKey",
    "date",
    "status",
  ];
  @ViewChild(MatPaginator) paginator: MatPaginator;

  constructor(private tokenService: TokenService) {
    this.getTokens({ limit: 2 }, "cache-first");
  }
  ngAfterViewInit() {
    this.paginator.page
      .pipe(takeUntil(this.comopnentDestroyed))
      .subscribe(() => {
        this.getTokens(
          {
            limit: 10,
            after: this.tokensPage.pageInfo.endCursor,
          },
          "network-only"
        );
      });
  }
  getTokens(params: GetTokensQueryInput, fetchPolicy: FetchPolicy) {
    this.loading = true;
    this.tokenService.getTokens(params, fetchPolicy).then((r) => {
      console.log(r.data);
      this.tokensPage = r.data.getTokens;
      this.loading = false;
    });
  }
  ngOnDestroy(): void {
    this.comopnentDestroyed.next(true);
  }
  ngOnInit(): void {}
}
