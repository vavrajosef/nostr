import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators';

const defaultHostname = 'localhost'
const defaultPort = '4318'


export interface GetHealthRequest {

}

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

  GetHealth(r: GetHealthRequest): Observable<GetHealthResponse> {
    return this.httpClient.get<GetHealthResponse>(`http://${defaultHostname}:${defaultPort}/health`)
  }
}
