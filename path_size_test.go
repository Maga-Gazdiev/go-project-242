package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "file.txt")
	err := os.WriteFile(file, []byte("1234"), 0644) // 4B
	require.NoError(t, err)

	size, err := GetPathSize(file, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "4B", size)
}

func TestGetPathSize_Dir(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("1234"), 0644) // 4B
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("1234567"), 0644) // 7B
	require.NoError(t, err)

	size, err := GetPathSize(tmpDir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "11B", size)
}

func TestGetPathSize_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "dir1")
	err := os.Mkdir(subDir, 0755)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(subDir, "a.txt"), []byte("123"), 0644) // 3B
	require.NoError(t, err)

	size, err := GetPathSize(subDir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "3B", size)
}

func TestGetPathSize_HiddenFiles(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.WriteFile(filepath.Join(tmpDir, "a.txt"), []byte("123"), 0644)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmpDir, ".hidden.txt"), []byte("4567"), 0644)
	require.NoError(t, err)

	size, err := GetPathSize(tmpDir, false, false, false)
	require.NoError(t, err)
	require.Equal(t, "3B", size)

	sizeAll, err := GetPathSize(tmpDir, false, false, true)
	require.NoError(t, err)
	require.Equal(t, "7B", sizeAll)
}

func TestGetPathSize_HumanReadable(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("1234"), 0644) // 4B
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("1234567"), 0644) // 7B
	require.NoError(t, err)

	size, err := GetPathSize(tmpDir, false, true, false)
	require.NoError(t, err)
	require.Equal(t, "11.0B", size)
}

func TestGetPathSize_Recursive(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("123"), 0644)
	require.NoError(t, err)
	subDir := filepath.Join(tmpDir, "dir1")
	err = os.Mkdir(subDir, 0755)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(subDir, "file2.txt"), []byte("1234"), 0644)
	require.NoError(t, err)

	size, err := GetPathSize(tmpDir, true, false, false)
	require.NoError(t, err)
	require.Equal(t, "7B", size)
}

func TestFormatSize(t *testing.T) {
	require.Equal(t, "123B", FormatSize(123, false))
	require.Equal(t, "1.2KB", FormatSize(1234, true))
	require.Equal(t, "24.0MB", FormatSize(25165824, true))
	require.Equal(t, "1.0GB", FormatSize(1073741824, true))
}
