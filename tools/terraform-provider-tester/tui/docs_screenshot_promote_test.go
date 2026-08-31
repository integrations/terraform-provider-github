package tui

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	errEmptyPNG          = errors.New("png is empty")
	errDecodePNG         = errors.New("decode png")
	errUnexpectedPNGSize = errors.New("unexpected png size")
	errPromotePNG        = errors.New("promote png")
)

func Test_validateAndPromotePNG_rejectsInvalidSources(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		write   func(t *testing.T, path string)
		wantErr error
	}{
		{
			name: "not png",
			write: func(t *testing.T, path string) {
				t.Helper()
				if err := os.WriteFile(path, []byte("not a png"), 0o600); err != nil {
					t.Fatalf("write invalid png: %v", err)
				}
			},
			wantErr: errDecodePNG,
		},
		{
			name: "truncated png",
			write: func(t *testing.T, path string) {
				t.Helper()
				data := mustEncodePNG(t, docsScreenshotWidth, docsScreenshotHeight, color.RGBA{R: 0x22, G: 0x44, B: 0x66, A: 0xff})
				if len(data) < 32 {
					t.Fatalf("encoded png too small: %d", len(data))
				}
				if err := os.WriteFile(path, data[:32], 0o600); err != nil {
					t.Fatalf("write truncated png: %v", err)
				}
			},
			wantErr: errDecodePNG,
		},
		{
			name: "wrong dimensions",
			write: func(t *testing.T, path string) {
				t.Helper()
				mustWritePNG(t, path, 1200, docsScreenshotHeight, color.RGBA{R: 0x88, G: 0x55, B: 0x22, A: 0xff}, 0o600)
			},
			wantErr: errUnexpectedPNGSize,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			scratch := t.TempDir()
			dstPath := filepath.Join(scratch, "existing.png")
			before := mustWritePNG(t, dstPath, docsScreenshotWidth, docsScreenshotHeight, color.RGBA{R: 0xaa, G: 0xbb, B: 0xcc, A: 0xff}, 0o644)
			tmpPath := filepath.Join(scratch, "candidate.png")
			tc.write(t, tmpPath)

			err := validateAndPromotePNG(tmpPath, dstPath, docsScreenshotWidth, docsScreenshotHeight)
			if err == nil {
				t.Fatal("expected validation error")
			}
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error %v does not match %v", err, tc.wantErr)
			}

			after, readErr := os.ReadFile(dstPath)
			if readErr != nil {
				t.Fatalf("read destination after failed promote: %v", readErr)
			}
			if !bytes.Equal(after, before) {
				t.Fatal("destination changed after failed promote")
			}
		})
	}
}

func Test_validateAndPromotePNG_rejectsPostIHDRTruncation(t *testing.T) {
	t.Parallel()

	scratch := t.TempDir()
	dstPath := filepath.Join(scratch, "existing.png")
	before := mustWritePNG(t, dstPath, docsScreenshotWidth, docsScreenshotHeight, color.RGBA{R: 0xaa, G: 0xbb, B: 0xcc, A: 0xff}, 0o644)

	data := mustEncodePNG(t, docsScreenshotWidth, docsScreenshotHeight, color.RGBA{R: 0x22, G: 0x44, B: 0x66, A: 0xff})
	if len(data) <= 33 {
		t.Fatalf("encoded png too small: %d", len(data))
	}
	truncated := data[:33]
	cfg, err := png.DecodeConfig(bytes.NewReader(truncated))
	if err != nil {
		t.Fatalf("DecodeConfig should accept post-IHDR truncation: %v", err)
	}
	if cfg.Width != docsScreenshotWidth || cfg.Height != docsScreenshotHeight {
		t.Fatalf("DecodeConfig dimensions = %dx%d, want %dx%d", cfg.Width, cfg.Height, docsScreenshotWidth, docsScreenshotHeight)
	}

	tmpPath := filepath.Join(scratch, "candidate.png")
	if err := os.WriteFile(tmpPath, truncated, 0o600); err != nil {
		t.Fatalf("write truncated png: %v", err)
	}

	err = validateAndPromotePNG(tmpPath, dstPath, docsScreenshotWidth, docsScreenshotHeight)
	if err == nil {
		t.Fatal("expected validation error for post-IHDR truncation")
	}
	if !errors.Is(err, errDecodePNG) {
		t.Fatalf("error %v does not match %v", err, errDecodePNG)
	}

	after, readErr := os.ReadFile(dstPath)
	if readErr != nil {
		t.Fatalf("read destination after failed promote: %v", readErr)
	}
	if !bytes.Equal(after, before) {
		t.Fatal("destination changed after failed promote")
	}
}

func Test_validateAndPromotePNG_replacesDestinationWhenValid(t *testing.T) {
	t.Parallel()
	if runtime.GOOS == "windows" {
		t.Skip("atomic destination replacement is only exercised on the generator's supported Unix/macOS path")
	}

	scratch := t.TempDir()
	dstPath := filepath.Join(scratch, "existing.png")
	_ = mustWritePNG(t, dstPath, docsScreenshotWidth, docsScreenshotHeight, color.RGBA{R: 0x10, G: 0x20, B: 0x30, A: 0xff}, 0o644)
	tmpPath := filepath.Join(scratch, "candidate.png")
	want := mustWritePNG(t, tmpPath, docsScreenshotWidth, docsScreenshotHeight, color.RGBA{R: 0xde, G: 0xad, B: 0xbe, A: 0xff}, 0o600)

	if err := validateAndPromotePNG(tmpPath, dstPath, docsScreenshotWidth, docsScreenshotHeight); err != nil {
		t.Fatalf("validateAndPromotePNG: %v", err)
	}
	got, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("read promoted destination: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatal("destination contents do not match promoted png")
	}
	if _, err := os.Stat(tmpPath); !os.IsNotExist(err) {
		t.Fatalf("temporary png should be gone after promotion, stat err=%v", err)
	}
	info, err := os.Stat(dstPath)
	if err != nil {
		t.Fatalf("stat promoted destination: %v", err)
	}
	if info.Mode().Perm() != 0o644 {
		t.Fatalf("promoted destination mode = %o, want 644", info.Mode().Perm())
	}
}

func validateAndPromotePNG(srcPath, dstPath string, wantWidth, wantHeight int) error {
	info, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("stat png %s: %w", srcPath, err)
	}
	if info.Size() == 0 {
		return fmt.Errorf("%w: %s", errEmptyPNG, srcPath)
	}
	f, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("open png %s: %w", srcPath, err)
	}

	img, err := png.Decode(f)
	closeErr := f.Close()
	if err != nil {
		decodeErr := fmt.Errorf("%w %s: %w", errDecodePNG, srcPath, err)
		if closeErr != nil {
			return errors.Join(decodeErr, fmt.Errorf("close png %s: %w", srcPath, closeErr))
		}
		return decodeErr
	}
	if closeErr != nil {
		return fmt.Errorf("close png %s: %w", srcPath, closeErr)
	}

	bounds := img.Bounds()
	if bounds.Dx() < wantWidth || bounds.Dy() < wantHeight {
		return fmt.Errorf("%w for %s: got %dx%d want at least %dx%d", errUnexpectedPNGSize, srcPath, bounds.Dx(), bounds.Dy(), wantWidth, wantHeight)
	}
	if err := os.Chmod(srcPath, 0o644); err != nil {
		return fmt.Errorf("chmod png %s: %w", srcPath, err)
	}
	if err := os.Rename(srcPath, dstPath); err != nil {
		return fmt.Errorf("%w %s -> %s: %v", errPromotePNG, srcPath, dstPath, err)
	}
	return nil
}

func mustWritePNG(t *testing.T, path string, width, height int, fill color.Color, mode os.FileMode) []byte {
	t.Helper()
	data := mustEncodePNG(t, width, height, fill)
	if err := os.WriteFile(path, data, mode); err != nil {
		t.Fatalf("write png %s: %v", path, err)
	}
	return data
}

func mustEncodePNG(t *testing.T, width, height int, fill color.Color) []byte {
	t.Helper()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, fill)
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatalf("encode png: %v", err)
	}
	return buf.Bytes()
}
