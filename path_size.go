package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !info.IsDir() {
		size := info.Size()
		return FormatSize(size, human), nil
	}

	var total int64
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		name := entry.Name()
		if !all && strings.HasPrefix(name, ".") {
			continue
		}

		entryPath := filepath.Join(path, name)
		if entry.IsDir() && recursive {
			subSizeStr, err := GetPathSize(entryPath, recursive, false, all)
			if err != nil {
				return "", err
			}
			var subSize int64
			fmt.Sscanf(subSizeStr, "%dB", &subSize)
			total += subSize
		} else if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return "", err
			}
			total += info.Size()
		}
	}

	return FormatSize(total, human), nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	s := float64(size)
	i := 0
	for s >= 1024 && i < len(units)-1 {
		s /= 1024
		i++
	}

	return fmt.Sprintf("%.1f%s", s, units[i])
}
