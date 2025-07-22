# 📝 Task Management API Documentation

**Base URL**: `http://localhost:8080/api`

---

## 📦 Task Object Schema

```json
{
  "id": "64b3f13e8f1b2c0012345678",
  "title": "Complete assignment",
  "description": "Finish the Go task manager API",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "pending"
}
```

| Field       | Type     | Description                              |
|-------------|----------|------------------------------------------|
| `id`        | `string` | MongoDB ObjectID of the task            |
| `title`     | `string` | Title of the task                       |
| `description` | `string` | Description of the task               |
| `due_date`  | `string` | Due date in ISO 8601 format             |
| `status`    | `string` | Status (e.g., `pending`, `completed`)   |

---

## 🔒 Authentication and Authorization

### Authentication
- All endpoints (except `/api/users/register` and `/api/users/login`) require a valid **JWT token**.
- The token must be included in the `Authorization` header with the `Bearer` prefix:
  ```
  Authorization: Bearer <token>
  ```

### Authorization
- Some endpoints require specific roles:
  - **Admin Role**: Only users with the `admin` role can access certain endpoints (e.g., promoting users).
  - **User Role**: Regular users can access task-related endpoints.

---

## 📘 Endpoints

### ✅ User Endpoints

#### 🔹 POST `/api/users/register`
Register a new user.

##### Example Request:
```bash
POST /api/users/register
Content-Type: application/json
```

##### Request Body:
```json
{
  "username": "john",
  "password": "password123"
}
```

##### Success Response:
**Code**: `201 Created`

```json
{
  "id": "64b3f13e8f1b2c0012345678",
  "username": "john",
  "role": "user"
}
```

---

#### 🔹 POST `/api/users/login`
Log in and receive a JWT token.

##### Example Request:
```bash
POST /api/users/login
Content-Type: application/json
```

##### Request Body:
```json
{
  "username": "john",
  "password": "password123"
}
```

##### Success Response:
**Code**: `200 OK`

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

#### 🔹 GET `/api/users/:username`
Get a user by username. **Authentication required**.

##### Example Request:
```bash
GET /api/users/john
Authorization: Bearer <user-token>
```

##### Success Response:
**Code**: `200 OK`

```json
{
  "id": "64b3f13e8f1b2c0012345678",
  "username": "john",
  "role": "user"
}
```

##### Error Response:
**Code**: `404 Not Found`

```json
{
  "error": "user not found"
}
```

---

#### 🔹 POST `/api/users/promote/:id`
Promote a user to the `admin` role. **Admin access required**.

##### Example Request:
```bash
POST /api/users/promote/64b3f13e8f1b2c0012345678
Authorization: Bearer <admin-token>
```

##### Success Response:
**Code**: `200 OK`

```json
{
  "id": "64b3f13e8f1b2c0012345678",
  "username": "john",
  "role": "admin"
}
```

##### Error Response:
**Code**: `403 Forbidden`

```json
{
  "error": "Admin Access only"
}
```

---

### ✅ Task Endpoints

#### 🔹 GET `/api/tasks/`
Retrieve a list of all tasks. **Authentication required**.

##### Example Request:
```bash
GET /api/tasks/
Authorization: Bearer <user-token>
```

##### Success Response:
**Code**: `200 OK`

```json
[
  {
    "id": "64b3f13e8f1b2c0012345678",
    "title": "Complete assignment",
    "description": "Finish the Go task manager API",
    "due_date": "2025-08-01T00:00:00Z",
    "status": "pending"
  }
]
```

---

#### 🔹 GET `/api/tasks/:id`
Retrieve details of a single task by ID. **Authentication required**.

##### Example Request:
```bash
GET /api/tasks/64b3f13e8f1b2c0012345678
Authorization: Bearer <user-token>
```

##### Success Response:
**Code**: `200 OK`

```json
{
  "id": "64b3f13e8f1b2c0012345678",
  "title": "Complete assignment",
  "description": "Finish the Go task manager API",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "pending"
}
```

---

#### 🔹 POST `/api/tasks/`
Create a new task. **Authentication required**.

##### Example Request:
```bash
POST /api/tasks/
Authorization: Bearer <user-token>
Content-Type: application/json
```

##### Request Body:
```json
{
  "title": "New Task",
  "description": "This is a new task",
  "due_date": "2025-08-05T00:00:00Z",
  "status": "pending"
}
```

##### Success Response:
**Code**: `201 Created`

```json
{
  "id": "64b3f13e8f1b2c0012345679",
  "title": "New Task",
  "description": "This is a new task",
  "due_date": "2025-08-05T00:00:00Z",
  "status": "pending"
}
```

---

#### 🔹 PUT `/api/tasks/:id`
Update an existing task. **Authentication required**.

##### Example Request:
```bash
PUT /api/tasks/64b3f13e8f1b2c0012345678
Authorization: Bearer <user-token>
Content-Type: application/json
```

##### Request Body:
```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-08-10T00:00:00Z",
  "status": "completed"
}
```

##### Success Response:
**Code**: `200 OK`

```json
{
  "id": "64b3f13e8f1b2c0012345678",
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-08-10T00:00:00Z",
  "status": "completed"
}
```

---

#### 🔹 DELETE `/api/tasks/:id`
Delete a task by ID. **Authentication required**.

##### Example Request:
```bash
DELETE /api/tasks/64b3f13e8f1b2c0012345678
Authorization: Bearer <user-token>
```

##### Success Response:
**Code**: `204 No Content`

---

## 🛑 Error Codes

| Status Code | Description                          | When Returned                          |
|-------------|--------------------------------------|----------------------------------------|
| `200`       | OK                                   | Successful GET or PUT                  |
| `201`       | Created                              | Task or user created successfully      |
| `204`       | No Content                           | Task deleted successfully              |
| `400`       | Bad Request                          | Invalid input or malformed JSON        |
| `401`       | Unauthorized                        | Missing or invalid JWT token           |
| `403`       | Forbidden                           | User lacks the required permissions    |
| `404`       | Not Found                            | Task or user with the given ID does not exist  |
| `500`       | Internal Server Error                | Server-side error                      |

---

## 🧪 Testing Instructions (Postman or curl)

- **GET** `/api/tasks/` → Returns all tasks (requires valid JWT token).
- **POST** `/api/users/login` → Logs in a user and returns a JWT token.
- **GET** `/api/users/:username` → Returns user details by username (requires valid JWT token).
- **POST** `/api/users/promote/:id` → Promotes a user to admin (requires admin token).