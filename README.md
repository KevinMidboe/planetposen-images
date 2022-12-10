# Planetposen image proxy

App for fetching and uploading images to cloud bucket storage. Should also be a on-the-fly image formatting proxy.

## Setup

From source:

1. `git clone https://github.com/kevinmidboe/planetposen-images`
2. `cp .env.example .env`
3. Update variables in `.env` file
4. `make install`

## Run

Run api from command line with:

```bash
make server
```

Run as docker container using:

```bash
(sudo) docker run -d --name planetposen-images -p 8000:8000 \
    -e PORT=8000
    -e HOSTNAME=planet.schleppe.cloud
    -e GCS_BUCKET=planetposen-images
```

or

```bash
(sudo) docker run -d --name planetposen-images -p 8000:8000 --env-file .env planetposen-images
```
