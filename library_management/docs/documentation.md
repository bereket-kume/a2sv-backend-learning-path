# Console-Based Library Management System

## Features
- Add and remove books
- Borrow and return books
- List available books
- List borrowed books by member

## Data Structures Used
- Structs (Book, Member)
- Interfaces (LibraryManager)
- Maps (for storing books and members)
- Slices (for borrowed books)

## Folder Structure

library_management/
├── main.go
├── go.mod
├── controllers/
│ └── library_controller.go
├── models/
│ ├── book.go
│ └── member.go
├── services/
│ └── library_service.go
├── docs/
│ └── documentation.md


## How to Run

```bash
go mod init library_management
go run main.go
