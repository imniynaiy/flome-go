# Flome

Simple short note app.

## Usage

1. Generate password (go run tool/genpassword.go -p your-password -s salt-as-in-config)
2. Update the username and password in init/database.sql
3. Run the sql script to set up database.
4. Build the [frontend project](https://github.com/theoriz0/flome-react), copy the dist files under 
5. Update configs (default config locates at `configs/config.yml`)
6. Run the server (Or build docker image `make docker`)