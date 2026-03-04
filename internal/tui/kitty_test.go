package tui

import (
	"strings"
	"testing"
)

func TestKittyUploadAndPlaceSmall(t *testing.T) {
	result := kittyUploadAndPlace([]byte("hello"), 1, 10, 5)
	// Should contain upload sequence
	if !strings.Contains(result, "\033_Ga=t,") {
		t.Error("should contain upload APC sequence")
	}
	// Should contain placement sequence
	if !strings.Contains(result, "\033_Ga=p,U=1,i=1,c=10,r=5") {
		t.Error("should contain virtual placement sequence")
	}
}

func TestKittyUploadAndPlaceEmpty(t *testing.T) {
	result := kittyUploadAndPlace(nil, 1, 10, 5)
	if result != "" {
		t.Errorf("empty data should return empty string, got %q", result)
	}
}

func TestKittyUploadAndPlaceLarge(t *testing.T) {
	data := make([]byte, 4000)
	for i := range data {
		data[i] = byte(i % 256)
	}
	result := kittyUploadAndPlace(data, 1, 10, 5)
	// Should have continuation chunks (m=1)
	if !strings.Contains(result, "m=1") {
		t.Error("large data should have continuation chunks with m=1")
	}
	// Should end with placement
	if !strings.Contains(result, "a=p,U=1") {
		t.Error("should end with virtual placement")
	}
}

func TestRenderImagePlaceholder(t *testing.T) {
	result := renderImagePlaceholder(1, 3, 2)
	// Should contain the placeholder character
	if !strings.ContainsRune(result, placeholder) {
		t.Error("should contain U+10EEEE placeholder character")
	}
	// Should contain foreground color sequence (ID=1 → 0,0,1)
	if !strings.Contains(result, "\033[38;2;0;0;1m") {
		t.Error("should contain foreground color encoding image ID")
	}
	// Should contain reset
	if !strings.Contains(result, "\033[39m") {
		t.Error("should contain foreground color reset")
	}
	// Should have a newline between rows
	if !strings.Contains(result, "\n") {
		t.Error("should have newline between rows")
	}
}

func TestRenderImagePlaceholderZero(t *testing.T) {
	result := renderImagePlaceholder(1, 0, 0)
	if result != "" {
		t.Errorf("zero dimensions should return empty string, got %q", result)
	}
}

func TestImageDimensionsFallback(t *testing.T) {
	// Invalid image data should use fallback dimensions
	cols, rows := imageDimensions([]byte("not an image"), 80)
	if cols != 40 || rows != 10 {
		t.Errorf("fallback dimensions = (%d, %d), want (40, 10)", cols, rows)
	}
}
