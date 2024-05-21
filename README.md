# URL-Shortener

A repository that can host a URL shortener service.

Each URL will be converted into a short URL with **1** to **7** hash characters, supporting a maximum of **2048** URL generations per second, and designed to run for up to 15 years. The service can handle up to 4 machines simultaneously.  
To achieve this, the ID generator is implemented by slightly modifying the [bwmarrin/snowflake][SnowFlake] package.

This project is implemented in **Golang** and follows the clean architecture principles. You can quickly set up a complete example using Docker and Docker Compose. The repository provides all necessary configurations and scripts to get the service up and running, ensuring a scalable and maintainable codebase. 

## About The Project

This project is intended to practice and enhance my backend development skills, such as Golang, Clean Architecture, Docker, Nginx, and more.  

The system's specifications are designed based on the requirements mentioned in Chapter 8 of this [System Design Project][System Design].  

Through this project, I aim to deepen my understanding of various backend technologies and apply them in a practical context 🐱‍🚀.

## Getting Started

You can start this project by following the steps below.  

Most basic use cases are provided in the examples/ directory.  
You can also create your custom image by editing the Dockerfile in the docker/ directory.  

Make sure all [Prerequisites](#prerequisites) are installed.

### Prerequisites
- [Golang][Go]
  
  > develop on go version go1.22.2 windows/amd64
  
- [Docker] (Optional)
  
  > Install this only if you want to try the docker example.  
  > Docker Desktop 4.12.0, Docker Compose version v2.10.2 I used.

### Host a URL shortener service locally
You can just clone the repository.
```
git clone https://github.com/birdmandayum0131/Url_Shortener.git
```
#### docker
  run the docker compose example.
  ```
  sh ./examples/docker/docker-run.sh
  ```
  if you have a gateway container(docker network) running, like nginx server with https.
  you can join the network by setup the environment variable of `GATEWAY_NETWORK` which called by docker example.
#### windows
  If you don't want use docker, and you have a mysql server exist.  
  
  1. Setup necessary environment variable or provide a env.cmd under examples/windows.  
      
      You can write in format like:
      ```
      export DB_USER=${USERNAME}
      export DB_PASSWORD=${PASSWORD}
      ```
  2. run windows example  
      ```
      sh ./examples/windows/run.sh
      ```  
## Timeline
  - start the project `2024/04/16`
  - successfully complete the basic implementation `2024/04/18`
  - studying and refactor it to clean architecture ~ `2024/05/09`
  - add docker support to this project `2024/05/10`
  - add docker compose example to Host a URL shortener service from scratch `2024/05/16`
  - refactor configs to yaml file `2024/05/21`
    
## Roadmap
  - [x] Clean Architecture
  - [x] Docker
  - [x] Docker Compose
  - [x] Nginx
  - [ ] Goroutine
  - [ ] CI/CD (Jenkins?)
  - [ ] Load Testing

[SnowFlake]:                        https://github.com/bwmarrin/snowflake                                                                 "bwmarrin/snowflake"
[System Design]:                    https://github.com/Admol/SystemDesign/blob/main/CHAPTER%208%EF%BC%9ADESIGN%20A%20URL%20SHORTENER.md   "Admol/SystemDesign"
[Go]:                               https://go.dev/                                                                                       "Golang"
[Docker]:                           https://www.docker.com/                                                                               "Docker"
