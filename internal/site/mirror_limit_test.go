package site

import "testing"

func TestMediaSlug(t *testing.T) {
	tests := map[string]string{
		"/media/serie-a":   "serie-a",
		"/media/serie-a/1": "serie-a",
		"/catalogo":        "",
	}
	for path, want := range tests {
		got, ok := mediaSlug(path)
		if ok != (want != "") || got != want {
			t.Errorf("mediaSlug(%q) = %q, %v; want %q", path, got, ok, want)
		}
	}
}

func TestAdmitSeriesLimit(t *testing.T) {
	m := &Mirror{seriesLimit: 3}
	selected := map[string]bool{}
	for _, raw := range []string{
		"https://hentaila.com/media/uno",
		"https://hentaila.com/media/dos/1",
		"https://hentaila.com/media/tres",
	} {
		if !m.admitSeries(raw, selected) {
			t.Fatalf("expected %q to be admitted", raw)
		}
	}
	if m.admitSeries("https://hentaila.com/media/cuatro", selected) {
		t.Fatal("fourth series was admitted despite a limit of three")
	}
	if !m.admitSeries("https://hentaila.com/media/dos/2", selected) {
		t.Fatal("another episode from an admitted series was rejected")
	}
}
