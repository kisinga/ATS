import {async, ComponentFixture, TestBed} from '@angular/core/testing';

import {MeterManagementComponent} from './meter-management.component';

describe('MeterManagementComponent', () => {
  let component: MeterManagementComponent;
  let fixture: ComponentFixture<MeterManagementComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [MeterManagementComponent]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MeterManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
