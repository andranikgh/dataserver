# dataserver
Redis-backed Data Server in Golang
Project Name
This project is a simple web application that displays promotions loaded from a CSV file in a Redis database. It also includes a command-line utility for loading the CSV file into Redis.

Requirements
Go (version X or higher)
Redis (version X or higher)
Installation
Clone the repository: git clone https://github.com/your-username/your-project.git
Change into the project directory: cd your-project
Install dependencies: go get ./...
Build the project: go build
Usage
Web Application
Start Redis: redis-server
Start the web application: ./your-project
Open a web browser and navigate to http://localhost:8080
Command-line Utility
Start Redis: redis-server
Load the CSV file into Redis: ./your-project load-promotions path/to/csv/file.csv
Contributing
If you would like to contribute to this project, please create a pull request with your changes.

License
This project is licensed under the MIT License.