package providers

import "testing"

func TestClassifyHentaiLAProviders(t *testing.T) {
	tests := map[string]string{
		"https://mega.nz/embed/example#key":        "mega",
		"https://www.yourupload.com/embed/example": "yourupload",
		"https://streamwish.to/e/example":          "streamwish",
		"https://www.mp4upload.com/embed-example":  "mp4upload",
		"https://vidhidepro.com/v/example":          "vidhide",
		"https://voe.sx/e/example":                  "voe",
	}
	for raw, want := range tests {
		if got := classify(raw); got != want {
			t.Errorf("classify(%q) = %q; want %q", raw, got, want)
		}
	}
}

func TestDetectUsesConfiguredOrder(t *testing.T) {
	body := `https://voe.sx/e/a https://mega.nz/embed/b#key https://streamwish.to/e/c`
	got := Detect(body, []string{"mega", "streamwish", "voe"})
	if len(got) != 3 {
		t.Fatalf("Detect returned %d sources; want 3", len(got))
	}
	for i, want := range []string{"mega", "streamwish", "voe"} {
		if got[i].Provider != want {
			t.Errorf("source %d provider = %q; want %q", i, got[i].Provider, want)
		}
	}
}
