# 🔖 Bookmark Manager

A full-stack distributed bookmark management application built with **Go, Angular, gRPC, PostgreSQL, and JWT Authentication**. The project demonstrates microservices architecture, secure user authentication, and modern frontend/backend development practices.

---

## 🌐 Live Demo

* Frontend: https://bookmark-frontend-rho.vercel.app/
* Backend API: https://bookmark-backend-h81i.onrender.com/bookmarks

---

## 🧠 Overview

Bookmark Manager is a production-style distributed application that allows users to securely save and manage bookmarks through an intuitive interface.

The application follows a microservices architecture where a REST API communicates with a dedicated gRPC service for bookmark preview generation.

Users can:

* Create an account and securely log in
* Add bookmarks with title & description
* View bookmarks in a clean card-based UI
* Delete bookmarks
* Preview bookmark metadata through a gRPC service
* Access protected features using JWT-based authentication

---

## 🏗️ Architecture

```text
Angular (Vercel)
        ↓
JWT Authentication
        ↓
Go REST API (Render)
        ↓
gRPC Preview Service
        ↓
PostgreSQL (Supabase)
```

---

## ⚙️ Tech Stack

### Backend

* Go (Gin)
* JWT Authentication
* gRPC
* PostgreSQL (Supabase)
* Docker

### Frontend

* Angular
* Responsive UI (Card-based layout)
* Route Guards

### Infrastructure

* Render (Backend Deployment)
* Vercel (Frontend Deployment)

---

## 🔑 Key Features

* JWT-based Authentication & Authorization
* User Signup, Login & Logout
* Bookmark CRUD Operations
* Microservices architecture (REST + gRPC)
* Asynchronous service communication via gRPC
* Production-style deployment across
* Database Connection Retry Mechanism

---

## 🧩 System Design Highlights

* Separate **API service** and **Preview service** for better modularity
* gRPC used for efficient internal communication
* Stateless backend services
* JWT-based authentication workflow
* Clear separation of concerns
* Database retry mechanism for reliability

---

## 🛠️ Local Setup

### 1. Clone the repository

```bash
git clone https://github.com/your-username/bookmark-manager.git
cd bookmark-manager
```

---

### 2. Backend Setup

```bash
cd api
go mod tidy
go run main.go
```

---

### 3. gRPC Service

```bash
cd preview-service
go run main.go
```

---

### 4. Frontend Setup

```bash
cd frontend
npm install
ng serve
```

---

## 🧪 API Endpoints

### Authentication
| Method | Endpoint       | Description       |
| ------ | -------------- | ----------------- |
| POST   | /signup        | Register new user |
| POST   | /login         | Authenticate user |

### Bookmarks
| Method | Endpoint       | Description     |
| ------ | -------------- | --------------- |
| GET    | /bookmarks     | Fetch bookmarks |
| POST   | /bookmarks     | Create bookmark |
| DELETE | /bookmarks/:id | Delete bookmark |

---


## 🚀 Future Improvements

* Password Reset Flow
* Bookmark categorization / tagging
* Bookmark search and filtering
* Pagination and sort functionality

---

## 💡 What I Learned

* Building microservices in Go
* gRPC communication between services
* JWT Authentication & Authorization
* Angular Route Guards & Form Validation
* Docker-Based Development
* Handling Networking, CORS, and Production Challenges
* Designing Scalable Backend Architectures

---

## 📬 Contact

* LinkedIn: https://www.linkedin.com/in/suryaprakash-singh/
* Email: [singh.suraj1025@gmail.com](mailto:singh.suraj1025@gmail.com)

---

⭐ If you found this project interesting, consider giving the repository a star.
