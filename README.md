# 🔖 Bookmark Manager

A full-stack distributed application to save, manage, and preview bookmarks with a modern microservices architecture.

---

## 🌐 Live Demo

* Frontend: https://bookmark-frontend-rho.vercel.app/
* Backend API: https://bookmark-backend-h81i.onrender.com/bookmarks

---

## 🧠 Overview

This project demonstrates a production-style system built using **Go, Angular, gRPC, and PostgreSQL**, focusing on scalability, clean architecture, and real-world backend patterns.

Users can:

* Add bookmarks with title & description
* View bookmarks in a clean card-based UI
* Delete bookmarks
* Preview metadata using a separate gRPC service

---

## 🏗️ Architecture

```text
Angular (Vercel)
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
* gRPC
* PostgreSQL (Supabase)
* Docker

### Frontend

* Angular
* Responsive UI (Card-based layout)

### Infrastructure

* Render (Backend Deployment)
* Vercel (Frontend Deployment)

---

## 🔑 Key Features

* Microservices architecture (REST + gRPC)
* Distributed system simulation with multiple services
* Clean API design with proper layering
* Asynchronous service communication via gRPC
* Production-style deployment across multiple platforms
* Database connection retry mechanism for reliability

---

## 🧩 System Design Highlights

* Separate **API service** and **Preview service** for better modularity
* gRPC used for efficient internal communication
* Stateless backend services
* Designed for horizontal scalability
* Handles network failures with retry logic

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

| Method | Endpoint       | Description     |
| ------ | -------------- | --------------- |
| GET    | /bookmarks     | Fetch bookmarks |
| POST   | /bookmarks     | Create bookmark |
| DELETE | /bookmarks/:id | Delete bookmark |

---


## 🚀 Future Improvements

* User authentication (multi-user support)
* Bookmark categorization / tagging
* Rate limiting
* Caching layer (Redis)
* CI/CD pipeline

---

## 💡 What I Learned

* Building microservices in Go
* gRPC communication between services
* Real-world deployment challenges (Docker, networking, CORS)
* Handling cloud limitations (IPv4/IPv6, cold starts)
* Structuring scalable backend systems

---

## 📬 Contact

* LinkedIn: https://www.linkedin.com/in/suryaprakash-singh/
* Email: [singh.suraj1025@gmail.com](mailto:singh.suraj1025@gmail.com)

---
