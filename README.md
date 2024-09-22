

# Project Overview
This document describes a microservice for managing members and assigning recipients in a Secret Gift Exchange. 
This microservice focuses on the data management aspects of a Secret Gift Exchange. It provides functionalities to manage family members and their assigned gift recipients. The service adheres to the following principles:

- **Simplicity**: As it was designed to be completed within a few hours, prioritizing clean and maintainable code.
- **Conciseness and Readability**: The code is written in a concise and readable style, following Go's best practices for formatting and naming conventions.
- **Modularity**: The code is organized into well-defined modules and functions, promoting reusability and maintainability.
- **web framework for Go**: Gin, A popular web framework for Go, providing a concise and expressive API for building web applications.
- **Persistent Database**: SQLite3, A lightweight, serverless, embedded SQL database engine.
- **Object-Relational Mapper**:GORM, An Object-Relational Mapper (ORM) library for Go that simplifies database interactions.
- **Minimal API**: Offers basic CRUD operations for family members and retrieval of assigned recipients. 
- **Unit Testing**: Includes unit tests to validate service functionality.

## API Reference
<p>The service exposes a REST API for managing family members and gift exchange assignments. A Postman collection is provided to facilitate manual testing of the API.</p>

<p>A Postman collection containing pre-configured requests for testing the API is available. This collection can be imported into Postman to easily interact with the service and verify its functionality.</p>

<p>Please find the collection in <em><strong> GiftSwap -> test -> postman -> GiftSwap Test collection.postman_collection.json</strong></em> </p>



# Data Model

> #### Family Member:
>
> - **id**: Unique identifier (uint)
> - **name**: Name of the family member (string)

> #### Gift Exchange:
> 
> - **AssignerID**:  ID of the member the family member will be gifting (uint)
> - **RecipientID**:  ID of the family member participating in the exchange (uint)
> - **Year**: Year when they exchanged the gift.


# REST Endpoints
- **GET /members**: Lists all family members.
- **GET /members/{id}**: Retrieves a single family member by ID.
- **POST /members**: Creates a new family member.
- **PUT /members/{id}**: Updates an existing family member.
- **DELETE /members/{id}**: Deletes a family member.
- **GET /gift_exchange**: Lists all family members with their assigned recipient IDs (requires implementation of gift exchange logic).

# Current Implementation
- The current version focuses on managing family members.
- Current version has some issues with correct implementation the gift exchange logic (assigning recipients).
- Data is stored in database sqlite.
- Basic Auth is applied to **DELETE /members/{id}** API.
- Few Unit tests are provided to validate the functionality of the service for CRUD operations on family members.
- Postman collection is provided to test application manually.


# Future Improvements
- Implement the correct gift exchange logic to randomly assign recipients while adhering to constraints (no self-gifting, avoiding repeated pairs within 3 years time frame).
- Explore concurrency handling mechanisms if the service needs to handle simultaneous requests.
- Introduce a robust authentication (e.g., JWT, OAuth) and authorization / RBAC mechanism to protect sensitive endpoints.
- Need to implement strict input validation to ensure all user-provided data is validated against expected formats and ranges.
- Implement sensitive data encryption to encrypt sensitive data at rest (e.g., stored in databases like password) and in transit (e.g., network communication).

# Getting Started
- Clone the Repository.
  ``git clone <repository_url>``
- Install dependencies.
  ``go mod tidy``
- Build and run the service.
  ``go run main.go``

# Disclaimer
This project serves as a foundation for a Secret Gift Exchange microservice. It is intended to be a starting point and can be extended and improved upon.