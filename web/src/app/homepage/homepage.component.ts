import { Component, OnInit } from '@angular/core';
import { AppService } from '../app.service';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {
  status = ''
  timestamp = -1

  constructor(
    private readonly appService: AppService,
  ) { }

  ngOnInit() {
    this.appService.GetHealth().subscribe((res) => {
      this.status = res.status
      this.timestamp = res.timestamp
    })
  }
}
