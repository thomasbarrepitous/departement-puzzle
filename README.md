# Guess the departement !

This is a simple game that demonstrates a Go + HTMX use case. It is lightweight, requires few ressources and is fun to play !

![Screenshot](/screenshot.png)

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/thomasbarrepitous/departement-puzzle.git
    cd departement-puzzle
    ```

2. **Build and run the Docker container:**

    ```bash
    docker-compose up
    ```

    This command will build the Docker image and start the container. You can access the application at [http://localhost:8080](http://localhost:8080).

3. **To stop the container, press `Ctrl + C` in the terminal where `docker-compose up` is running. Alternatively, run:**

    ```bash
    docker-compose down
    ```

## Customization

- If your Go application has external dependencies, update the Dockerfile or Docker Compose configuration accordingly.
- Adjust the `docker-compose.yml` file based on your application's specific requirements.

## Contributing

If you'd like to contribute to this project, feel free to open an issue or submit a pull request. Contributions are welcome!
