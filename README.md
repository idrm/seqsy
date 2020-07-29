# Seqsy

A simple sequence generator implemented as an HTTP server that returns a monotonically increasing sequence of positive integers.

The HTTP server runs on port 8080, and the generated numbers can be retrieved from the root path (e.g. `curl http://localhost:8080/`)

The sequence is guaranteed to survive restarts by occasionally persisting the latest number to a file; that file is read from whenever the server is restarted. It is located at `/data/counter.txt` and should be mounted using a persistent volume.

Only a single instance of the server per logical sequence should be running at a time, and it shouldn't share the persistence file with the rest.

A basic health-check endpoint is available at `/how/you/doing` to see if the server still responds to HTTP requests (e.g. `curl http://localhost:8080/how/you/doing`).

The server is implemented in Go, which yields a minimal memory (and Docker image) footprint.

# Usage

- Create a Docker volume: `docker volume create seqsy`
- Run Docker image: `docker run -d --name seqsy -v seqsy:/data -P 8080:8080 idrm/seqsy:1.0.0`
- Get a sequence number: `curl http://localhost:8080/`

# Building

Run `docker build -t idrm/seqsy:1.0.0 .`

# License

Apache 2.0

# Miscellaneous

Yes, it's pronounced as you thought -- "six-eye".