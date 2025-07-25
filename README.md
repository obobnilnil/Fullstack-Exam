# Live Order Monitoring System

A fullstack real-time order monitoring dashboard designed for internal staff.  
Built with microservices architecture using Go, NestJS, PostgreSQL, Redis, WebSocket, and Next.js.

---

## 🛠️ Setup Instructions

### Prerequisites
- Docker & Docker Compose
- Node.js (for local development only)
- Go (for local backend testing)

### Run the System
```bash
# Go to project root
cd answer2/live-order-monitoring

# Start all services
docker compose up --build

```

Access Points
Frontend (Next.js): http://localhost:3001

API Gateway (NestJS): http://localhost:3000

pgAdmin: http://localhost:9977

PostgreSQL: localhost:5611

Redis: localhost:6379


🧱 Architecture Diagram

                      +----------------+
                      |   Frontend     |
                      | (Next.js)      |
                      +-------+--------+
                              |
                              v
                    +---------+----------+
                    |      Gateway       |
                    |    (NestJS)        |
                    +--+------------+----+
                       |            |
                       v            v
                 +-----+--+     +---+------+
                 | Orders |     |  Users   |
                 | (Go)   |     |  (Go)     |
                 +--------+     +----------+
                       |
                       v
                 +-----------+
                 | PostgreSQL |
                 +-----------+

    WebSocket ⟷ Redis Pub/Sub ⟷ Order Service

🧩 Service Responsibilities
Service	Tech Stack	Responsibilities
Frontend	Next.js	Staff dashboard with real-time order updates
Gateway	NestJS	Auth, routing, WebSocket handling
Orders	Go	Order creation, status update, Redis publish
Users	Go	Staff/admin login, JWT issuance
PostgreSQL	DB	Stores orders, users
Redis	Redis	WebSocket Pub/Sub events for real-time sync

📡 WebSocket Use Case
When a new order is created or updated via orders service

The service publishes an event to Redis

The gateway service (NestJS) subscribes to Redis channel

Gateway emits the event via WebSocket to connected frontend clients

Frontend dashboard reflects the update in real-time


📁 Folder Structure
bash
Copy
Edit
answer2/live-order-monitoring/
├── frontend/       # Next.js frontend
├── gateway/        # NestJS API Gateway
├── services/
│   ├── orders/     # Go order service
│   └── users/      # Go user service
├── .env.example    # Sample environment file
├── docker-compose.yml
