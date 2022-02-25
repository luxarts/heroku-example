# Purpose of this repository
Know how to create a Go application, deploy to Heroku (automatically) and configure your domain to point the domain provided by Heroku.

## Things you will learn
- Use Git and GitHub
- Create a simple Go backend server
- Use Docker
- Use Heroku platform
- DNS configuration

# Requirements
- Go
- Heroku CLI
- Git
- Docker
- Domain (optional)

---

# Instructions
## Create a project in GitHub
1. Go to [www.github.com](www.github.com) and click "New"
2. Check _Add README_ and _Add .gitignore_ and select _Go_ from the list

## Configure your SSH keys for GitHub

## Clone the repository
```
git clone git@github.com:<YOUR_USERNAME>/<YOUR_APP>.git
```
## Write the code
We will write a simple server with single endpoint.
The endpoint will be `GET /ping` and it will response a `200 OK` status code and the message `"pong"`.
The folder structure will be:
```
|-cmd
| |-main.go         // Contains the entry point
|
|-internal
| |-router
|   |-router.go     // Contains the router initialization
|   |-endpoints.go  // Contains all the endpoints
|
|-defines
| |-paths.go        // Contains the paths of each endpoint
```

1. Init the module
   ```
   go mod init <YOUR_APP>
   ```
2. Download `gin-gonic` library
   ```
   go get -u github.com/gin-gonic/gin
   ```
## Create the Dockerfile
1. Create a file with name `Dockerfile` in the root of the project with the following content
   ```Dockerfile
   # Use alpine (very lightweight linux distro) with Go installed as a base image
   FROM golang:rc-alpine3.14 AS builder

   # Create a folder called build
   RUN mkdir /build
   # Change the working directory to the new folder
   WORKDIR /build
   # Copy all the contents to the new folder
   COPY . .
   # Build binary
   RUN go build -o output ./cmd/main.go
   
   # Use alpine as a base image
   FROM alpine
   # Create a user
   RUN adduser -S -D -H -h /app appuser
   # Change current user
   USER appuser
   # Copy the binary created before to a new folder
   COPY --from=builder /build/output /app/
   # Change the working directory to the new folder
   WORKDIR /app
   # Execute the binary
   CMD ["./output"]
   ```
2. Build docker image
```
docker build -t <YOUR_APP> .
```
3. Run the image
```
docker run --rm -p 8080:8080 <YOUR_APP>
```
## Create the project in Heroku
1. Go to [Heroku Dashboard](https://dashboard.heroku.com/apps) and click on _New_ > _New app_
2. Go to _Deploy_ tab > _Deployment method_ > Github
3. Login with GitHub and search for <YOUR_APP>
4. Click on _Connect_
5. Click _Enable Automatic Deploys_
## Configure your app to use containers
1. Login in Heroku CLI
   ```
   heroku login
   ```
2. Set stack to container
   ```
   heroku stack:set container --app <YOUR_APP>
   ```
## Create the heroku.yml
1. Create a file with name `heroku.yml` in the root of the project with the following content
   ```yml
   build:
     docker:
       web: Dockerfile
   ```

## Push your local changes
```
git add .
git commit -m "Initial version"
git push
```

## Configure your app to point to your domain

---

# Links
