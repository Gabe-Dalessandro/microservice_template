# FanFit iOS

## Architecture

- We run a micro service oriented architecture
- Each service is based on this template within the repo
- Details about how its designed can be found here:
  - https://medium.com/gdg-vit/clean-architecture-the-right-way-d83b81ecac6<br><br>



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
    - Once docker has finished building you can open a PG-Admin window to see your database with nice UI
    - Open PG-Admin: go to browser and navigate to localhost:16543
    - Enter username and password: 
          - Username: test@gmail.com
          -	Password: test123!
    -	Right click on the “Server” and click “Create new server”
    -	Get the IP Address postgres is running on
          - Open a new terminal
          - Type “docker ps” and find the id of the container running postgres
          - Type command “docker inspect (id for postgres contatiner) | grep IPAddress”
          - Copy that to use as “Host Address” In PG-Admin
    - Enter the following when creating the server:
          - Under General -> Name: root
    - Under Connection:
          - Host Name: the IP Address for postgres
          - Port: 5432
          - Maintenance Database: root
          - Username: root
          - Password: root

3. Using Swagger to Test APIs:
    - o	Go to browser and navigate to http://localhost:8080/swagger/index.html#




## File Structure

- **api:** Holds all code related to routes and endpoints. Uses code from other folders
- **docs:** tests our API endpoints using Swaggo (a version of Swagger)
    - This code is generated based on code throughout the repository
    - Swaggo Documentation: https://docs.sqlc.dev/en/latest/tutorials/getting-started.html
- **models:** uses sqlc library to generate go code based on SQL code
    - Sqlc documentation: https://docs.sqlc.dev/en/latest/tutorials/getting-started.html
    -	Holds the different structs and queries that will be used inside of the service
    - Can be broken out into different tags that use different data structres in the service
    - Example: UserService might have the “Creators” and “Clients” within the models folder
          - Both directories need you to call “sqlc generate” to create the data structures
    - Repository: Uses SQL terminology for functions
    - Service: Higher level business logic functions that use the SQL functions (More business readable)
- **server:** code to generat the server and has the go routine to continuously run
- **sql:** Has the SQL code that generates the tables in Postgres for this particular service. Also can use inserts here as well.
-	**application.go:** Creates the services, repositories (from the model layer), routers, and server. Also starts each piece as well
-	**docker-compose.yml:** Script that automatically builds the entire Service, Postgres DB, and launches a PG-Admin screen so you have a nice UI to see database changes
-	**dockerfile:** Script that launches everything for you
-	**makefile:** contains other build process that may help you generate different tools (i.e. swagger)



## Tool

-	**Swagger:** Used to test APIs
    -	Documentation: https://github.com/swaggo/swag
    -	Used in the browser to test endpoints
    -	Found at localhost:8080/swagger/index.html
    -	Code is set up in server/server.go
    -	To find the “swag” command in order to generate the swagger.json/yaml: 
          -	export PATH="$PATH:$HOME/go/bin"
    - To use “Swaggo”: 
          - make swagger
    - If you are getting the old index.html page, you must delete your cache
          - SHIFT + (click the refresh page button)
-	**Sqlc:** generates extremely useful go code used within the repository
    -	Documentation: https://docs.sqlc.dev/en/latest/tutorials/getting-started.html
    -	Write SQL code within the models/tag/repository directory
    -	Then use command “sqlc compile” to ensure you wrote correct SQL code
    -	Use command: “sqlc generate” to create the go files
    -	These files have structs and functions that can be used throughout the repo



## Acknowledgments
-	**Harshil Mavani:** Kickass, Business Savy CTO
-	**Jason Gomez:** Young, Hungry and Smart developer
-	**Gabe Dalessandro:** Developer

