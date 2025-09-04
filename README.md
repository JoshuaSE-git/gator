# 🐊 Gator - An RSS Aggregator

> Gator is a CLI application that allows you to track and browse RSS feeds.

## 📋 Pre-requisites

You will need to install Go and PostgreSQL.

### Go Installation

Follow the instructions for your OS [here](https://go.dev/doc/install)

Ensure installation was successful.

```bash
go version
```

### PostgreSQL Installation

**macOS with [brew](https://brew.sh/)**

```bash
brew install postgresql@15
```

**Linux / WSL (Debian)**

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

Ensure installation was successful. "psql" is the default client for PostgreSQL.

```bash
psql --version
```

## 📦 Installation

1) Clone the repository.

```bash
git clone https://github.com/JoshuaSE-git/gator
```

2) Change directory to cloned repository.

```bash
cd gator
```

3) Install the binary using `go install`.

```bash
go install
```

4) Test installation with `gator` command.  The application should panic.

```bash
gator
```

5) Create config file.

Gator uses a simple json config file to manage users and database connection.  Simply create a .gatorconfig.json file in your home directory.

```bash
touch ~/.gatorconfig.json
```

6) Databse connection string.

Get your connection string. A connection string is just a URL with all of the information needed to connect to a database. The format is:
`protocol://username:password@host:port/database`

Here are examples:

- macOS (no password, your username): `postgres://wagslane:@localhost:5432/gator`
- Linux: `postgres://postgres:postgres@localhost:5432/gator`

Test your connection string by running psql, for example:

```bash
psql "postgres://wagslane:@localhost:5432/gator"
```

It should connect you to the gator database directly. If it's working, great. Exit out of psql and save the connection string in your .gatorconfig.json file.

```json
{
  "db_url": "postgres://example"
}
```


## ⚙️ Commands

| Command  | Description                                           |
| -------- | ----------------------------------------------------- |
| `register`    | Register a new user                              |
| `login` | Login as an existing user |
| `addfeed`   | Add a new RSS feed (name and url required)                      |
| `feeds`   | View all existing RSS feeds                   |
| `users`   | View all registed users          |
| `follow`   | Follow a registered RSS feed (url required)         |
| `following`   | View all followed RSS feeds     |
| `unfollow`   | Unfollow an RSS feed    |
| `browse`   | View most recent posts from followed RSS feeds (default result size is 2)   |
| `agg`   | Run this command to aggregate registered RSS feeds (required time string e.g. '30s')   |
| `reset`   | Resets the database  |

## 📄 License

[MIT License](https://github.com/JoshuaSE-git/wallet-watcher-cli/blob/main/LICENSE) © Joshua Emralino

