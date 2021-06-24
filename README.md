# FanFit iOS

## Architecture

- We run a micro service oriented architecture
- Each service is based on this template within the repo
- Details about how its designed can be found here:
  - https://medium.com/gdg-vit/clean-architecture-the-right-way-d83b81ecac6


## Installation

- You will need to install the following dependencies to effectively write code for this repository (note, if you are just looking to deploy the service, this step can be skipped since you will use docker)
- Swaggo: Testing APIs
  - go get -u github.com/swaggo/swag/cmd/swag
  - go get -u github.com/swaggo/gin-swagger
  - go get -u github.com/swaggo/files
- Sqlc: generates extremely useful go code based on sql code
  - Mac: brew install sqlc
- Go:  
- Docker:


## Deployment

1. Building the project:
  - Works fine right now but still a work in progress (looking for a way to build individual dependencies in docker instead of the whole thing when a single change in made)
  - Ensure all containers are stopped: docker-compose down
  - Build all containers: docker-compose up –build
  - If stopped and want to start again: docker-compose up
2. Accessing the PG-Admin window:
  -	Once docker has finished building you can open a PG-Admin window to see your database with nice UI
  -	Open PG-Admin: go to browser and navigate to localhost:16543
  -	Enter username and password: 
    - Username: test@gmail.com
    -	Password: test123!
  -	Right click on the “Server” and click “Create new server”
  -	Get the IP Address postgres is running on
    -Open a new terminal
    -Type “docker ps” and find the id of the container running postgres
    -Type command “docker inspect (id for postgres contatiner) | grep IPAddress”
    -Copy that to use as “Host Address” In PG-Admin
  - Enter the following when creating the server:
    - Under General -> Name: root
    - Under Connection:
    - Host Name: the IP Address for postgres
    - Port: 5432
    - Maintenance Database: root
    - Username: root
    - Password: root


