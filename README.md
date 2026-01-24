# Gossip With Go
Developer: haoyi

## Overview:
Gossip With Go is a discussion platform where users can create topics, post content, comment, and interact through likes and dislikes. Authentication is required to access most features.

## List of Contents:
- [Gossip With Go](#gossip-with-go)
  - [Overview:](#overview)
  - [List of Contents:](#list-of-contents)
  - [Prerequisites](#prerequisites)
  - [Getting Started](#getting-started)
    - [Quick Setup Steps:](#quick-setup-steps)
      - [1. Frontend Setup (2 commands):](#1-frontend-setup-2-commands)
      - [2. Backend Setup (5 commands):](#2-backend-setup-5-commands)
    - [Detailed Setup Instructions](#detailed-setup-instructions)
      - [1. Clone the Repository](#1-clone-the-repository)
      - [2. Frontend Setup (React + TypeScript)](#2-frontend-setup-react--typescript)
      - [3. Backend Setup (Go)](#3-backend-setup-go)
    - [Available Scripts](#available-scripts)
    - [Troubleshooting](#troubleshooting)
  - [User Guide](#user-guide)
    - [User Access](#user-access)
      - [Login](#login)
      - [Register](#register)
    - [Topics](#topics)
      - [Add Topic](#add-topic)
      - [Update Topic](#update-topic)
      - [Delete Topic](#delete-topic)
      - [Search Topic](#search-topic)
    - [Posts](#posts)
      - [Add Post](#add-post)
      - [Update Post](#update-post)
      - [Delete Post](#delete-post)
      - [Search Post](#search-post)
      - [Like / Dislike Post](#like--dislike-post)
    - [Comments](#comments)
      - [Add Comment](#add-comment)
      - [Update Comment](#update-comment)
      - [Delete Comment](#delete-comment)
      - [Like / Dislike Comment](#like--dislike-comment)
  - [Use of AI](#use-of-ai)

## Prerequisites
- **Node.js** (v16+) and npm
- **Go** (v1.25.5+)
- **Goose** (database migration tool, installed during setup)
- **Docker** and **Docker Compose**
- **PostgreSQL** (via Docker or local installation)

## Getting Started

### Quick Setup Steps:

#### 1. Frontend Setup (2 commands):
- `npm install` –Install dependencies
- `npm run dev` – Start dev server on `http://localhost:5173`

#### 2. Backend Setup (5 commands):
- `cd backend && docker-compose up -d` – Start PostgreSQL
- `go mod download` – Install Go dependencies
- Create a `.env` file with database credentials and a JWT secret key
- `goose up` – Run database migrations
- `go run ./main` – Start the backend on `http://localhost:3000`

---

### Detailed Setup Instructions

#### 1. Clone the Repository
```bash
git clone https://github.com/haobuhaoo/gossip-with-go.git
cd gossip-with-go
```

#### 2. Frontend Setup (React + TypeScript)
In the root directory, run the following commands:
```bash
# Install frontend dependencies
npm install

# Start the development server
npm run dev
```
The frontend will be available at `http://localhost:5173` (or the port assigned by Vite).

#### 3. Backend Setup (Go)

**Step 1: Start PostgreSQL Database**
```bash
# From the root directory, navigate to the backend folder
cd backend

# Start PostgreSQL using Docker Compose
docker-compose up -d
```
This starts a PostgreSQL instance on `localhost:5433` with:
- Username: `postgres`
- Password: `postgres`
- Database: `gossip-with-go`

You can also use your own credentials for username, password and database name.

**Step 2: Install Go Dependencies**
```bash
# From the backend directory
go mod download
```

**Step 3: Set Up Environment Variables**

Create a `.env` file in the `backend` directory and fill in the `GOOSE_DBSTRING` and `JWT_SECRET_KEY` values. Refer to `.env.example` file in the `backend` directory for guidance. Example:
```bash
GOOSE_DBSTRING="<driver>://<username>:<password>@localhost:5433/<database>?sslmode=disable"
JWT_SECRET_KEY="your secret key"
```

**Step 4: Install Goose CLI**
```bash
# Install goose globally (if not already installed)
go install github.com/pressly/goose/v3/cmd/goose@latest
```

**Step 5: Run Database Migrations**
```bash
# From the backend directory
goose up
```
This runs all pending migrations in `backend/internal/postgresql/migrations/`.

**Step 6: Run the Backend Server**
```bash
# From the backend directory
go run ./main
```
The backend will be available at `http://localhost:3000`.

---

### Available Scripts

**Frontend:**
- `npm run dev` – Start development server
- `npm run build` – Build for production

**Backend:**
- `go run ./main` – Start the server

### Troubleshooting

- **Port Already in Use:** Change the frontend port in `vite.config.ts` or the backend port in `.env`
- **Database Connection Failed:** Ensure Docker is running and PostgreSQL is accessible on `localhost:5433`
- **CORS Issues:** Check backend CORS middleware configuration in `backend/internal/api/api.go`
- **Node Modules Issues:** Delete `node_modules` and `package-lock.json`, then run `npm install` again

## User Guide

### User Access
- Users must be logged in to access the protected routes (such as topics, posts and comments).

#### Login

- Enter your username in the input field and click **LOGIN**.
- Upon successful login, you will be directed to the home page where the list of topics will be displayed.

#### Register

- Click **REGISTER** on the login page to register for an account.
- Enter a unique username in the input field that satisfy the following constraints and click **REGISTER**:
  - Must be alphanumeric (`a-z`, `A-Z`, `0-9`) and may contain period (`.`), hyphen (`-`) and underscore (`_`).
  - Must be between 3 to 50 characters long.
- Upon successful login, you will be directed back to the login page.

  **Note:**
  - Usernames are case-sensitive (e.g. `Admin` and `admin` are treated as different users).

---

### Topics
- Topics are displayed in alphabetical order.
- Click a topic to view the list of posts under that topic.

#### Add Topic

- Click the **+ ADD** button at the top right corner of the screen to open the modal form.
- Enter the topic title and click **ADD**.
- The new topic will be appear in the topic list.

  **Note:**
  - The title must be a non-empty string.

#### Update Topic

- Click the **pencil** icon next to a topic to open the update form.
- Enter the new title and click **UPDATE**.
- The updated topic will appear in the topic list.

  **Note:**
  - Only the author of the topic can update it.
  - The title must be a non-empty string.

#### Delete Topic

- Click the **trash** icon next to a topic to delete it.
- The selected topic will be removed from the topic list.

  **Note:**
  - Only the author of the topic can delete it.
  - No confirmation prompt is shown. **Please proceed with caution.**

#### Search Topic

- Enter a query in the search bar and click **SEARCH**.
- Topics with title containing the query will be displayed.

  **Note:**
  - Click the **X** button to reset the topic list.
  - The query must be a non-empty string.

---

### Posts
- Posts are displayed in order of most likes, followed by most recent updated time.
- Posts are truncated in the list view.
- Click a post to view its full content and associated comments.

#### Add Post

- Click the **+ ADD** button at the top right corner of the screen to open the modal form.
- Enter the post title and description and click **ADD**.
- The new post will be appear in the post list.

  **Note:**
  - Both the title and description must be a non-empty string.

#### Update Post

- Click the **pencil** icon next to a post to open the update form.
- Enter the new title and description and click **UPDATE**.
- The updated post will be appear in the post list.

  **Note:**
  - Only the author of the post can update it.
  - Both the title and description must be a non-empty string.

#### Delete Post

- Click the **trash** icon next to the post to delete it.
- The selected post will be removed from the post list.

  **Note:**
  - Only the author of the post can delete it.
  - No confirmation prompt is shown. **Please proceed with caution.**

#### Search Post

- Enter a query in the search bar and click **SEARCH**.
- Posts with title or description containing the query will be displayed.

  **Note:**
  - Click the **X** button to reset the post list.
  - The query must be a non-empty string.

#### Like / Dislike Post

- Click the **thumbs up** icon to like a post.
- Click the **thumbs down** icon to dislike a post.

  **Note:**
  - Click the same button again will remove your reaction.

---

### Comments
- Comments are displayed in order of most likes, followed by most recent updated time.

#### Add Comment

- Enter your comment in the input field below the post and click **ADD**.
- The new comment will be appear in the comment list.

  **Note:**
  - The comment must be a non-empty string.

#### Update Comment

- Click the **pencil** icon next to the comment to update it.
- Update the text and click **UPDATE**.
- The updated comment will be appear in the comment list.

  **Note:**
  - Only the author of the comment can update it.
  - The input description must be a non-empty string.
  - Clicking anywhere outside the comment will cancel update mode.

#### Delete Comment

- Click the **trash** icon next to the comment to delete it.
- The selected comment will be removed from the comment list.

  **Note:**
  - Only the author of the comment can delete it.
  - No confirmation prompt is shown. **Please proceed with caution.**

#### Like / Dislike Comment

- Click the **thumbs up** icon to like a comment.
- Click the **thumbs down** icon to dislike a comment.

  **Note:**
  - Click the same button again will remove your reaction.

## Use of AI

AI was used in this project to:
- Review the `README.md` file for grammer and spelling, and to improve sentence clarity.
- Review and refine application setup instructions for new users.
- Research how to implement JWT-based authentication and authorisization for both frontend and backend.
- Research how to deploy the application on Render.
- Research how to use Docker and Goose for backend containers and database migrations.
