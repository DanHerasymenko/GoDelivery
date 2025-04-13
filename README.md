# 📦 GoDelivery — Delivery Management System (Microservices Architecture)

**Author:** Daniil Herasymenko  
**Tech Stack:** Go · Gin · PostgreSQL · Redis · Kafka · Docker · Kubernetes · JWT · CI/CD · Testing

---

## 🚀 Project Description

GoDelivery is a real-time microservice platform that simulates how modern delivery platforms (like Glovo/Uber Eats) operate. It is designed with production-level architecture and best backend practices using Golang.

---

## ⚙️ Architecture Overview

- **Microservices-based** (Auth, Order, Courier, Notify, API Gateway)
- **RESTful APIs** for communication
- **Kafka** for asynchronous event handling
- **JWT authentication** and Redis for session/cache
- **PostgreSQL** as primary storage
- **Docker & Kubernetes** for containerization and orchestration
- **GitLab CI/CD Pipelines** for build, test and deploy
- **Structured logging** via `slog` to stdout and file
- **Testing:** unit + integration with mocks

---

## 🧩 Microservices

| Service        | Description |
|----------------|-------------|
| `AuthService`   | Login, register, JWT issuing and validation |
| `OrderService`  | Order creation, status tracking |
| `CourierService`| Courier registration, location updates (via Redis) |
| `NotifyService` | Consumes Kafka events and sends mock notifications |
| `API Gateway`   | Entry point for clients with JWT validation and routing |

---

## 🛠 Technologies Used

- **Backend:** Go, Gin
- **Auth:** JWT, Redis
- **Databases:** PostgreSQL, Redis
- **Messaging:** Kafka
- **DevOps:** Docker, Kubernetes, GitLab CI
- **Docs:** Swagger (OpenAPI 3.0)
- **Testing:** testify, GoMock

---

## 📦 Project Setup

```bash
# Clone project
git clone https://github.com/yourname/go-delivery-platform.git
cd go-delivery-platform

# Build all services
make build

# Start development environment
docker-compose up --build
