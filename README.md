# Dataserver
This project is a simple web application that displays promotions loaded from a CSV file in a Redis database.

# Requirements
Go (version 1.19 or higher)

Redis (version 7.2 or higher)

# Installation
Clone the repository: ` git clone https://github.com/andranikgh/dataserver.git `

Change into the project directory: ` cd dataserver `

Install dependencies: ` go mod tidy ./ `

Build the project: ` go build `
# Usage

Web Application

Start Redis: redis

Start the web application: ` ./dataserver `

Open a web browser and navigate to http://localhost:1321



Or just run using docker composer: `` docker-compose -up [-d] ``


License
This project is licensed under the MIT License.
