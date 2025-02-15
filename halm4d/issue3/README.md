# halm4d's solution to issue3

## Usage

### Run the backend server
Set SERVER_PORT environment variable to the port you want the server to listen to. __Default is__ `8080`.
```shell
$ cd backend
$ go run main.go
```

### Run batch writer terminal application
Set the following environment variables or use the default values:
- `SERVER_URL` - The URL of the backend server. __Default is__ `http://localhost:8080`.
- `BATCH_SIZE` - The number of items to flush in a single batch. __Default is__ `3`.
- `OUTPUT_FILE` - The file to write the output to. __Default is__ `output.txt`.
- `PANIC_RATE` - The rate at which the application should panic. __Default is__ `10`.
```shell
$ cd batch-writer
$ go run main.go
```