package web

import (
	"testing"

	"github.com/dart998/hentaila-archive-ex4100/internal/database"
	"github.com/dart998/hentaila-archive-ex4100/internal/hentaila"
)

func TestLocalFolderCandidatesKeepsNumberedSequelsSeparate(t *testing.T) {
	item := hentaila.Item{Title: "Onichichi 2"}
	lib := []database.LibraryItem{
		{Name: "Onichichi", Path: "/library/Onichichi", Files: 2},
		{Name: "Onichichi 2", Path: "/library/Onichichi 2", Files: 1},
	}

	got := localFolderCandidates(item, lib)
	if len(got) != 1 {
		t.Fatalf("got %d candidates, want 1: %#v", len(got), got)
	}
	if got[0].Item.Name != "Onichichi 2" {
		t.Fatalf("matched %q, want Onichichi 2", got[0].Item.Name)
	}
}

func TestLocalFolderCandidatesDoesNotUseSequelForBaseTitle(t *testing.T) {
	item := hentaila.Item{Title: "Onichichi"}
	lib := []database.LibraryItem{{Name: "Onichichi 2", Path: "/library/Onichichi 2", Files: 1}}
	if got := localFolderCandidates(item, lib); len(got) != 0 {
		t.Fatalf("matched sequel for base title: %#v", got)
	}
}
