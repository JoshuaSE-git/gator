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

macOS (no password, your username): `postgres://wagslane:@localhost:5432/gator`
Linux: `postgres://postgres:postgres@localhost:5432/gator`
Test your connection string by running psql, for example:

```bash
psql "postgres://wagslane:@localhost:5432/gator"
```

It should connect you to the gator database directly. If it's working, great. exit out of psql and save the connection string in your .gatorconfig.json file.

```json
{
  "db_url": "postgres://example"
}
```


## ⚙️ Commands

| Command  | Description                                           |
| -------- | ----------------------------------------------------- |
| `add`    | Add a new expense entry                               |
| `delete` | Delete entries by ID, date, category, or amount range |
| `edit`   | Edit an existing expense by ID                        |
| `list`   | View filtered and sorted expenses                     |
| `undo`   | Undo recent changes (deletions, edits, adds)          |

| Flags                        | Available Commands              | Description                       |
| ---------------------------- | ------------------------------- | --------------------------------- |
| `--date, -d`                 | `add`, `delete`, `edit`, `list` | Date in YYYY-MM-DD format         |
| `--category, -c`             | `add`, `delete`, `edit`, `list` | Expense category (ex. "Gaming")   |
| `--description, -s`          | `add`, `delete`, `edit`, `list` | Expense description (ex. "Dota2") |
| `--id, -i`                   | `delete`, `edit`, `list`        | Expense id (ex. 1)                |
| `--amount, -a`               | `delete`, `edit`, `list`        | Expense amount (ex. 10.32)        |
| `--min-amount, --max-amount` | `delete`, `list`                | Min/Max amount                    |
| `--min-date, --max-date`     | `delete`, `list`                | Min/Max date                      |

Use wallet [command] --help to see full options and flag descriptions.

## ▶️ Usage Examples

### ➕ Adding Expenses

```bash
wallet add 10.24
wallet add 10.24 --date "2027-07-08" --category "Food" --description "Wendys"
wallet add 10.24 --category "Food" -s "Wendys"
wallet add 10.24 -c "Food" -s "Wendys"
```

### ❌ Deleting Expenses

```bash
wallet delete --id 1 
wallet delete --id 1 5 12 8
wallet delete --category "Food"
wallet delete --date 2027-07-08
wallet delete --min-date 2027-06-01 --max-date 2027-07-01 --category "Food"
wallet delete --min-amount 10.23 --max-amount 23.23
```

### ✏️ Editing Expenses

```bash
wallet edit --id 1 -c "Gaming" --date 2027-05-05 --description "League"
wallet edit -i 13 --amount 24.67
wallet edit --id 12 --s "Starbucks"
```

### 📋 Listing Expenses

```bash
wallet list --day
wallet list --month 2027-07 --category Food --min-amount 5.00
wallet list --year 2027 --sort-by amount --desc
```

## 📄 License

[MIT License](https://github.com/JoshuaSE-git/wallet-watcher-cli/blob/main/LICENSE) © Joshua Emralino
