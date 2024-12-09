# Technical Specification for E-commerce Backend

## Overview

**Purpose:** This document outlines a comprehensive technical specification for developing a robust, scalable, and secure RESTful API backend for an e-commerce platform.

## Functional Requirements

### Product Management

- **Creation:** Add new products with attributes including name, description, price, image, category, and tags.
- **Retrieval:** Fetch all products or specific products by ID, category, or tags.
- **Update:** Modify existing product details (name, description, price, image, category, tags).
- **Deletion:** Remove products from the inventory.

### User Management

- **Registration:** Enable new user registration with attributes such as name, email, password, address, and phone number.
- **Login:** Authenticate users using email and password.
- **Profile Management:** Allow users to update their profile information (address, phone number, preferences).
- **Password Reset:** Provide functionality for users to reset their passwords.

### Order Management

- **Creation:** Create new orders with attributes like user ID, products, shipping address, billing address, and payment method.
- **Retrieval:** Retrieve all orders or specific orders by ID, user ID, or status.
- **Update:** Change order status (e.g., pending, shipped, delivered, canceled).
- **Cancellation:** Allow users to cancel orders.

### Cart Management

- **Addition:** Add products to a user's cart.
- **Removal:** Remove products from a user's cart.
- **Update:** Change the quantity of items in a user's cart.

## Non-Functional Requirements

### Security

- Implement strong authentication and authorization mechanisms using JSON Web Tokens (JWT).
- Use a robust hashing algorithm (e.g., bcrypt) for password encryption.
- Validate user inputs (to prevent SQL injection and cross-site scripting - XSS attacks).

## Technical Specifications

- **Programming Language:** Golang
- **Framework:** Gin
- **Database:** MongoDB
- **ORM:** Mongoose (or a suitable ORM for MongoDB)
- **Authentication Method:** JSON Web Tokens (JWT)
- **API Documentation Tool:** Swagger or OpenAPI
- **Testing Frameworks:** Unit testing, integration testing, end-to-end testing
- **Deployment Strategy:** Containerization (e.g., Docker) on cloud platforms (e.g., AWS, GCP, Azure)

## Database Schema

### Products Collection

| Field         | Type              |
| ------------- | ----------------- |
| `_id`         | Unique Identifier |
| `zid`         | Unique Identifier |
| `name`        | String            |
| `description` | String            |
| `price`       | Number            |
| `image`       | String            |
| `category`    | String            |
| `tags`        | Array of Strings  |

### Admins Collection

| Field       | Type              |
| ----------- | ----------------- |
| `_id`       | Unique Identifier |
| `zid`       | Unique Identifier |
| `name`      | String            |
| `email`     | String            |
| `password`  | Hashed String     |
| `createdAt` | Hashed String     |
| `updatedAt` | Hashed String     |

### Users Collection

| Field          | Type              |
| -------------- | ----------------- |
| `_id`          | Unique Identifier |
| `zid`          | Unique Identifier |
| `name`         | String            |
| `email`        | String            |
| `password`     | Hashed String     |
| `address`      | String            |
| `phone_number` | String            |
| `preferences`  | Object            |

### Orders Collection

| Field              | Type                                 |
| ------------------ | ------------------------------------ |
| `_id`              | Unique Identifier                    |
| `zid`              | Unique Identifier                    |
| `user_id`          | Reference to Users                   |
| `products`         | Array of Product IDs with Quantities |
| `shipping_address` | String                               |
| `billing_address`  | String                               |
| `payment_method`   | String                               |
| `status`           | String                               |

### Carts Collection

| Field     | Type                                 |
| --------- | ------------------------------------ |
| `_id`     | Unique Identifier                    |
| `zid`     | Unique Identifier                    |
| `user_id` | Reference to Users                   |
| `items`   | Array of Product IDs with Quantities |

## API Endpoints

### Product Endpoints

- **POST** `/products`
- **GET** `/products`
- **GET** `/products/:zid`
- **PUT** `/products/:zid`
- **DELETE** `/products/:id`

### User Endpoints

- **POST** `/users`
- **POST** `/login`
- **GET** `/users/:zid`
- **PUT** `/users/:zid`
- **DELETE** `/users/:zid`

### Order Endpoints

- **POST** `/orders`
- **GET** `/orders`
- **GET** `/orders/:zid`
- **PUT** `/orders/:zid`
- **DELETE** `/orders/:zid`

### Cart Endpoints

- **POST** `/carts`
- **GET** `/carts/:userZid`
- **PUT** `/carts/:userZid/items`
- **DELETE** `/carts/:userZid/items/:productId`
