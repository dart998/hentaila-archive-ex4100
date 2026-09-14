package web

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/dart998/hentaila-archive-ex4100/internal/config"
	"github.com/dart998/hentaila-archive-ex4100/internal/crawler"
	"github.com/dart998/hentaila-archive-ex4100/internal/database"
)

func TestDownloadItemIndexesSeriesOutsidePersonalLibrary(t *testing.T) {
	const slug = "ojisan-de-umeru-ana-the-animation"
	const cookie = "session=test-cookie"

	var receivedCookie string
	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/media/"+slug {
			http.NotFound(w, r)
			return
		}
		receivedCookie = r.Header.Get("Cookie")
		fmt.Fprintf(w, `<html><body><h1>Ojisan de Umeru Ana The Animation</h1><a href="/media/%s/1">1</a><a href="/media/%s/2">2</a></body></html>`, slug, slug)
	}))
	defer origin.Close()

	db, err := database.Open(filepath.Join(t.TempDir(), "archive.sqlite"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err = db.SetSetting("hentaila_session_cookie", cookie); err != nil {
		t.Fatal(err)
	}

	metadataDir := t.TempDir()
	crawl := crawler.New(config.Config{BaseURL: origin.URL, MetadataDir: metadataDir}, db)
	s := &Server{db: db, crawl: crawl, baseURL: origin.URL, libraryRoot: t.TempDir()}

	item, total, err := s.downloadItem(context.Background(), slug)
	if err != nil {
		t.Fatalf("downloadItem: %v", err)
	}
	if item.Slug != slug {
		t.Fatalf("slug = %q, want %q", item.Slug, slug)
	}
	if receivedCookie != cookie {
		t.Fatalf("cookie = %q, want %q", receivedCookie, cookie)
	}
	if total != 2 {
		t.Fatalf("total = %d, want 2", total)
	}
	if got := db.SeriesEpisodeCount(slug); got != 2 {
		t.Fatalf("SeriesEpisodeCount = %d, want 2", got)
	}
}
