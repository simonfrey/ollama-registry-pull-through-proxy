package invalidate_manifests_worker

import (
	"github.com/rs/zerolog/log"
	"ollama-registry-pull-through-cache/internal/worker/cache_worker"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Run(cacheDir string, maxAge time.Duration) {
	for {
		// Wall all files recursively in cacheDir
		err := filepath.Walk(cacheDir, func(fPath string, info os.FileInfo, err error) error {
			// Check if file is a manifest file
			if !strings.Contains(fPath, "/manifests/") || filepath.Ext(fPath) == ".todo" {
				return nil
			}

			if info.ModTime().After(time.Now().Add(-maxAge)) {
				// File is newer than maxAge. Skip it
				return nil
			}

			// File is older than maxAge. Mark it as .todo file
			log.Info().Str("component", "manifest-invalidation-worker").Msgf("Invalidate manifest %q", fPath)

			cache_worker.QueueFileForDownload(fPath)
			return nil
		})
		if err != nil {
			log.Error().Str("component", "manifest-invalidation-worker").Err(err).Msg("Failed to walk cache dir")
		}

		time.Sleep(10 * time.Second)
	}
}
