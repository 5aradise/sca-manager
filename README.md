# Spy Cat Agency management application

## Local Development

Make sure you're on Go version 1.23+.

Rename `.example.env` to `.env` and change environment variables you want.

### Run the project:

You can run both app and db in Docker:

```bash
docker-compose up -d
```

Or run db in Docker and app locally:

```bash
docker-compose up -d db
```

```bash
make run
```
