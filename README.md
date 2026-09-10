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

| Variable | Valeur par défaut | Description |
|----------|-------------------|-------------|
| Hôte     | `localhost`       | Adresse du serveur PostgreSQL |
| Port     | `5432`            | Port exposé par Docker |
| Utilisateur | `postgres`     | Utilisateur PostgreSQL |
| Mot de passe | `devpass`     | Mot de passe (provisoire, à sécuriser plus tard) |
| Base de données | `gradestream` | Nom de la base créée au lancement |

# Migrations PostgreSQL locales

Cette feature ajoute les tables PostgreSQL nécessaires à Gradestream : `learners`, `cohorts` et `evaluations`.

## 1. Démarrer PostgreSQL

```bash
docker run --name gradestream-pg \
  -e POSTGRES_PASSWORD=devpass \
  -e POSTGRES_DB=gradestream \
  -p 5432:5432 \
  -d postgres:16
```

Cette commande démarre un conteneur PostgreSQL accessible sur le port `5432`. Les variables `POSTGRES_PASSWORD` et `POSTGRES_DB` configurent le mot de passe et la base initiale. 

## 2. Créer les migrations

```bash
mkdir -p internal/storage/postgres/migrations
```

Créer ensuite les fichiers suivants :

```text
internal/storage/postgres/migrations/
├── 0001_create_learners.sql
├── 0002_create_cohorts.sql
└── 0003_create_evaluations.sql
```

### `0001_create_learners.sql`

```sql
CREATE TABLE learners (
    id   TEXT PRIMARY KEY,
    name TEXT NOT NULL
);
```

### `0002_create_cohorts.sql`

```sql
CREATE TABLE cohorts (
    id   TEXT PRIMARY KEY,
    name TEXT NOT NULL
);
```

### `0003_create_evaluations.sql`

```sql
CREATE TABLE evaluations (
    id          TEXT PRIMARY KEY,
    learner_id  TEXT NOT NULL REFERENCES learners(id),
    cohort_id   TEXT NOT NULL REFERENCES cohorts(id),
    score       DOUBLE PRECISION NOT NULL,
    max_score   DOUBLE PRECISION NOT NULL,
    recorded_at TIMESTAMPTZ NOT NULL,
    UNIQUE (learner_id, cohort_id, recorded_at)
);
```

Les migrations doivent être exécutées dans l’ordre, car `evaluations` dépend des tables `learners` et `cohorts`.

## 3. Appliquer les migrations

```bash
docker exec -i gradestream-pg psql -U postgres -d gradestream < internal/storage/postgres/migrations/0001_create_learners.sql

docker exec -i gradestream-pg psql -U postgres -d gradestream < internal/storage/postgres/migrations/0002_create_cohorts.sql

docker exec -i gradestream-pg psql -U postgres -d gradestream < internal/storage/postgres/migrations/0003_create_evaluations.sql
```

- `docker exec` exécute une commande dans un conteneur actif.
- `-i` permet de transmettre le contenu des fichiers SQL via l’entrée standard.
- `psql` est le client PostgreSQL.
- `-U postgres` sélectionne l’utilisateur PostgreSQL.
- `-d gradestream` sélectionne la base de données.
- `< fichier.sql` envoie le contenu du fichier à `psql`.

## 4. Vérifier les tables

```bash
docker exec -it gradestream-pg \
  psql -U postgres -d gradestream -c "\dt"
```

L’option `-c` exécute une commande puis quitte. La commande `\dt` affiche les tables disponibles dans la base.

La structure générale à retenir est :

```text
docker exec [-i|-it] <conteneur> <commande> [arguments]
```

Utilise `-i` pour envoyer un flux et `-it` pour une session interactive.