# Pokedex REPL

This is a command-line Pokedex REPL (Read-Eval-Print Loop) written in Go. It allows you to explore the world of Pokemon, catch them, and inspect the ones you've caught.

## Features

* **Explore Pokemon world:** Discover new areas and the Pokemon that inhabit them.
* **Catch Pokemon:** Try your luck at catching wild Pokemon.
* **Pokedex:** Keep track of all the Pokemon you've caught.
* **Inspect Pokemon:** View detailed information about the Pokemon in your Pokedex.

## Installation

### Prerequisites

* **Go:** This project requires Go. You can download and install it from the official Go website: [https://golang.org/dl/](https://golang.org/dl/)

### Dependencies

This project uses Go Modules to manage dependencies. To install the necessary dependencies, run the following command in your project's root directory:

```bash
go mod tidy
```

## How to Use

1.  Run the application from your terminal.
2.  You will be greeted with the `Pokedex >` prompt.
3.  Enter one of the commands listed below.

## Commands

* **`help`**: Displays a help message with a list of all available commands.
* **`map`**: Shows the next 20 locations in the Pokemon world.
* **`mapb`**: Shows the previous 20 locations.
* **`explore <location_name>`**: Lists all the Pokemon found in a specific location.
* **`catch <pokemon_name>`**: Attempts to catch a Pokemon.
* **`inspect <pokemon_name>`**: Shows details about a caught Pokemon, including its name, height, weight, stats, and types.
* **`pokedex`**: Lists all the Pokemon you have caught.
* **`exit`**: Exits the Pokedex REPL.