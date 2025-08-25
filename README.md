# gator

A fast, no-nonsense RSS reader CLI written in Go. Fetch RSS feeds, store posts in Postgres, browse your personalized timeline, and keep the aggregator running in the background.

## Requirements

- Go 1.21+
- PostgreSQL 14+

## Install

You can install the CLI directly with Go:

```bash
# From anywhere
go install github.com/heretic1321/gator/cmd/gator@latest
```

This builds a static binary and places it in your $GOPATH/bin (or $GOBIN). Ensure that directory is on your PATH so you can run `gator` from your shell.

## Configure

gator uses a simple JSON config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://USER:PASSWORD@HOST:PORT/DBNAME?sslmode=disable",
  "current_user_name": ""
}
```

- db_url: standard Postgres connection string
- current_user_name: will be set by the `login` command

Apply database migrations using your preferred tool (e.g., goose). Migrations are in `sql/schema/`.

## Usage

Run the CLI:

```bash
gator <command> [args]
```

Key commands:

- register <username>: Create a new user
- login <username>: Set the current user in the config
- users: List users (current indicated)
- addfeed <name> <url>: Add a feed and follow it
- feeds: List all feeds
- follow <url>: Follow a feed by URL
- unfollow <url>: Unfollow a feed by URL
- following: Show feeds the current user follows
- agg <duration>: Start the aggregator loop (e.g. `agg 1m`, `agg 10s`)
- browse [limit]: Browse recent posts from followed feeds (default limit 2)
- reset: Danger! Deletes all users (and cascades)

Examples:

```bash
# One-time setup
createdb gator
# Update ~/.gatorconfig.json with your DB URL

# Register and login
gator register alice
gator login alice

# Add and follow feeds
gator addfeed "TechCrunch" https://techcrunch.com/feed/
gator follow https://news.ycombinator.com/rss

# Run the aggregator (never-ending loop)
gator agg 1m

# In another terminal, browse posts
gator browse 10
```

## Aggregator

The `agg` command runs forever, every N duration:

- Picks the next feed to fetch (oldest last_fetched_at first)
- Marks it fetched
- Downloads and parses the RSS feed
- Stores posts (unique by URL)

Use Ctrl+C to stop. Be respectful of remote servers: choose reasonable durations (e.g., ≥ 30s).

## Development

During development you can run:

```bash
make run
```

This runs `go run` on the CLI entrypoint. For production usage prefer `go build` or `go install`. After building, the binary can be run without the Go toolchain.

## Project Layout

- cmd/gator/: CLI entrypoint
- internal/commands/: Commands and middleware
- internal/database/: sqlc-generated queries and helpers
- sql/: SQL schema and queries
- pkg/rss/: RSS client and XML parsing types

## Troubleshooting

- Ensure ~/.gatorconfig.json exists and db_url is correct
- Verify Postgres is running and accessible
- If you see rate limits or slow responses, increase the aggregator duration

## License

MIT
