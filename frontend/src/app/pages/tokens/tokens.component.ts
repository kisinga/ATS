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
import { TokenStatus } from "app/models/token.model";
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
  currentPage = 0;
  displayedColumns: string[] = [
    "meter_number",
    "tokenNumber",
    "apiKey",
    "date",
    "status",
  ];
  @ViewChild(MatPaginator) paginator: MatPaginator;
  tokenStatus = TokenStatus;

  constructor(private tokenService: TokenService) {
    this.getTokens({ limit: 10 }, "cache-first");
  }

  ngAfterViewInit() {
    this.paginator.page
      .pipe(takeUntil(this.comopnentDestroyed))
      .subscribe((page) => {
        if (page.pageIndex === 0 && page.pageIndex > this.currentPage) {
          //  The user had scrolled forwad and is now back to page 1, hence no pagination data
          this.getTokens({ limit: 10 }, "cache-first");
        } else {
          // The user is not on the first page
          if (page.pageIndex > this.currentPage) {
            // User has clicked next
            this.getTokens(
              {
                limit: 10,
                beforeOrAfter: this.tokensPage.pageInfo.endCursor,
              },
              "cache-first"
            );
          } else {
            // User has clicked next
            this.getTokens(
              {
                limit: 10,
                beforeOrAfter: this.tokensPage.pageInfo.startCursor,
                reversed: true,
              },
              "cache-first"
            );
          }
        }
        this.currentPage = page.pageIndex;
      });
  }

  getTokens(params: GetTokensQueryInput, fetchPolicy: FetchPolicy) {
    this.loading = true;
    this.tokenService.getTokens(params, fetchPolicy).then((r) => {
      // console.log(r.data);
      this.tokensPage = r.data.getTokens;
      this.loading = false;
    });
  }

  ngOnDestroy(): void {
    this.comopnentDestroyed.next(true);
  }

  ngOnInit(): void {}
}
