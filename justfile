set shell := ["bash", "-uc"]

default:
    @just --list

# Download and display example TOML file
setup:
    curl -O https://raw.githubusercontent.com/GGist/ssdp-rs/3d4dc17d63c0ec42b03b4ce8f07330a3352bc6d6/Cargo.toml

# Check TOML formatting and show what would be changed
check: setup
    dagger call run-toml-formatter --source=. --args="check","Cargo.toml" export --path=.

# Format TOML file in-place
format: setup
    dagger call run-toml-formatter --source=. --args="check","--fix-inplace","Cargo.toml" export --path=.
    @echo "Formatted file:"
    @cat Cargo.toml

# Run complete demo with setup, check and format
demo: setup check format

test:
    testscript tests
