set shell := ["bash", "-uc"]

default:
    @just --list

# Download and display example TOML file
setup:
    curl -O https://raw.githubusercontent.com/GGist/ssdp-rs/3d4dc17d63c0ec42b03b4ce8f07330a3352bc6d6/Cargo.toml
    @echo "Original file:"
    @cat Cargo.toml

# Check TOML formatting without modifying file
check:
    dagger call check --source=. --filename=Cargo.toml --fix-inplace=false export --path=.

# Format TOML file in-place
format:
    dagger call check --source=. --filename=Cargo.toml --fix-inplace=true export --path=.
    @echo "Formatted file:"
    @cat Cargo.toml

# Run complete demo with setup, check and format
demo: setup check format
