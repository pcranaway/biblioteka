package env

import env "github.com/Netflix/go-env"

type Environment struct {

    Postgres struct {
        Host      string `env:"POSTGRES_HOST"`
        Username  string `env:"POSTGRES_USERNAME"`
        Password  string `env:"POSTGRES_PASSWORD"`
    }

    MeiliSearch struct {
        Host      string `env:"MEILISEARCH_HOST"`
        APIKey    string `env:"MEILISEARCH_ACCESSKEY"`
    }

    MinIO struct {
        Host      string `env:"MINIO_HOST"`
        AccessKey string `env:"MINIO_ACCESS_KEY"`
        SecretKey string `env:"MINIO_SECRET_KEY"`
    }

}

func Load() *Environment {
    var e Environment

    _, err := env.UnmarshalFromEnviron(&e)
    if err != nil {
        panic(err)
    }

    return &e
}
