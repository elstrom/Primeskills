# Comments API - Golang with GORM and PostgreSQL

This project is a simple RESTful API built using **Golang**, **GORM**, and **PostgreSQL**. The API allows you to perform basic operations (POST, GET, DELETE) on comments data. This API supports storing comments with fields such as `userId`, `id`, `title`, and `body`.

## Features
- Create a new comment
- Retrieve all comments
- Delete a comment by ID

## Prerequisites
To get started with this project, you'll need the following installed:
- Go (Golang) v1.18 or higher
- PostgreSQL database
- GORM ORM for Golang

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/comments-api.git
    cd comments-api
    ```

2. **Set up your PostgreSQL database**:
    - Ensure you have a PostgreSQL server running on `localhost` (or configure the connection settings in the `database/database.go` file).
    - Create a database called `comments_db`.

    ```sql
    CREATE DATABASE comments_db;
    ```

3. **Install Go dependencies**:
    Make sure you have Go installed on your system and that the GOPATH is properly set up.
    Then, install all the required Go dependencies using:
    ```bash
    go mod tidy
    ```

4. **Run the application**:
    To start the application, run the following command:
    ```bash
    go run main.go
    ```
    The server will be available at `http://localhost:8080`.

## API Endpoints

1. **POST /comments**  
    Create a new comment.
    - **Request body**:
    ```json
    {
        "userId": 1,
        "title": "Example Comment Title",
        "body": "This is an example comment body."
    }
    ```
    - **Response**:
    ```json
    {
        "id": 1,
        "userId": 1,
        "title": "Example Comment Title",
        "body": "This is an example comment body."
    }
    ```

2. **GET /comments**  
    Get all comments.
    - **Response**:
    ```json
    [
        {
            "id": 1,
            "userId": 1,
            "title": "Example Comment Title",
            "body": "This is an example comment body."
        }
    ]
    ```

3. **DELETE /comments/{id}**  
    Delete a comment by its ID.
    - **URL**: `http://localhost:8080/comments/{id}`
    - **Response**: 204 No Content

## Running Tests

This project does not currently include unit tests. However, you can easily test the API using tools like [Postman](https://www.postman.com/).

### ChatGPT Guide

You can find the complete guide for building the Comments API with Golang, GORM, and PostgreSQL based on ChatGPT's instructions here: [ChatGPT Guide](https://chatgpt.com/share/678d5388-d4f4-8010-9487-f62fb384a7cc).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Author**: Ramdan Kurnia  
**Email**: Danram162@gmail.com
