# GoCat

GoCat is a Go-based project that appears to be a backend service or API. It uses various Go libraries and frameworks to handle database operations, HTTP requests, and configuration management.

## Features

- User management (signup, login, update, delete)
- Menu management
- Category management
- Role management
- Transaction management
- PostgreSQL database integration
- RESTful API endpoints

## Dependencies

The project uses several key dependencies, including:

- github.com/gin-gonic/gin
- github.com/lib/pq
- github.com/spf13/viper
- golang.org/x/crypto

## Project Structure

The project is organized into modules, each handling specific functionality:

- categories
- menu
- transaction0
- transaction1
- payment
- role
- user

Each module typically includes a repository layer for database operations.

## Database

The project uses PostgreSQL as its database. Table names are defined as constants, including:

- categories
- menu
- transaction0
- transaction1
- payment
- users
- role

## API Documentation

### User Endpoints

#### Sign Up
- **POST** `/signup`
- Creates a new user account

#### Login
- **POST** `/login`
- Authenticates a user and returns a token

#### Update User
- **PUT** `/users/{username}`
- Updates user information

#### Delete User
- **DELETE** `/users/{username}`
- Deletes a user account

#### Get All Users
- **GET** `/users`
- Retrieves a list of all users

### Menu Endpoints

#### Create Menu Item
- **POST** `/menu`
- Adds a new item to the menu

#### Get All Menu Items
- **GET** `/menu`
- Retrieves all items in the menu

#### Get Menu Item by ID
- **GET** `/menu/{id}`
- Retrieves a specific menu item

#### Update Menu Item
- **PUT** `/menu/{id}`
- Updates an existing menu item

#### Delete Menu Item
- **DELETE** `/menu/{id}`
- Removes an item from the menu

### Category Endpoints

#### Create Category
- **POST** `/categories`
- Creates a new category

#### Get All Categories
- **GET** `/categories`
- Retrieves all categories

#### Get Category by ID
- **GET** `/categories/{id}`
- Retrieves a specific category

#### Update Category
- **PUT** `/categories/{id}`
- Updates an existing category

#### Delete Category
- **DELETE** `/categories/{id}`
- Removes a category

### Role Endpoints

#### Create Role
- **POST** `/roles`
- Creates a new role

#### Get All Roles
- **GET** `/roles`
- Retrieves all roles

#### Get Role by ID
- **GET** `/roles/{id}`
- Retrieves a specific role

#### Update Role
- **PUT** `/roles/{id}`
- Updates an existing role

#### Delete Role
- **DELETE** `/roles/{id}`
- Removes a role

## Getting Started

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Set up your PostgreSQL database
4. Configure your environment variables
5. Run the application: `go run main.go`

## API Endpoints

(Here you would list the available API endpoints and their functionalities)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

@undrtdprojects