import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

export type ToastType = 'success' | 'error' | 'info';

export interface ToastMessage {
  id: number;
  message: string;
  type: ToastType;
  duration: number;
}

@Injectable({
  providedIn: 'root',
})
export class ToasterService {
  private toastQueue: ToastMessage[] = [];
  private toastSubject = new BehaviorSubject<ToastMessage[]>([]);
  toasts$ = this.toastSubject.asObservable();
  private nextId = 1;

  show(message: string, type: ToastType = 'info', duration = 5000) {
    const toast: ToastMessage = {
      id: this.nextId++,
      message,
      type,
      duration,
    };

    this.toastQueue = [...this.toastQueue, toast];
    this.toastSubject.next(this.toastQueue);

    window.setTimeout(() => this.dismiss(toast.id), duration);
  }

  showError(message: string, duration = 5000) {
    this.show(message, 'error', duration);
  }

  showSuccess(message: string, duration = 4000) {
    this.show(message, 'success', duration);
  }

  showInfo(message: string, duration = 4000) {
    this.show(message, 'info', duration);
  }

  dismiss(id: number) {
    this.toastQueue = this.toastQueue.filter((toast) => toast.id !== id);
    this.toastSubject.next(this.toastQueue);
  }
}
