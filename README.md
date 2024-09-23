# Groupie Tracker Search Bar

Groupie Tracker Search Bar is a web application that provides information about various music artists, including their locations, concert dates, and other related details. The application fetches data from a public API and presents it in a user-friendly interface, featuring a search bar that allows users to easily find specific artists by name or filter results based on location or concert dates.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [Templates](#templates)
- [Styling](#styling)
- [Contributing](#contributing)
- [License](#license)
- [Authors](#authors)

## Features

- View a list of artists along with their details.
- Search for specific artists.
- View detailed information about a selected artist.
- View the locations, concert dates, and relations between artists.
- Serve static files such as images and CSS.

## Installation

To run this project locally, follow these steps:

1. **Clone the repository:**

    ```bash
    git clone https://learn.zone01kisumu.ke/git/bernaotieno/groupie-tracker-search-bar.git
    ```

2. **Navigate to the project directory:**

    ```bash
    cd groupie-tracker-search-bar
    ```

3. **Install dependencies:**

    Since this is a Go project, ensure you have Go installed on your machine. You can check this by running:

    ```bash
    go version
    ```

    If Go is not installed, download and install it from [here](https://golang.org/dl/).

4. **Run the application:**

    ```bash
    go run .
    ```

5. **Open your browser and navigate to:**

    ```
    http://localhost:8081
    ```

## Usage

The web application has the following endpoints:

- `/` - Home page displaying a list of artists.
- `/artists` - Endpoint to fetch a list of all artists.
- `/locations` - Endpoint to fetch a list of all locations.
- `/dates` - Endpoint to fetch a list of all concert dates.
- `/relations` - Endpoint to fetch relations between artists.
- `/artistInfo` - Endpoint to fetch and display detailed information about a specific artist.
- `/static/` - Serves static files (e.g., images, CSS).

## Project Structure

The project is structured as follows:

```
├── api/
│ ├── api.go # Contains functions to fetch data from the API
│
├── cmd/
| ├──cmd.go # HAndles functions and listening to server
├── handlers/
│ ├── handlers.go # Contains HTTP handlers for different routes
│ 
├── models/
│ ├── models.go # Defines the data models  and has the search functionalities
│ 
├── static/
│ ├── Static files (i.e. images, CSS) and the JavaScript file.
│ 
├── templates/
│ ├── Home.html # Template for the home page
│ ├── artistPage.html # Template for artist details page
| ├── errorPage.html # Template for error messages
│ 
├── main.go # Entry point of the application
|
├── go.mod # Go module file
|
└── README.md 
```


## API Endpoints

The following endpoints are used to fetch data from the Groupie Tracker API:

- **Artists:** `https://groupietrackers.herokuapp.com/api/artists`
- **Locations:** `https://groupietrackers.herokuapp.com/api/locations`
- **Dates:** `https://groupietrackers.herokuapp.com/api/dates`
- **Relations:** `https://groupietrackers.herokuapp.com/api/relation`

The `api.go` file in the `api/` directory contains the logic for fetching and processing this data.

## Templates

- **Home.html:** The main page that lists all the artists.
- **artistPage.html:** The detailed page for a specific artist.
- **errorPage.html:** The Page that is shown when an error occurs.

These templates are stored in the `templates/` directory and are rendered using Go's `html/template` package.

## Styling

The application uses custom CSS for styling, located in the `static/` directory. The main styles are defined in the `style.css` file.

## Contributing

If you'd like to contribute to this project, feel free to open an issue or submit a pull request. Please ensure that your code follows the existing style and includes tests where appropriate.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Authors

**[Bernad Okumu](https://learn.zone01kisumu.ke/git/bernaotieno)**

**[Raymond Caleb](https://learn.zone01kisumu.ke/git/rcaleb)**
