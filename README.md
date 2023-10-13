# go-webapp
Simple web app using golang and docker multistage builds.

Pre-requisite:
- Any IDE
- Golang 
- Docker

Build and run docker image locally:
    
    cd go-webapp
  
    docker build -t go-webapp:v1 .
    
    docker run -p 80:80 -it go-webapp:v1


Open Browser and try below API:

    http://localhost:80/
    http://localhost:80/getUser