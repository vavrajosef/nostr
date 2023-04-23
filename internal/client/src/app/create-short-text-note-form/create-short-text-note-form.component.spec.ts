import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateShortTextNoteFormComponent } from './create-short-text-note-form.component';

describe('CreateShortTextNoteFormComponent', () => {
  let component: CreateShortTextNoteFormComponent;
  let fixture: ComponentFixture<CreateShortTextNoteFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateShortTextNoteFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateShortTextNoteFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
