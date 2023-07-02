<p align="center">
  This is @kmakashe and @SofyaOspanova Ascii-Art-WEB-Dockerize project
</p>

## 🧐 About <a name="about"></a>

This project aims to dockerize the Ascii-Art-WEB application, which allows users to generate ASCII art banners using different styles. By following the instructions below, you will learn how to containerize the application using Docker while adhering to good coding practices and Dockerfile best practices.

### Objectives

The objectives for this project are as follows:
- Create at least one Dockerfile.
- Build an image based on the Dockerfile.
- Run a container using the image.

### Usage
After successfully building the Docker image, you can create and run a Docker container using the image. Use the following command:

`docker container run -p 8080:8080 ascii-go`

Additional Docker Commands:

Here are some additional Docker commands that might be useful:

- Build the Docker image:
`docker image build -t ascii-go .`

- Check the status of running containers:
`docker ps -a`

-  Remove a specific container (replace CONTAINER_ID with the actual ID):
`docker rm CONTAINER_ID`

- View the list of Docker images:
`docker images`

Accessing the Application

Once the Docker container is running, you can access the Ascii-Art-WEB-Stylize application by opening a web browser and navigating to http://localhost:8080.

## Prerequisites

- [![My Skills](https://skillicons.dev/icons?i=vscode)](https://skillicons.dev) VScode
- [![My Skills](https://skillicons.dev/icons?i=go)](https://go.dev/) [Go](https://go.dev/) 1.20

## ⛏️ Built Using <a name="built_using"></a>

- [![My Skills](https://skillicons.dev/icons?i=go)](https://skillicons.dev) [Go](https://go.dev/) - Programming language

## ✍️ Authors <a name="authors"></a>

- [@kmakashe](https://01.alem.school/git/kmakashe)
- [@SofyaOspanova](https://01.alem.school/git/SofyaOspanova)

## 🎉 References and Inspiration <a name="references"></a>

- [ASCII Table](https://www.alpharithms.com/ascii-table-512119/)
- ASCII Art Archive
- FIGlet font library
