# htmxual ✨
htmxual is a demonstration application for building frontends without JavaScript. It utilizes [Go Gin](https://github.com/gin-gonic/gin) as the web server, [Templ](https://github.com/a-h/templ) (a library for JSX-style templating in Go), and [HTMX](https://github.com/bigskysoftware/htmx) in the frontend. 🚀

The application showcases a simple counter example accessible via /count/ and a todo list example on /todo/. 📝

## Features 🌟
Counter Example: Visit /count/ to see a basic counter implemented without using JavaScript.

Todo List Example: Access /todo/ to explore a simple todo list demonstration.

## Technologies Used 💻
Go Gin: Web framework used for the backend.

Templ: Library for JSX-style templating in Go.

HTMX: Used in the frontend to enable interactions without traditional JavaScript.

## Getting Started 🏁
To get the project up and running locally, follow these steps:

## Prerequisites: Ensure you have Go installed. ✅
Clone the Repository: `git clone https://github.com/niklastreml/htmxual.git`

Navigate to the Directory: `cd htmxual`

Run the Application: Execute `go run ./cmd/app` in your terminal.

Access the Application: Open your browser and go to http://localhost:8080/. 🌐

## Developing the Project ▶️

To run `htmxual`, you'll need to use [Air](https://github.com/cosmtrek/air) to streamline the development process. Follow these steps:

### Prerequisites ✔️

Make sure you have Go installed on your machine.

### Installation 🛠️

1. Install Air by running: 
    ```
    go get -u github.com/cosmtrek/air
    ```

### Starting the Application 🚀

1. Once Air is installed, navigate to the project directory.
2. In your terminal, run the following command:
    ```
    air
    ```
3. Air will watch for file changes and automatically rebuild and restart the server.

4. Access the application by opening your browser and going to `http://localhost:8080/`.

### Usage 🎯

- Access `/count/` to view the counter example.
- Visit `/todo/` to explore the todo list functionality.

## Contributing 🤝

Contributions are welcome! If you'd like to add features, fix bugs, or improve the documentation, please feel free to submit pull requests.

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

## License 📜

This project is licensed under the MIT License - see the `LICENSE` file for details. 📄
