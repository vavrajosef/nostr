import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { HomepageComponent } from './home-page/home-page.component';
import { ProfilePageComponent } from './profile-page/profile-page.component';
import { AppearanceSettingsPageComponent } from './appearance-settings-page/appearance-settings-page.component';
import { AccountSettingsPageComponent } from './account-settings-page/account-settings-page.component';
import { BackupSettingsPageComponent } from './backup-settings-page/backup-settings-page.component';
import { CreateShortTextNoteFormComponent } from './create-short-text-note-form/create-short-text-note-form.component';
import { ListNotificationsPageComponent } from './list-notifications-page/list-notifications-page.component';
import { NetworkSettingsPageComponent } from './network-settings-page/network-settings-page.component';
import { SignInPageComponent } from './sign-in-page/sign-in-page.component';
import { SignUpPageComponent } from './sign-up-page/sign-up-page.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { NotFoundPageComponent } from './not-found-page/not-found-page.component';

@NgModule({
  declarations: [
    AppComponent,
    HomepageComponent,
    ProfilePageComponent,
    AppearanceSettingsPageComponent,
    AccountSettingsPageComponent,
    BackupSettingsPageComponent,
    CreateShortTextNoteFormComponent,
    ListNotificationsPageComponent,
    NetworkSettingsPageComponent,
    SignInPageComponent,
    SignUpPageComponent,
    LandingPageComponent,
    NotFoundPageComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [

  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
