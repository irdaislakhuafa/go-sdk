# Go Software Development Kit

This is my code collection for development with Go Programming Language.
Here are some packages that i often use for developing Golang projects.

- [appcontext](./appcontext/): Provides utilities for managing request-scoped data within a Go application's context, simplifying data sharing and retrieval across different parts of the request lifecycle.
- [auth](./auth/)
    - [jwt](./auth/jwt.go): Effortless JWT Generation, Validation, and Claims Extraction
- [codes](./codes/): A package for defining and handling custom error codes in Go applications, enabling streamlined error management and identification.
- [collections](./collections/)
    - [slice](./collections/slice.go): Provides utility functions for manipulating and comparing slices, simplifying common slice operations.
- [concurrency](./concurrency/)
    - [concurrency](./concurrency/concurrency.go): A package for running asynchronous functions in Go with controlled goroutine limits, preventing memory leaks caused by uncontrolled goroutine execution. It simplifies concurrent task management and ensures efficient resource utilization.
- [convert](./convert/)
    - [ptr](./convert/ptr.go): Provides utilities for working with pointers, including type-safe value extraction without nil checks and direct conversion of function return values to pointers. This simplifies pointer handling and reduces boilerplate code.
- [cryptography](./cryptography/)
    - [aes256gcm](./cryptography/aes256gcm.go): Provides functions for encrypting and decrypting data using the AES-256-GCM algorithm. This package simplifies the usage of AES-256-GCM in Go applications for secure data protection.
    - [argon2](./cryptography/argon2.go): Argon2 implementation in Go with a standard format compatible with [node-argon2](https://github.com/ranisalt/node-argon2), ensuring interoperability between Go and Node.js applications.
    - [bcrypt](./cryptography/bcrypt.go): Provides a simplified interface for bcrypt hashing using a builder pattern, making it easier to generate and verify bcrypt hashes securely.
    - [sha256](./cryptography/sha256.go): Simplifies generating SHA256 hashes with or without a key, providing a convenient interface for incorporating SHA256 hashing into your Go applications.
- [datastructure](./datastructure/)
    - [set](./datastructure/set.go): Provides a hash set implementation in Go, utilizing a map for efficient storage and retrieval of unique elements. This package simplifies working with sets of data and ensures element uniqueness.
    - [queue](./datastructure/queue.go): Provides a basic FIFO (First-In, First-Out) queue implementation in Go, enabling the management of ordered data with enqueue and dequeue operations.
- [errors](./errors/): Provides utilities for generating and handling errors with custom error codes defined in the `codes` package. This enables efficient error identification and allows for implementing different logic based on specific error types.
- [files](./files/): Provides utilities for managing local files, such as checking file existence, getting file extensions, and performing other file-related operations.
- [header](./header/): Offers functions for manipulating and extracting information from HTTP headers, streamlining common header-related tasks in Go applications.
- [language](./language/): Provides support for managing multi-language messages within your Go application, enabling easy localization and internationalization.
- [log](./log/): A simplified logging wrapper around the [zerolog](https://github.com/rs/zerolog.git) package. It provides convenient functions for retrieving values from the context, adding custom fields, and logging messages with various levels.
- [operator](./operator/): Provides implementations of common logical operators, including a ternary operator, to enhance code readability and conciseness in Go applications.
- [smtp](./smtp/): Provides a simplified interface for sending emails using SMTP, acting as a wrapper around the [gomail](https://github.com/go-mail/gomail.git) library. It simplifies email sending tasks and offers convenient configuration options.
- [sorter](./sorter/): Provides a generic sorting function that works with any type implementing the `sort.Interface` interface, making it easy to sort slices of custom data structures.
- [strformat](./strformat/): Provides a string formatting utility similar to JavaScript's template literals, but for Go. It leverages the `text/template` package from the Go standard library to enable dynamic string generation with placeholders and expressions.

## Documentations

You can read full documentations [here](https://pkg.go.dev/github.com/irdaislakhuafa/go-sdk)

## How to install

```bash
go get github.com/irdaislakhuafa/go-sdk@latest
```

## TODO
- [ ] Implement storage file operation for type disk
