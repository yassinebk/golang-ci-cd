# Go Web Server with CI/CD

A production-ready Go web server with comprehensive CI/CD pipeline, automated testing, and semantic versioning releases.

[![Go CI/CD Pipeline](https://github.com/username/repo/actions/workflows/ci.yml/badge.svg)](https://github.com/username/repo/actions/workflows/ci.yml)
[![Release & Deploy](https://github.com/username/repo/actions/workflows/cd.yml/badge.svg)](https://github.com/username/repo/actions/workflows/cd.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/username/repo)](https://goreportcard.com/report/github.com/username/repo)
[![Coverage Status](https://codecov.io/gh/username/repo/branch/main/graph/badge.svg)](https://codecov.io/gh/username/repo)

## Features

- **RESTful Endpoints**:
  - `/` - Home page showing hostname
  - `/health` - Health check endpoint with uptime
  - `/info` - Server information
  - `/metrics` - Basic server metrics

- **Production Ready**:
  - Comprehensive test suite
  - Docker support
  - Automated CI/CD pipeline
  - Semantic versioning
  - Security scanning

## Prerequisites

- Go 1.23.4 or higher
- Docker
- Git
- Make (optional)

## Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/username/repo.git
   cd repo
   ```

2. **Run locally**
   ```bash
   go run main.go
   ```

3. **Build and run with Docker**
   ```bash
   docker build -t webserver .
   docker run -p 80:80 webserver
   ```

4. **Run tests**
   ```bash
   go test -v ./...
   ```

## Development

### Project Structure

```
.
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── cd.yml
├── main.go
├── main_test.go
├── Dockerfile
├── .gitignore
├── .commitlintrc.yml
└── README.md
```

### Testing

Run all tests with coverage:
```bash
go test -race -coverprofile=coverage.out -covermode=atomic ./...
go tool cover -html=coverage.out -o coverage.html
```

### Commit Guidelines

This project follows [Conventional Commits](https://www.conventionalcommits.org/). Each commit message should be structured as follows:

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Code style changes
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Tests
- `build`: Build system
- `ci`: CI/CD
- `chore`: Maintenance
- `revert`: Revert changes

Example:
```
feat(api): add new metrics endpoint

Added a new endpoint that exposes basic server metrics.
Includes memory usage and goroutine count.

BREAKING CHANGE: Previous metrics format has been updated to include new fields.
```

### CI/CD Pipeline

The project includes a comprehensive CI/CD pipeline:

1. **CI Pipeline** (`ci.yml`):
   - Linting
   - Security scanning
   - Unit testing
   - Integration testing
   - Docker build verification

2. **CD Pipeline** (`cd.yml`):
   - Automatic semantic versioning
   - Release creation
   - Docker image publishing
   - Automated deployment

### Docker Support

Build the image:
```bash
docker build -t webserver .
```

Run the container:
```bash
docker run -p 80:80 webserver
```

## Configuration

The server can be configured through environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `80` |
| `LOG_LEVEL` | Logging level | `info` |

## API Documentation

### GET /
Returns a simple HTML page showing the hostname.

### GET /health
Returns server health status and uptime.

Response:
```json
{
  "status": "healthy",
  "uptime": "1h2m3s"
}
```

### GET /info
Returns detailed server information.

Response:
```json
{
  "hostname": "server-1",
  "timeStarted": "2024-12-23T10:00:00Z"
}
```

### GET /metrics
Returns basic server metrics.

Response:
```json
{
  "memoryAlloc": 1234567,
  "goroutines": 5
}
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Project Link: [https://github.com/username/repo](https://github.com/username/repo)