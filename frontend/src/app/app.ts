import { Component, signal, ViewChild } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { BookmarkForm } from './components/bookmark-form/bookmark-form';
import { BookmarkList } from './components/bookmark-list/bookmark-list';

@Component({
  selector: 'app-root',
  imports: [BookmarkForm,BookmarkList],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  @ViewChild(BookmarkList)
  bookmarkList!:BookmarkList;

  refreshBookmarks(){
    this.bookmarkList.loadBookmarks();
  }
}
