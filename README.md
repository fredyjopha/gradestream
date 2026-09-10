# GradeStream

> Ingestion and tracking pipeline for assessment results for training organisations.

[![Go](https://img.shields.io/badge/Go-1.27+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/) [![TypeScript](https://img.shields.io/badge/TypeScript-native-3178C6?style=flat-square&logo=typescript&logoColor=white)](https://www.typescriptlang.org/) [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16+-4169E1?style=flat-square&logo=postgresql&logoColor=white)](https://www.postgresql.org/) [![Redis](https://img.shields.io/badge/Redis-7+-DC382D?style=flat-square&logo=redis&logoColor=white)](https://redis.io/) [![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com/) [![HTML5](https://img.shields.io/badge/HTML5-native-E34F26?style=flat-square&logo=html5&logoColor=white)](https://developer.mozilla.org/docs/Web/HTML) [![CSS3](https://img.shields.io/badge/CSS3-native-1572B6?style=flat-square&logo=css3&logoColor=white)](https://developer.mozilla.org/docs/Web/CSS)


## Prérequis

- Docker installé et fonctionnel
- Go 1.27.1

## Démarrage rapide

### 1. Lancer PostgreSQL localement

```bash
docker run --name gradestream-pg \
  -e POSTGRES_PASSWORD=devpass \
  -e POSTGRES_DB=gradestream \
  -p 5432:5432 \
  -d postgres:16
```

Vérifie que le conteneur tourne :

```bash
docker ps
```

### 2. Installer les dépendances Go

```bash
go get github.com/jackc/pgx/v5@latest
go mod tidy
```

### 3. Lancer l’application

```bash
go run ./cmd/gradestream
```

Si tout fonctionne, tu devrais voir :

```text
GradeStream — connexion PostgreSQL OK
```

## Configuration

La chaîne de connexion utilisée par défaut est :

```text
postgres://postgres:devpass@localhost:5432/gradestream
```
