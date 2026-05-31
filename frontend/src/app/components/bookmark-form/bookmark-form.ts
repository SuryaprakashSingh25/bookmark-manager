import { Component, EventEmitter, Output } from '@angular/core';
import { BookmarkService } from '../../services/bookmark.service';
import { ToasterService } from '../../services/toaster.service';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-bookmark-form',
  imports: [FormsModule],
  templateUrl: './bookmark-form.html',
  styleUrls: ['./bookmark-form.css'],
})
export class BookmarkForm {
  url='';
  @Output()
  bookmarkCreated=new EventEmitter<void>();

  constructor(
    private bookmarkService: BookmarkService,
    private toaster: ToasterService
  ){}

  submit(){
    if(!this.url.trim()){
      this.toaster.showError('Please enter a valid URL for the bookmark.');
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
        this.toaster.showError('Unable to create bookmark. Please try again.');
      }
    });
  }
}
