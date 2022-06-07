# Backend application for IDLER chat service

###
- go 1.18.2
- docker

For application need EnvFile by Borys Pierov plugin by .env file which contains:
```dotenv
POSTGRES_VERSION=[your postgres version here]
POSTGRES_PORT=[your postgres port here]
POSTGRES_SCHEMA=idler
POSTGRES_USER=[your postgres user here]
POSTGRES_PASSWORD=[your postgres password here]

LIQUIBASE_VERSION=[your liquibase version here]

REDIS_VERSION=[your redis version here]
REDIS_PORT=[your redis port here]
REDIS_PASSWORD=[your redis password here]

MONGO_VERSION=[your mongo version here]
MONGO_HOST=host.docker.internal
MONGO_PORT=[your mongo port here]
MONGO_DB=idler
MONGO_INITDB_ROOT_USERNAME=[your mongo username here]
MONGO_INITDB_ROOT_PASSWORD=[your mongo password here]

PASSWORD_SALT=[your salt here]
JWT_SIGNING_KEY=[your signing here]
```