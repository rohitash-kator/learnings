# Event Booking

A REST API for event management and registration built with **Go** and **Gin**. Users can sign up, log in with JWT, create and manage events, and register for or cancel registration on events.

## Tech Stack

- **Go** (1.25+)
- **Gin** – HTTP web framework
- **SQLite** – Database (`api.db`)
- **JWT** – Authentication (Bearer token)
- **golang.org/x/crypto** – Password hashing (bcrypt)

## Project Structure

```
event-booking/
├── main.go              # App entry, DB init, route registration
├── db/
│   └── db.go            # SQLite connection, table creation (users, events, registrations)
├── routes/
│   └── routes.go        # Route definitions (public vs protected)
├── controllers/
│   ├── users.go         # Signup, Login
│   ├── events.go        # Events CRUD
│   └── registrations.go # Register / cancel for event
├── models/
│   ├── user.go          # User model, save, validate credentials
│   └── event.go         # Event model, CRUD, register/cancel
├── middlewares/
│   └── auth.go          # JWT auth middleware
└── utils/
    ├── jwt.go           # Token generation and verification
    └── hash.go          # Password hashing
```

## Prerequisites

- Go 1.25 or later
- CGO enabled (required for SQLite driver)

## Setup & Run

```bash
# Install dependencies
go mod download

# Run the server (creates api.db if missing)
go run main.go
```

Server runs at **http://localhost:3000**.

## API Endpoints

### Auth (no token)

| Method | Path     | Description        |
|--------|----------|--------------------|
| POST   | `/signup` | Create user        |
| POST   | `/login`  | Login, returns JWT |

### Events (public)

| Method | Path          | Description     |
|--------|---------------|-----------------|
| GET    | `/events`     | List all events |
| GET    | `/events/:id` | Get one event   |

### Events & registrations (protected – require `Authorization: <token>`)

| Method | Path                    | Description              |
|--------|-------------------------|--------------------------|
| POST   | `/events`               | Create event             |
| PUT    | `/events/:id`           | Update event (owner only)|
| DELETE | `/events/:id`           | Delete event (owner only)|
| POST   | `/events/:id/register`  | Register for event       |
| DELETE | `/events/:id/register`  | Cancel registration      |

## Example Usage

### 1. Sign up

```bash
curl -X POST http://localhost:3000/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'
```

### 2. Login

```bash
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"secret123"}'
```

Response includes a `token`; use it in the `Authorization` header for protected routes.

### 3. Create event (authenticated)

```bash
curl -X POST http://localhost:3000/events \
  -H "Content-Type: application/json" \
  -H "Authorization: YOUR_JWT_TOKEN" \
  -d '{
    "name":"Meetup",
    "description":"Dev meetup",
    "location":"City Hall",
    "date_time":"2025-03-01T18:00:00Z"
  }'
```

### 4. Register for an event

```bash
curl -X POST http://localhost:3000/events/1/register \
  -H "Authorization: YOUR_JWT_TOKEN"
```

### 5. Cancel registration

```bash
curl -X DELETE http://localhost:3000/events/1/register \
  -H "Authorization: YOUR_JWT_TOKEN"
```

## Database

- **File:** `api.db` (SQLite, created on first run)
- **Tables:** `users`, `events`, `registrations`
- Passwords are stored hashed (bcrypt). JWT is used for session/auth.

## License

MIT
