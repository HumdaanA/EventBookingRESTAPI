# Event Booking REST API

## Overview

This is a personal project designed to build a robust Event Booking REST API using GoLang. The application supports various event-related functionalities, including user registration, event management, and secure authentication. It incorporates several advanced features to ensure scalability, security, and maintainability.

## Features
### Event Management:

- **GET**: Retrieve a list of available events
- **GET**: Fetch details of a single event
- **POST**: Create a new bookable event (authentication required)
- **PUT**: Update an existing event (authentication required)
- **DELETE**: Remove an event (authentication and ownership required)

### User Management:

- **POST**: Create a new user
- **POST**: User login with token-based authentication

### Registration Management:

- **POST**: Register a user for an event (authentication required)
- **DELETE**: Cancel a user's registration for an event (authentication required)

### Authentication:

- Secured endpoints using JWT-based authentication
- Token verification required for sensitive routes
- User-specific permissions for event modifications
  
## Technologies Used

- **GoLang**: Primary programming language for building the REST API
- **Gin Framework**: Used for routing and handling HTTP requests
- **Golang-JWT**: Provides token-based authentication
- **Crypto/Bcrypt**: Secures user passwords by hashing them
- **Database/SQL & Go-SQLite3**: Manages persistent storage and handles SQL operations

## Project Structure

- **Routes**: Defined in routes.go within the routes package using the Gin framework
- **Database Initialization**: Managed in db.go to handle table creation and data persistence
- **Authentication**: Implemented using middleware for secure API endpoints
- **Password Hashing**: Utility functions in hash.go for password hashing and verification

## Security Highlights

- Token-based authentication for sensitive routes
- Password hashing to prevent plain-text password storage
- Ownership validation for event modifications














  
