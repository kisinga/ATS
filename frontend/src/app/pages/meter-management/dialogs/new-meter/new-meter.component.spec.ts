import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NewMeterComponent } from './new-meter.component';

describe('NewMeterComponent', () => {
  let component: NewMeterComponent;
  let fixture: ComponentFixture<NewMeterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NewMeterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NewMeterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
