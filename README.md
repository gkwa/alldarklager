# TOML Formatter Dagger Module

A Dagger module that provides TOML file formatting capabilities using the [toml-formatter](https://github.com/paulovcmedeiros/toml-formatter) tool.

## Features

- Format TOML files in-place
- Run arbitrary toml-formatter commands
- Debug mode for interactive use

## Prerequisites

- Dagger CLI installed
- Docker running

## Usage

### Format a TOML file

To format a TOML file in-place:

```shell
# Format and overwrite original file
dagger call check --source=. --filename=your-file.toml export --path=.
```

### Example

Here's an example using a sample Cargo.toml file:

```shell
# Download sample TOML file
curl -O https://raw.githubusercontent.com/GGist/ssdp-rs/3d4dc17d63c0ec42b03b4ce8f07330a3352bc6d6/Cargo.toml

# Check initial format
cat Cargo.toml

# Format the file
dagger call check --source=. --filename=Cargo.toml export --path=.

# Check the formatted result
cat Cargo.toml
```

### Run Custom Commands

You can run any toml-formatter command using the RunTomlFormatter function:

```shell
# Show help
dagger call run-toml-formatter --source=. --args="--help" export --path=.
```

### Debug Mode

For interactive debugging and experimentation:

```shell
dagger call debug --source=.
```

This will open an interactive shell in the container where you can run toml-formatter commands directly.

## Development

Building upon this module:

- The module uses Python 3.12 base image
- Poetry is used for dependency management
- The toml-formatter tool is installed directly from GitHub

To modify or extend the module, see the functions in `main.go`:

- CreateBaseContainer: Sets up the base Python container
- InstallPoetry: Installs Poetry package manager
- InstallProject: Sets up the project and installs toml-formatter
- RunTomlFormatter: Core function for running formatter commands
- Check: Convenience function for formatting files in-place
- Debug: Interactive debugging support
