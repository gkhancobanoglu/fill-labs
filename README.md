# 🌟 Word Sorting by 'a' Count and Length

This Go program sorts a list of words based on:
1. The number of occurrences of the letter "a" in each word (in decreasing order).
2. If the "a" count is the same, it sorts the words by length (in decreasing order).

## How to Use

1. Clone the repository.
2. Run the program: `go run q1_main.go`
3. Enter words to add to the list or type `exit` to quit.

## Additional Features

- Users can add new words to the list during runtime.
- The program exits when the user types `exit`.


# 🔄 Recursive Function to Reverse Sequence

This Go program defines a recursive function that takes an integer, divides it by 2 recursively, and prints the sequence in reverse order.

## How to Use

1. Clone the repository.
2. Run the program: `go run q2_main.go`
3. The function will print the reversed sequence of numbers generated from the input.

## Example

For input `9`, the output will be:

2 4 9


---

# 🔍 Find Most Repeated Data in an Array

This Go program defines a function to find the most repeated element in a given array of strings.

## Problem Description

The program accepts an array or list of strings and finds the most frequently occurring element. If there are multiple elements with the same highest frequency, it returns the first one.

## How to Use

1. Clone the repository.
2. Run the program: `go run q3_main.go`
3. The program will output the most repeated element in the array.

## Additional Features

- Users can add new strings to the array during runtime.
- The program exits when the user types `exit`.


# 👥 User Management System

This project consists of a user management system with both a backend and a frontend. It allows users to perform CRUD operations, including creating, editing, and deleting users. The system is built with a Go backend and a React/Next.js frontend.

## Features

- **Master View**: Lists all users in a data grid, with buttons for New, Edit, and Delete operations.
- **Detailed View**: Displays a form with two buttons (**Action**, **Back**). The action button changes text based on the selected operation:
    - **New** → "Create"
    - **Edit** → "Save"
    - **Delete** → "Delete"
- **Backend**: A RESTful API that supports operations for listing users, retrieving a user by ID, creating, updating, and deleting users. The API follows REST standards for paths, HTTP methods, and status codes.
- **Persistent Storage**: SQLite database for data persistence, with the database file included in the project folder.

---

## 🔧 Backend Setup

1. Navigate to the **backend** directory:

    ```bash
    cd user-management-system/backend/api
    ```

2. Run the backend server:

    ```bash
    go run main.go
    ```

---

## 💻 Frontend Setup

1. Navigate to the **frontend** directory:

    ```bash
    cd user-management-system/frontend
    ```

2. Install the required dependencies:

    ```bash
    npm install
    ```

3. Run the frontend application:

    ```bash
    npm run dev
    ```

The frontend will now be running and accessible at [http://localhost:3000](http://localhost:3000).

---

## 🌐 API Endpoints

The following REST API endpoints are available:

- **GET** `/users`: Returns a list of all users.
- **GET** `/users/{id}`: Returns the user with the specified ID.
- **POST** `/users`: Saves a new user.
- **PUT** `/users/{id}`: Updates the user with the specified ID.
- **DELETE** `/users/{id}`: Deletes the user with the specified ID.

---

## 🛠️ Technologies Used

- **Backend**: Go
- **Frontend**: TypeScript, React, Next.js
- **Database**: SQLite
