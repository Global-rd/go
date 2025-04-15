# Dockerize a golang application

## Run via docker-compose

```bash
docker-compose up
```

## Run via docker

```bash
docker build -t health-checker .
```

```bash
docker run -p 8080:8080 health-checker
```
