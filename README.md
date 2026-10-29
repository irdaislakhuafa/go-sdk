# Go Software Development Kit

A comprehensive collection of production-ready utility packages and helpers designed to simplify Go application development, reduce boilerplate, and enforce robust patterns.

## Table of Contents

- [Packages](#packages)
  - [appcontext](#appcontext)
  - [auth](#auth)
  - [caches](#caches)
  - [codes](#codes)
  - [collections](#collections)
  - [concurrency](#concurrency)
  - [convert](#convert)
  - [cryptography](#cryptography)
  - [datastructure](#datastructure)
  - [db](#db)
  - [errors](#errors)
  - [files](#files)
  - [header](#header)
  - [http](#http)
  - [language](#language)
  - [log](#log)
  - [operator](#operator)
  - [querybuilder/sqlc](#querybuildersqlc)
  - [random](#random)
  - [smtp](#smtp)
  - [sorter](#sorter)
  - [storage](#storage)
  - [strformat](#strformat)
- [Documentation](#documentation)
- [Installation](#installation)

## Packages

### [appcontext](./appcontext/)
Provides utilities for managing request-scoped data within context objects, simplifying sharing and retrieval of lifecycle data across application layers.

### [auth](./auth/)
Authentication and token management tools:
- **jwt**: Simplifies JSON Web Token generation, validation, and claims extraction using `golang-jwt/jwt/v5`.

### [caches](./caches/)
Caching abstractions supporting multiple backend stores:
- **memory**: In-memory thread-safe cache implementation.
- **sqlite**: SQLite-backed persistent cache storage.

### [codes](./codes/)
Standardized custom error codes and mapping functions to streamline error classification, identification, and response formatting across services.

### [collections](./collections/)
Generic collection manipulation utilities:
- **slice**: Operations for filtering, mapping, checking, and comparing slices.
- **map**: Map manipulation and key-value utilities.
- **reduce**: Functional reduction helpers for slices and iterables.

### [concurrency](./concurrency/)
Utilities for safe asynchronous execution:
- **concurrency**: Runs asynchronous functions with controlled goroutine limits to prevent memory leaks and resource exhaustion.
- **semaphore**: Lightweight concurrency limiting primitives.

### [convert](./convert/)
Type conversion and pointer helpers:
- **ptr**: Utilities for working with pointers, including type-safe value extraction without nil checks and direct conversion of primitive values to pointers.

### [cryptography](./cryptography/)
Secure hashing and encryption primitives:
- **aes256gcm**: Encrypts and decrypts data using the AES-256-GCM authenticated encryption standard.
- **argon2**: Argon2 password hashing compatible with standard Node.js implementations for cross-language interoperability.
- **bcrypt**: Fluent builder interface for secure password hashing and verification.
- **sha256**: Convenient hashing utility supporting keyed and unkeyed SHA-256 generation.

### [datastructure](./datastructure/)
Common data structures implemented in generic Go:
- **set**: Hash set implementation utilizing maps for efficient storage and retrieval of unique elements.
- **queue**: First-In-First-Out queue management with standard enqueue and dequeue operations.

### [db](./db/)
Database connection initializers and connection managers:
- **mysql**: MySQL connection setup using `go-sql-driver/mysql`.
- **postgresql**: PostgreSQL connection setup using `lib/pq`.
- **sqlite**: SQLite connection setup using `go-sqlite3`.

### [errors](./errors/)
Advanced error handling wrapper supporting custom error codes, contextual metadata, and stack trace capture for deep debugging.

### [files](./files/)
Local file management utilities including existence verification, extension extraction, and path helper functions.

### [header](./header/)
HTTP header manipulation utilities for extracting, setting, and parsing incoming and outgoing request headers.

### [http](./http/)
HTTP debugging utilities for inspecting requests, responses, and payload data during development.

### [language](./language/)
Multi-language support and internationalization messages for localizing application output.

### [log](./log/)
A structured logging wrapper around `zerolog` featuring context integration, custom field injection, and multi-destination output including file rotation (`lumberjack`) and console formatting.

### [operator](./operator/)
Logical operators including a concise ternary operator function to improve inline conditional readability.

### [querybuilder/sqlc](./querybuilder/sqlc/)
Helpers and type-safe wrappers for SQLC generated code to simplify database operations and transaction management.

### [random](./random/)
Random integer and token generation utilities.

### [smtp](./smtp/)
Simplified email sending wrapper around `gomail` supporting HTML bodies, attachments, and SMTP configuration options.

### [sorter](./sorter/)
Generic sorting functions compatible with custom data structures and custom comparators.

### [storage](./storage/)
Cloud and local object storage abstraction layer supporting:
- **minio**: S3 compatible object storage integration via MinIO SDK.
- **disk**: Local disk file storage operations.
- **dummy**: Mock storage backend for unit testing.

### [strformat](./strformat/)
String formatting utility inspired by template literals, leveraging Go standard `text/template` for dynamic interpolation with placeholders.

## Documentation

Full package documentation and API references are available on [GoPkgDev](https://pkg.go.dev/github.com/irdaislakhuafa/go-sdk).

## Installation

```bash
go get github.com/irdaislakhuafa/go-sdk@latest
```
