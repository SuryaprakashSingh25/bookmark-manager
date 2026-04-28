import { Component, EventEmitter, Output } from '@angular/core';
import { BookmarkService } from '../../services/bookmark.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-bookmark-form',
  imports: [FormsModule],
  templateUrl: './bookmark-form.html',
  styleUrl: './bookmark-form.css',
})
export class BookmarkForm {
  url='';
  @Output()
  bookmarkCreated=new EventEmitter<void>();

  constructor(
    private bookmarkService: BookmarkService
  ){}

  submit(){
    if(!this.url.trim()){
      return;
    }

    this.bookmarkService.createBookmark(this.url)
    .subscribe({
      next: () => {
        this.url = '';
        this.bookmarkCreated.emit();
      },
      error: (err) => {
        console.error(err);
      }
    });
  }
}
