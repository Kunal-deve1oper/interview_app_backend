# Go HTTP Server Setup Guide

This document explains how to set up and run the backend for the \`interview_app_backend\` project.

## Repository

Clone the repository from GitHub:

```bash
git clone https://github.com/Kunal-deve1oper/interview_app_backend.git
```

## Prerequisites

Make sure you have the following installed on your system:
- Go (version 1.20 or higher recommended)
- Git

## Steps to Set Up

1. **Navigate to the Project Directory**
   ```bash
   cd interview_app_backend
   ```

2. **Install Dependencies**
   Run the following command to download and install the project's dependencies:
   ```bash
   go mod tidy
   ```

3. **Set Up Environment Variables**
   Create a `.env` file in the root directory of the project and add the following variables:
   ```
   DB_CONNECTION_STRING=<your-database-connection-string>
   MAIL_SERVICE_EMAIL=<your-mail-service-email>
   MAIL_SERVICE_PASSWORD=<your-mail-service-password>
   ```

   Replace `<your-database-connection-string>`, `<your-mail-service-email>`, and `<your-mail-service-password>` with your actual credentials.

   Example:
   ```
   DB_CONNECTION_STRING=postgres://username:password@localhost:5432/interview_db
   MAIL_SERVICE_EMAIL=example@gmail.com
   MAIL_SERVICE_PASSWORD=yourpassword123
   ```

4. **Run the Server**
   Start the server using the following command:
   ```bash
   go run main.go
   ```

   The server will typically run on `http://localhost:8080`. You can configure the port in the code if needed.

