# Ollama Registry Pull Through Cache

The [ollama](https://ollama.com/) registry is somewhat a docker registry, but also somewhat not. 
Hence, normal docker pull through caches do not work. This project aims to pull through cache, that you can 
deploy in your local network, to speed up the pull times of the ollama registry.

## Usage

### Start the cache

#### Golang

1. Install golang
2. Clone the repo
3. Build the code via `go build -o proxy main.go`
4. Execute the binary via `./proxy`
5. The cache will be available at http://localhost:9200`

#### Docker

1. Run via docker `docker run -p 9200:9200 -v ./cache_dir_docker:/pull-through-cache d9abbf62c6ec`
   2. This mounts the local folder `cache_dir_docker` to the cache directory of the container. This will contain ollama files.
5. The cache will be available at http://localhost:9200`

### Use the cache

This is quite easy, just prepend `http://localhost:9200/library/` to the image you want to run/pull

This `ollama pull <image>:<tag>` becomes `ollama pull --insecure http://localhost:9200/library/<image>:<tag>`
**Note: As we run on a non-https endpoint we need to add the `--insecure` to the command**

p.S. Please give a thumbs up on this [PR](https://github.com/ollama/ollama/pull/5241), so that the default behavior of the ollama client can be overwrite to use the cache. 
Will look nicer and work better.

## Contributing

Feel free to create issues and PRs. The project is tiny as of now, so no dedicated guidelines. 

Disclaimer: This is a side project. Don't expect any fast responses on anything. 

## Related ollama issues & PRs

- It is a fix to https://github.com/ollama/ollama/issues/914#issuecomment-1953482174
- To make its behavior work better, we would need this PR merged: https://github.com/ollama/ollama/pull/5241

## License

MIT