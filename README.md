# Go Stack Package

A simple, flexible Go package that implements multiple types of stack data structures.

This package offers a variety of stack implementations (e.g., array-backed stack, linked list-backed stack), allowing you to choose the one that best fits your needs. It supports basic stack operations such as push, pop, peek, and checking if the stack is empty.
<!-- TOC -->
* [Go Stack Package](#go-stack-package)
    * [Features](#features)
  * [Installation](#installation)
  * [Usage](#usage)
    * [Example Code](#example-code)
    * [Supported Operations](#supported-operations)
    * [Stack Implementations](#stack-implementations)
  * [Contribute](#contribute)
  * [License](#license)
<!-- TOC -->


### Features

- Array-backed stack implementation

- Linked list-backed stack implementation

- Clean, modular code for ease of use and extendability

- Simple and efficient API for stack operations

- Supports a variety of use cases in algorithms, system design, and more

## Installation

To install this package in your Go project, run:

```bash
go get github.com/amarsinghrathour/go-stack/stack
```

## Usage

### Example Code

Here is an example demonstrating how to use the stack package:

```go
package main

import (
"fmt"
"github.com/amarsinghrathour/go-stack/stack"
)

func main() {
	// Create a new array stack
	arrStack := stack.NewArrayStack()

	// Push elements onto the stack
	arrStack.Push(10)
	arrStack.Push(20)
	arrStack.Push(30)

	// Peek the top element
	top, _ := arrStack.Peek()
	fmt.Println("Top element:", top)  // Output: Top element: 30

	// Pop an element from the stack
	popValue, _ := arrStack.Pop()
	fmt.Println("Popped element:", popValue)  // Output: Popped element: 30

	// Check if stack is empty
	isEmpty := arrStack.IsEmpty()
	fmt.Println("Is stack empty?", isEmpty)  // Output: Is stack empty? false
}
```

### Supported Operations

All stack types support the following operations:

- Push(value interface{}) - Adds an element to the top of the stack.

- Pop() (returns interface{}, error) - Removes and returns the top element of the stack.

- Peek() (returns interface{}, error) - Returns the top element without removing it.

- IsEmpty() (returns bool) - Checks if the stack is empty.

### Stack Implementations

1. Array Stack (ArrayStack)
   - A simple array-backed stack implementation.
   - Suitable for use cases with a known maximum size or where resizing is not a concern.

2. Linked List Stack (LinkedListStack)

   - A more flexible linked-list-backed stack implementation.
   - Ideal when you need dynamic resizing or have unpredictable stack sizes.

## Contribute

If you'd like to contribute to this project, feel free to fork the repository and submit pull requests. Here are some guidelines to keep in mind:

- Add tests for new features or bug fixes.

- Follow Go's conventions for code formatting.

- Ensure that all tests pass before submitting your pull request.

## License

This project is licensed under the MIT [License]() - see the LICENSE file for details.