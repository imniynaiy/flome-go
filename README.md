# Flome

Simple short note app.

## Usage

1. Generate password (go run tool/genpassword.go -p your-password -s salt-as-in-config)
2. Update the username and password in init/database.sql, run it to set up database.
3. Build the [frontend project](https://github.com/theoriz0/flome-react)
4. Update configs (default config locates at `configs/config.yml`)
5. Run the server