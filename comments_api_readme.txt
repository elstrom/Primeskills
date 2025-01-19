
# Comments API Project

This project is a simple **Comments API** implemented in **Go** using **GORM** and **PostgreSQL**. It allows you to create, retrieve, and delete comments in a PostgreSQL database. The API supports CRUD operations for the `Comment` resource.

## Features:
- **Create a comment**: Allows you to create a new comment with `userId`, `title`, and `body`.
- **Get all comments**: Retrieves all comments stored in the database.
- **Delete a comment**: Allows deletion of a comment by its ID.

## Setup Instructions

### 1. Install Dependencies:
- Install Go from [https://go.dev/dl/](https://go.dev/dl/).
- Install PostgreSQL from [https://www.postgresql.org/download/](https://www.postgresql.org/download/).
- Install GORM:
  ```bash
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/postgres
  ```

### 2. Set Up PostgreSQL:
- Create a database in PostgreSQL:
  ```sql
  CREATE DATABASE comments_db;
  ```
- Ensure that the connection details (user, password, dbname) in the code match your PostgreSQL setup.

### 3. Running the Application:
- Clone or download this repository.
- Run the Go application:
  ```bash
  go run main.go
  ```
- The API will be available at `http://localhost:8080`.

### 4. Testing with Postman:
You can test the API using Postman or any HTTP client.

#### Endpoints:
- **POST** `/comments` to create a comment.
- **GET** `/comments` to get all comments.
- **DELETE** `/comments/{id}` to delete a comment by ID.

### 5. Link to ChatGPT Help:
For further help, refer to this link: [ChatGPT Guide](https://chatgpt.com/share/678d5388-d4f4-8010-9487-f62fb384a7cc)

## License
This project is licensed under the MIT License.

