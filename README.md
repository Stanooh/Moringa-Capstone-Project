## Prompt-Powered Kickstart: Beginner’s Toolkit for Go (Golang)

## Overview

This project is a beginner-friendly RESTful API built using Go (Golang).

The goal of this project was to explore backend development using a compiled, strongly typed language and to understand how scalable APIs are structured using Go’s standard library.

Go is widely used in cloud computing, DevOps tooling, and distributed systems, making it a valuable skill for modern backend development.

# Project Objective

Build a functional RESTful API that:

Handles HTTP requests

Returns structured JSON responses

Supports GET and POST operations

Demonstrates routing and request validation

Uses only Go’s standard library

Can be easily understood and replicated by beginners

# About Go

Go is an open-source, statically typed, compiled programming language developed at Google. It is designed for:

Simplicity

High performance

Scalability

Built-in concurrency support

## Common Use Cases

Backend services

Microservices

Cloud-native applications

DevOps tooling

## Notable Projects Built with Go

Docker

Kubernetes

## System Requirements

Windows / macOS / Linux

Go 1.20+

VS Code or any Go-supported IDE

Terminal (PowerShell / Bash / Command Prompt)

# Installation & Setup

# Install Go

Download and install Go from:
https://go.dev/dl/

Verify installation:

go version

## Initialize the Project

mkdir go-beginner-toolkit
cd go-beginner-toolkit
go mod init beginner-toolkit

## Create main.go

Create the main application file:

touch main.go

Add the project source code.

## Run the Application

go run main.go

Server runs at:

http://localhost:8080

## API Endpoints

Method Endpoint Description
GET /items Retrieve all items
GET /items/{id} Retrieve item by ID
POST /items Create a new item
📦 Example Request & Response
POST /items

Request Body

{
"name": "Book",
"price": 500
}

Response

{
"id": 1,
"name": "Book",
"price": 500
}

## Project Structure

go-beginner-toolkit/
│── main.go
│── README.md

## AI-Assisted Development

This project was developed with guided assistance from AI to enhance understanding and implementation.

Key Prompts Used

“Give me a beginner-friendly step-by-step guide to building a REST API in Go using only the standard library.”
→ Helped structure routing and JSON handling using net/http.

“How do I parse URL parameters and convert them safely in Go?”
→ Implemented dynamic routes using strconv.Atoi with proper error handling.

“How should a beginner structure a simple in-memory database in Go?”
→ Implemented an in-memory slice of structs with manual ID management.

“Generate a professional README.md file for a beginner Go REST API project.”
→ Improved documentation clarity and structure.

Reflection

AI was used as a learning assistant to:

Strengthen understanding of Go’s HTTP package

Improve backend architecture decisions

Enhance debugging skills

Improve documentation quality

The focus remained on understanding the concepts rather than copying solutions.

## Common Issues & Troubleshooting

'go' is not recognized

Ensure Go is installed correctly

Confirm Go is added to your system PATH

Port already in use

Change the port in main.go

Stop the process currently using port 8080

Invalid JSON errors

Ensure request body is valid JSON

Verify required fields are included

## Learning Resources

Official Documentation → https://go.dev/doc/

Go by Example → https://gobyexample.com/

Tour of Go → https://go.dev/tour/

👨‍💻 Author

Stanley Kimani
Aspiring Backend Developer | Building Scalable Systems with Go
