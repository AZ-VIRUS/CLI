# Anime Booking CLI 

A simple command-line ticket booking system written in Go that allows users to book tickets for popular anime shows. The project demonstrates core Go concepts such as structs, slices, functions, input validation, packages, and goroutines.

---

## Features

* Select from multiple anime shows
* Book tickets using the command line
* Input validation for name, email, and ticket count
* View all bookings
* Display a list of users who booked tickets
* Simulated ticket sending using goroutines
* Asynchronous ticket confirmation

---

## Project Structure

```
anime-booking/
│
├── main.go
├── helper/
│   └── helper.go
└── README.md
```

### main.go

Contains the main application logic including ticket booking, anime selection, booking display, and ticket sending.

### helper/helper.go

Handles input validation such as validating names, email format, and ticket availability.

---

## Available Anime Shows

1. Naruto
2. Attack on Titan
3. One Piece
4. Demon Slayer
5. Jujutsu Kaisen

---

## Requirements

* Go 1.18 or newer

Check your Go version:

```
go version
```

---

## Running the Program

Clone the repository:

```
git clone https://github.com/AZ-VIRUS/anime-booking.git
```

Navigate to the project directory:

```
cd anime-booking
```

Run the application:

```
go run main.go
```

---

## Concepts Demonstrated

* Go structs
* Slices
* Loops
* Functions and return values
* Packages
* Input validation
* Goroutines and concurrency
* Building a command-line application

---



