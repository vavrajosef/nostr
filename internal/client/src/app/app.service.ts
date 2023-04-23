import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

const defaultHostname = '0.0.0.0'
const defaultPort = '4317'

export interface GetHealthResponse {
  status: string
  timestamp: number
}

@Injectable({
  providedIn: 'root'
})
export class AppService {

  constructor(
    private readonly httpClient: HttpClient,
  ) { }

  GetHealth(): Observable<GetHealthResponse> {
    return this.httpClient.get<GetHealthResponse>(`http://${defaultHostname}:${defaultPort}/health`)
  }
}
