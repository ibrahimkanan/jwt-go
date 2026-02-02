# JWT Authentication with Go & React

A robust full-stack web application demonstrating secure JWT (JSON Web Token) authentication. This project features a high-performance Go backend using the Echo framework and a modern React frontend with Vite.

## 🚀 Tech Stack

### Backend

- **Language:** Go (Golang)
- **Framework:** [Echo v4](https://echo.labstack.com/)
- **Database:** PostgreSQL
- **ORM:** [GORM](https://gorm.io/)
- **Authentication:** JWT (JSON Web Token) & Bcrypt
- **Environment Management:** Godotenv

### Frontend

- **Framework:** [React](https://react.dev/)
- **Build Tool:** [Vite](https://vitejs.dev/)
- **State Management:** [Zustand](https://github.com/pmndrs/zustand)
- **Routing:** React Router DOM
- **HTTP Client:** Axios
- **Notifications:** React Hot Toast
- **Styling:** CSS3

## 🛠️ Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/dl/) (v1.20+)
- [Node.js](https://nodejs.org/) (v18+)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/) (Optional)

## 📦 Installation & Setup

### Option 1: Docker (Recommended)

1. **Clone the repository**

   ```bash
   git clone https://github.com/ibrahimkanan/jwt-go.git
   cd jwt-go
   ```

2. **Run with Docker Compose**
   ```bash
   docker-compose up --build
   ```
   The backend will start on `http://localhost:3000` and the database on port `5435`.

### Option 2: Local Development

#### Backend Setup

1. Navigate to the server directory:

   ```bash
   cd Server
   ```

2. Create a `.env` file in the `Server` directory based on the variables required (see `ConnectToDB.go` and `usersController.go`):

   ```env
   DB_HOST=localhost
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=your_db_name
   DB_PORT=5432
   SECRET=your_jwt_secret_key
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the server:
   ```bash
   go run main.go
   ```

#### Frontend Setup

1. Navigate to the client directory:

   ```bash
   cd ../Client
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

## 🔌 API Endpoints

| Method | Endpoint    | Description                      | Auth Required |
| ------ | ----------- | -------------------------------- | ------------- |
| POST   | `/signup`   | Register a new user              | No            |
| POST   | `/login`    | Login and receive JWT            | No            |
| POST   | `/logout`   | Logout and clear cookie          | No            |
| GET    | `/validate` | Validate session / Get user info | Yes           |

## 📂 Project Structure

```
jwt-go/
├── Client/                 # React Frontend & Vite Configuration
│   ├── src/
│   │   ├── pages/         # Login & Signup Pages
│   │   ├── stores/        # Zustand State Store
│   │   └── ...
├── Server/                 # Go Backend
│   ├── controllers/       # Route Handlers (Auth logic)
│   ├── initializers/      # DB Connection & Env Loading
│   ├── middleware/        # JWT Validation Middleware
│   ├── models/            # GORM Models
│   └── main.go            # Entry point & Routes
└── docker-compose.yml     # Container Orchestration
```

## 📝 License

This project is open source and available under the [MIT License](LICENSE).
