# Backend application IDLER chat service

###
- GO 1.18.2
- GIN
- PGX
- GCACHE

For application need EnvFile by Borys Pierov plugin and .env file which contains:
```dotenv
POSTGRES_VERSION=14
POSTGRES_PORT=[your postgres port here]
POSTGRES_SCHEMA=[your postgres schema here]
POSTGRES_USER=[your postgres user here]
POSTGRES_PASSWORD=[your postgres password here]

LIQUIBASE_VERSION=4.11

MONGO_VERSION=4.4.6
MONGO_HOST=[your mongo host here]
MONGO_PORT=[your mongo port here]
MONGO_DB=[your mongo db here]
MONGO_INITDB_ROOT_USERNAME=[your mongo username here]
MONGO_INITDB_ROOT_PASSWORD=[your mongo password here]

PASSWORD_SALT=[your salt here]
JWT_SIGNING_KEY=[your signing key here]
```

Command for building application
```dotenv
- make build
```
Command for running tests application
```dotenv
- make build
```

Command for running docker containers
```dotenv
- make docker
```

Swagger documentation http://localhost:9000/swagger/index.html
