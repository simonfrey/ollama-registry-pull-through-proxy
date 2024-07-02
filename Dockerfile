#
# Builder
#

FROM golang:1.22-alpine AS builder

# Create a workspace for the app
WORKDIR /app

# Copy over the files
COPY . ./

# Build
RUN go build -o proxy

#
# Runner
#

FROM alpine AS runner

WORKDIR /


# Copy from builder the final binary
COPY --from=builder /app/proxy /proxy

ENV NUM_DOWNLOAD_WORKERS=1
ENV MANIFEST_LIFETIME="240h"
ENV LOG_FORMAT_JSON="true"
ENV CACHE_DIR="/pull-through-cache"

ENV PORT=9200
EXPOSE 9200

ENTRYPOINT ["/proxy"]
