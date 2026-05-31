import { Component, OnInit } from '@angular/core';
import { Bookmark } from '../../models/bookmark';
import { BookmarkService } from '../../services/bookmark.service';
import { ToasterService } from '../../services/toaster.service';
import { CommonModule } from '@angular/common';
import { BookmarkForm } from '../bookmark-form/bookmark-form';

@Component({
  selector: 'app-bookmark-list',
  imports: [CommonModule, BookmarkForm],
  templateUrl: './bookmark-list.html',
  styleUrls: ['./bookmark-list.css'],
})
export class BookmarkList implements OnInit{
  bookmarks: Bookmark[]=[];

  constructor(
    private bookmarkService: BookmarkService,
    private toaster: ToasterService
  ){}

  ngOnInit(): void {
    this.loadBookmarks();
  }

  loadBookmarks(){
    this.bookmarkService.getBookmarks()
    .subscribe({
      next: (data) => {
        this.bookmarks = data;
      },
      error: (err) => {
        console.error(err);
        this.toaster.showError('Could not load bookmarks. Please try again.');
      }
    });
  }

  deleteBookmark(id:number){
    this.bookmarkService.deleteBookmark(id)
    .subscribe({
      next: () => {
        this.bookmarks = this.bookmarks.filter(
          b => b.id !== id
        );
      },
      error: (err) => {
        console.error(err);
        this.toaster.showError('Failed to delete bookmark. Please try again.');
      }
    });
  }

}
