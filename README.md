# Ollama Registry Pull Through Cache

The [ollama](https://ollama.com/) registry is somewhat a docker registry, but also somewhat not. 
Hence, normal docker pull through caches do not work. This project aims to pull through cache, that you can 
deploy in your local network, to speed up the pull times of the ollama registry.

[![Demo video of the tool on YouTube](https://img.youtube.com/vi/7_nlWlhrqNw/0.jpg)](https://www.youtube.com/watch?v=7_nlWlhrqNw)

## Usage

### 1. Start the cache

#### Golang

1. Install golang
2. Clone the repo
3. Build the code via `go build -o proxy main.go`
4. Execute the binary via `./proxy`
5. The cache will be available at http://localhost:9200`

#### Docker

1. Run via docker 
    ```
    docker run -p 9200:9200 -v ./cache_dir_docker:/pull-through-cache ghcr.io/beans-bytes/ollama-registry-pull-through-proxy:latest
    ```
2. This mounts the local folder `cache_dir_docker` to the cache directory of the container. This will contain ollama files.
3. The cache will be available at http://localhost:9200`

### 2. Use the cache

This is quite easy, **just prepend `http://localhost:9200/library/` to the image you want to run/pull**

This `ollama pull <image>:<tag>` becomes 
```bash
ollama pull --insecure http://localhost:9200/library/<image>:<tag>
```

**Note: As we run on a non-https endpoint we need to add the `--insecure` to the command**

p.S. Please give a thumbs up on this [PR](https://github.com/ollama/ollama/pull/5241), so that the default behavior of the ollama client can be overwrite to use the cache. 
Will look nicer and work better.

## Architecture

This proxy is based on a worker architecture. It has a worker pool, that can be configured via the `NUM_DOWNLOAD_WORKERS` environment variable. 
When a request comes in, the proxy will check if the file is already in the cache. If not, it will mark it as queued and serve the request from the upstream. 
In the background the worker checks the queue and downloads the file. Going forward, the file will be served from the local cache, not upstream anymore

## Configuration options (via environment variables)

| Environment Variable | Description                                                                                                      | Default                       |
|----------------------|------------------------------------------------------------------------------------------------------------------|-------------------------------|
| `PORT` | The port the proxy listens on                                                                                    | `9200`                        |
| `DUMP_UPSTREAM_REQUESTS` | If the proxy should dump the upstream requests                                                                   | `false`                       |
| `CACHE_DIR` | Directory where the cache is stored                                                                              | `./cache_dir`                 |
| `NUM_DOWNLOAD_WORKERS` | Number of workers that download the files                                                                        | 1                             |
| `MANIFEST_LIFETIME` | The lifetime of the model manifest. These change from time to time on the registry, and we want to get updates   | `240h` / 10 days              |
| `UPSTREAM_ADDRESS` | The upstream ollama registry                                                                                     | `https://registry.ollama.ai/` |
| `LOG_LEVEL` | The log level of the proxy                                                                                       | `info`                        |
| `LOG_FORMAT_JSON` | If the log should be in json format. This is nice for external log tools. (e.g. when you run this in kubernetes) | `false`                       |

## Contributing

Feel free to create issues and PRs. The project is tiny as of now, so no dedicated guidelines. 

Disclaimer: This is a side project. Don't expect any fast responses on anything. 

## Related ollama issues & PRs

- It is a fix to https://github.com/ollama/ollama/issues/914#issuecomment-1953482174
- To make its behavior work better, we would need this PR merged: https://github.com/ollama/ollama/pull/5241

## License

MIT