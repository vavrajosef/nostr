import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListNotificationsPageComponent } from './list-notifications-page.component';

describe('ListNotificationsPageComponent', () => {
  let component: ListNotificationsPageComponent;
  let fixture: ComponentFixture<ListNotificationsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListNotificationsPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListNotificationsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
