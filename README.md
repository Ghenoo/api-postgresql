# API-POSTGRESQL

Esse projeto é uma implementação de uma API utilizando Golang + PostgreSQL, incluindo levantar um container PostgreSQL no docker, criar tabela, setar permissões e testar tudo isso no postman.


---

## Instalação

Parar usar esse projeto, você precisará realizar esses passos:

1. Clone o repositório: `git clone https://github.com/Ghenoo/api-postgresql.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`
5. Deploy :`docker run -d --name api-todo -p 5432:5432 -e POSTGRES_PASSWORD=senha postgres:13.5`




## Docker and Docker Compose

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment. Here are the most common Docker and Docker Compose commands you may want to use:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Run a container based on the built image. Replace `your-image-name` with the name you used when building the image. You can change the port number if necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Build the services defined in the `docker-compose.yml` file.
- `docker compose up`: Run the services defined in the `docker-compose.yml` file.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```sh
docker-compose down
```

Para mais informações sobre Docker e Docker Compose, entre no link da documentação oficial:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)


## Ferramentas Usadas

Esse projeto uso e seguiu as seguintes ferramentas:

- [Golang](https://golang.org/) for backend development
- [PostgreSQL](https://postgresql.org/) for backend development

## Usage
Once the API is running, you can use postman to perform tests on the port `http://localhost:8080/`

## Contributing

To contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them using Conventional Commits
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

## Créditos

Esse projeto foi inspirado no [thiaguinho](https://www.youtube.com/@AprendaGolang).
