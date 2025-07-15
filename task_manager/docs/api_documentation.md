# 📝 Task Management API Documentation

**Base URL**: `http://localhost:8080/api/tasks`

---

## 📦 Task Object Schema

```json
{
  "id": 1,
  "title": "Complete assignment",
  "description": "Finish the Go task manager API",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "pending"
}
```

| Field       | Type   | Description                              |
|-------------|--------|------------------------------------------|
| `id`        | `int`  | Auto-generated ID of the task           |
| `title`     | `string` | Title of the task                      |
| `description` | `string` | Description of the task              |
| `due_date`  | `string` | Due date in ISO 8601 format            |
| `status`    | `string` | Status (e.g., `pending`, `completed`) |

---

## 📘 Endpoints

### ✅ GET `/api/tasks/`
Retrieve a list of all tasks.

#### 🔸 Example Request:
```bash
GET /api/tasks/
```

#### 🔸 Success Response:
**Code**: `200 OK`

```json
[
  {
    "id": 1,
    "title": "Complete assignment",
    "description": "Finish the Go task manager API",
    "due_date": "2025-08-01T00:00:00Z",
    "status": "pending"
  }
]
```

---

### ✅ GET `/api/tasks/:id`
Retrieve details of a single task by ID.

#### 🔸 Example Request:
```bash
GET /api/tasks/1
```

#### 🔸 Success Response:
**Code**: `200 OK`

```json
{
  "id": 1,
  "title": "Complete assignment",
  "description": "Finish the Go task manager API",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "pending"
}
```

#### 🔸 Error Response:
**Code**: `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

### ✅ POST `/api/tasks/`
Create a new task.

#### 🔸 Example Request:
```bash
POST /api/tasks/
Content-Type: application/json
```

#### 🔸 Request Body:
```json
{
  "title": "New Task",
  "description": "This is a new task",
  "due_date": "2025-08-05T00:00:00Z",
  "status": "pending"
}
```

#### 🔸 Success Response:
**Code**: `201 Created`

```json
{
  "id": 2,
  "title": "New Task",
  "description": "This is a new task",
  "due_date": "2025-08-05T00:00:00Z",
  "status": "pending"
}
```

---

### ✅ PUT `/api/tasks/:id`
Update an existing task.

#### 🔸 Example Request:
```bash
PUT /api/tasks/1
Content-Type: application/json
```

#### 🔸 Request Body:
```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-08-10T00:00:00Z",
  "status": "completed"
}
```

#### 🔸 Success Response:
**Code**: `200 OK`

```json
{
  "id": 1,
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2025-08-10T00:00:00Z",
  "status": "completed"
}
```

#### 🔸 Error Response:
**Code**: `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

### ✅ DELETE `/api/tasks/:id`
Delete a task by ID.

#### 🔸 Example Request:
```bash
DELETE /api/tasks/1
```

#### 🔸 Success Response:
**Code**: `204 No Content`

#### 🔸 Error Response:
**Code**: `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

## 🛑 Error Codes

| Status Code | Description                          | When Returned                          |
|-------------|--------------------------------------|----------------------------------------|
| `200`       | OK                                   | Successful GET or PUT                  |
| `201`       | Created                              | Task created successfully              |
| `204`       | No Content                           | Task deleted successfully              |
| `400`       | Bad Request                          | Invalid input or malformed JSON        |
| `404`       | Not Found                            | Task with the given ID does not exist  |
| `500`       | Internal Server Error                | Server-side error                      |

---

## 🧪 Testing Instructions (Postman or curl)

- **GET** `/api/tasks/` → Returns all tasks.
- **GET** `/api/tasks/1` → Returns task with ID `1`.
- **POST** `/api/tasks/` → Creates a task with a JSON body.
- **PUT** `/api/tasks/1` → Updates task with ID `1`.
- **DELETE** `/api/tasks/1` → Deletes task with ID `1`.