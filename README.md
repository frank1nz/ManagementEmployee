# ManagementEmployee

A backend project written in Go using [Fiber](https://github.com/gofiber/fiber) and [GORM](https://gorm.io/) that provides **authentication for admins** and **management functionality for employees**.

## 📦 Features

- 🔐 JWT-based admin authentication
- 👥 Employee management (CRUD)
- 🗃️ MySQL database integration via GORM
- 🌐 RESTful API structure with Fiber
- 📂 .env config support using `godotenv`

---

## 🛠️ Technologies Used

- Go 1.23+
- Fiber v2
- GORM
- JWT (via `github.com/golang-jwt/jwt`)
- MySQL

---



## 📁 Project Structure

```bash
ManagementEmployee/
├── auth/              # Authentication logic
├── employee/          # Employee-related logic
├── middleware/        # JWT middleware
├── route/             # Route definitions
├── database/          # DB connection
├── models/            # GORM models
├── .env               # Environment variables (e.g., DB config)
├── go.mod / go.sum    # Go module files
├── main.go            # Entry point
```
## How to start
1. Clone the repository
```bash
git clone https://github.com/frank1nz/ManagementEmployee.git
cd ManagementEmployee
```

2. Set up .env
```bash
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=management_employee
JWT_SECRET=your_jwt_secret
```

3. Create MySQL Database
```bash
CREATE DATABASE management_employee CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

4. Run the project
```bash
go run main.go
```
