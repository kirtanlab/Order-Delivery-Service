<!-- tesssst -->
# Order Management Service (order-ms)

A microservice-based order management system built with Go and Gin framework that handles orders, brands, products, customers, and rider management. This service provides a complete backend for e-commerce or food delivery platforms.

## Features


- **Admin Management**: User authentication with role-based access control
- **Brand Management**: Create and manage brands, categories, and products
- **Order Processing**: Complete order lifecycle from creation to delivery
- **Customer Management**: Customer profiles with authentication
- **Rider Assignment**: Rider allocation and tracking
- **Hub System**: Geographic organization of service areas
- **RESTful API**: Complete API for all operations

## Tech Stack

- **Language**: Go 1.18+
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Development Tools**: Air (for live reload)

## Project Structure

```
order-ms/
├── .air.toml           # Air configuration for live reload
├── .env.example        # Environment variables template
├── main.go             # Application entry point
├── pkg/                # Package directory containing all modules
│   ├── admin/          # Admin user management
│   ├── brand/          # Brand, product, and category management
│   ├── hub/            # Geographic service areas
│   ├── market/         # Customer and store-front features
│   ├── middleware/     # Authentication middlewares
│   ├── order/          # Order processing and management
│   ├── rider/          # Delivery personnel management
│   ├── utils/          # Utility functions
│   └── validator/      # Input validation
└── v4.sql              # Database schema
```

## Setup Instructions

### Prerequisites

- Go 1.18+
- PostgreSQL
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/shariarfaisal/order-ms.git
   cd order-ms
   ```

2. Copy the environment file and configure it:
   ```bash
   cp .env.example .env
   # Edit the .env file with your database credentials and other settings
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Setup the database:
   ```bash
   # Create a PostgreSQL database
   # Import the schema from v4.sql or let GORM handle migrations
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

### Development with Live Reload

For development with automatic reloading:

1. Install Air:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

2. Run the application with Air:
   ```bash
   air
   ```

## API Endpoints

### Admin Routes
- `POST /admin/create` - Create admin user
- `POST /admin/login` - Admin login
- `GET /admin/me` - Get admin profile

### Brand Routes
- `POST /brand/create` - Create brand
- `GET /brand/` - List all brands
- `POST /brand/category/create` - Create brand category
- `GET /brand/category/` - List brand categories

### Product Routes
- `POST /products/create` - Create product
- `GET /products/` - List all products
- `DELETE /products/:id` - Delete product

### Order Routes
- `POST /orders/create` - Create order

### Customer Routes
- `POST /customer/signup` - Customer signup
- `POST /customer/login` - Customer login
- `GET /customer/me` - Get customer profile

### Hub Routes
- `POST /hubs/create` - Create hub
- `GET /hubs/` - List all hubs
- `GET /hubs/:id` - Get hub by ID

## Environment Variables

```
ENV=development
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=orderms
DB_PORT=5432
APP_SECRET=secret
```

## Database Schema

The database schema includes tables for:
- Brands and Products
- Orders and Order Items
- Customers and Delivery Addresses
- Riders and Assignments
- Payment Logs
- Hubs and Geographical organization

## Development Notes

- Use the `air` command for live reloading during development
- Add Go modules with `go get [package-name]`
- Run `go mod tidy` to clean up dependencies
<!-- comments hello-->

## License

[MIT License](LICENSE)
