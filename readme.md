# sockit

This is the code for the day 8 of the [Advent of Corona challenge](https://adventofcorona.hackersatupc.org/).

## Run

You may want to configure an `.env` file in the root directory with the following structure

```text
FLAG={some_flag}
SIZE={game_size}
TIMEOUT={timeout_in_seconds}
```

The defaults are:

```text
FLAG=FLAG
SIZE=3
TIMEOUT=3600
```

### *A pelo*

Just run it via `go` command, for developing, for example.
This will parse the root dir for an `.env` file.

```bash
go run .
```

### Docker

- Build `docker build -t sockit .`
- Run `docker run -p 8080:8080 sockit:latest`, or `docker run -p 8080:8080 --env-file .env sockit:latest` if run with an env file.

## Build from Docker Hub

If you are a lazy bastard, then just:

- Pull the image `docker pull tarasyarema/sockit:latest`
- Run the container `docker run -p 8080:8080 --env-file .env tarasyarema/sockit:tagname`
