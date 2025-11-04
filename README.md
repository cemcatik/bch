# bch

A simple command-line tool for generating and verifying bcrypt password hashes.

## Features

- Generate bcrypt hashes from passwords with configurable work factor
- Verify passwords against existing bcrypt hashes
- Secure password input (no echo to terminal)
- Fast and lightweight

## Installation

### Prerequisites

- Go 1.25.0 or later

### Install from source

```bash
go install github.com/cemcatik/bch@latest
```

Or clone and build locally:

```bash
git clone https://github.com/cemcatik/bch.git
cd bch
go build -o bch bch.go
```

## Usage

### Generate a hash

Generate a bcrypt hash with the default work factor (12):

```bash
bch hash
```

You'll be prompted to enter a password, which will be hashed and printed to stdout.

### Generate a hash with custom work factor

Use the `-f` or `--factor` flag to specify a different work factor:

```bash
bch hash --factor 14
```

**Work factor guidelines:**

- Default: 12 (recommended for most applications)
- Higher values = more secure but slower
- Each increment doubles the computational cost
- Range: 4-31 (bcrypt limitation)

### Verify a hash

Verify a password against an existing bcrypt hash:

```bash
bch verify '$2a$12$R9h/cIPz0gi.URNNX3kh2OPST9/PgBkqquzi.Ss7KIUgO2t0jWMUW'
```

You'll be prompted to enter a password. The tool will print `true` if the password matches, or `false` if it doesn't.

## Examples

### Example: Generate hash

```bash
$ bch hash
Enter password:
$2a$12$R9h/cIPz0gi.URNNX3kh2OPST9/PgBkqquzi.Ss7KIUgO2t0jWMUW
```

### Example: Generate hash with work factor 14

```bash
$ bch hash -f 14
Enter password:
$2a$14$xyz...
```

### Example: Verify hash (correct password)

```bash
$ bch verify '$2a$12$R9h/cIPz0gi.URNNX3kh2OPST9/PgBkqquzi.Ss7KIUgO2t0jWMUW'
Enter password:
true
```

### Example: Verify hash (incorrect password)

```bash
$ bch verify '$2a$12$R9h/cIPz0gi.URNNX3kh2OPST9/PgBkqquzi.Ss7KIUgO2t0jWMUW'
Enter password:
false
```

## Security Notes

- Passwords are read securely from stdin without echoing to the terminal
- The default work factor of 12 is considered secure for most applications as of 2025
- Bcrypt automatically handles salting - each hash is unique even for the same password
- Higher work factors provide better protection against brute-force attacks but increase CPU usage

## Development

### Build

```bash
go build -o bch bch.go
```

### Format code

```bash
go fmt ./...
```

### Update dependencies

```bash
go get -u ./...
go mod tidy
```

## License

MIT License - see [LICENSE](LICENSE) file for details.
