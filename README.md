# Glintfed

Glintfed is a PixelFed-compatible service written in Go.

## Development

Run the minimal web application:

```bash
go run ./cmd/glintfed
```

The server listens on `:8080` by default and accepts `PORT` to override the port.

## Database Schema

Glintfed keeps a migration-first database workflow based on Pixelfed's live MySQL schema.

Sync the ent schemas from `.pixelfed-local`:

```bash
python3 scripts/sync-pixelfed-ent-schema.py
```
