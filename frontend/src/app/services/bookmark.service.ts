import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Bookmark } from '../models/bookmark';

@Injectable({
  providedIn: 'root',
})
export class BookmarkService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http:HttpClient){}

  getBookmarks(): Observable<Bookmark[]> {
    return this.http.get<Bookmark[]>(
      `${this.apiUrl}/bookmarks`
    );
  }

  createBookmark(url: string): Observable<Bookmark> {
    return this.http.post<Bookmark>(
      `${this.apiUrl}/bookmarks`,
      {url}
    );
  }

}
