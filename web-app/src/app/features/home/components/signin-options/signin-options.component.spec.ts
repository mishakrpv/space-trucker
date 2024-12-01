import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SigninOptionsComponent } from './signin-options.component';

describe('SigninOptionsComponent', () => {
  let component: SigninOptionsComponent;
  let fixture: ComponentFixture<SigninOptionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SigninOptionsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SigninOptionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
