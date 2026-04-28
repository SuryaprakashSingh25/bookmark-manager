import { Component, OnInit } from '@angular/core';
import { Bookmark } from '../../models/bookmark';
import { BookmarkService } from '../../services/bookmark.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-bookmark-list',
  imports: [CommonModule],
  templateUrl: './bookmark-list.html',
  styleUrl: './bookmark-list.css',
})
export class BookmarkList implements OnInit{
  bookmarks: Bookmark[]=[];

  constructor(private bookmarkService: BookmarkService){}

  ngOnInit(): void {
    this.loadBookmarks();
  }

  loadBookmarks(){
    this.bookmarkService.getBookmarks()
    .subscribe({
      next: (data) => {
        this.bookmarks=data;
      },
      error: (err) => {
        console.error(err);
      }
    });
  }

}
