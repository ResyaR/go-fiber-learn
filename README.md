<p align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Logo" height="80"/>
  <img src="https://raw.githubusercontent.com/gofiber/docs/main/static/logo.svg" alt="Fiber Logo" height="80"/>
  <img src="https://gorm.io/assets/logo.svg" alt="GORM Logo" height="80"/>
</p>

<h1 align="center">belajargo</h1>

A simple REST API using Go, Fiber, and GORM (Go ORM). 
The project structure is organized for easy development and professionalism.

---

## Project Structure

- `main.go` — Application entry point
- `models/` — Database models (e.g., User)
- `db/` — Database connection and migration
- `handlers/` — API endpoint handlers
- `routes/` — Application routing

## How to Run

1. Make sure you have Go installed (minimum version 1.16)
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. (Optional) Install [Air](https://github.com/cosmtrek/air) for live reload:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```
4. Run the application:
   - With Air (recommended for development):
     ```bash
     air
     ```
   - Or with Go directly:
     ```bash
     go run main.go
     ```
5. The API will be available at `http://localhost:3000`

## Endpoints

- `GET    /users` — List users
- `POST   /users` — Add user
- `PUT    /users/:id` — Update user
- `DELETE /users/:id` — Delete user

## Main Dependencies

- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/index.html)

---

> This project is suitable for learning the basics of REST API development with Go and best practices for folder organization.
