# curly-palm-tree

A Hello World application written in Go, designed for GitHub Codespaces development.

## Description

This is a simple Hello World application that demonstrates Go development setup with unit tests, linting, and debugging capabilities. The project is configured for use with GitHub Codespaces for an instant development environment.

## Prerequisites

- Go 1.25.5 or later
- (Optional) golangci-lint for code linting

## Getting Started

### Running the Application

To run the Hello World application:

```bash
go run main.go
```

Expected output:
```
Hello, World!
```

### Building the Application

To build an executable:

```bash
go build -o hello-world
./hello-world
```

### Running Tests

To run the unit tests:

```bash
go test -v
```

To run tests with coverage:

```bash
go test -v -cover
```

To generate a detailed coverage report:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Development with GitHub Codespaces

This repository is configured for GitHub Codespaces with the following features:

- **Go 1.25.5** development environment
- **VS Code Extensions**:
  - Go (golang.go) - Official Go extension with autocomplete and IntelliSense
  - Live Share (ms-vsliveshare.vsliveshare) - Real-time collaborative development
  - Atom Material Icons (Equinusocio.vsc-material-theme-icons) - File icons theme
  - Material Product Icons (PKief.material-product-icons) - Product icons theme
  - Catppuccin for VS Code (Catppuccin.catppuccin-vsc) - Color theme
  - Code Runner (formulahendry.code-runner) - Run code snippets quickly

### Opening in Codespaces

1. Navigate to the repository on GitHub
2. Click the "Code" button
3. Select "Codespaces" tab
4. Click "Create codespace on main" (or your branch)

The development environment will automatically:
- Install Go 1.25.5
- Download dependencies
- Install golangci-lint
- Configure autocomplete and IntelliSense
- Set up debug/run configurations

## Linting

This project uses golangci-lint for code quality checks.

### Install golangci-lint

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Run Linter

```bash
golangci-lint run
```

The linter configuration is in `.golangci.yml` and includes checks for:
- Code errors and bugs
- Code style and formatting
- Code complexity
- Security issues
- Common mistakes

## Debugging

VS Code launch configurations are provided in `.vscode/launch.json`:

1. **Launch Package** - Run the entire package
2. **Debug Package** - Debug the entire package with breakpoints
3. **Launch File** - Run the currently open file
4. **Debug Test** - Debug unit tests

To debug:
1. Open the file you want to debug
2. Set breakpoints by clicking left of the line numbers
3. Press `F5` or go to Run → Start Debugging
4. Select the appropriate launch configuration

## Project Structure

```
.
├── .devcontainer/
│   └── devcontainer.json     # GitHub Codespaces configuration
├── .vscode/
│   └── launch.json           # Debug/run configurations
├── .gitignore                # Git ignore rules
├── .golangci.yml             # Linter configuration
├── go.mod                    # Go module definition
├── main.go                   # Main application
├── main_test.go              # Unit tests
├── LICENSE                   # MIT License
└── README.md                 # This file
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2025 TerraformTestLab

