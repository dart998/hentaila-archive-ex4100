package hentaila

import "testing"

func TestParseLibraryEntries(t *testing.T) {
	body := `<script>const data={libraryEntries:[{mediaId:42,status:0,episode:2,score:8,favorite:true,media:{title:"Serie piloto",episodesCount:4,slug:"serie-piloto",aka:{en:"Pilot Series"}}}]};</script>`
	items, err := parseLibraryEntries(body)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 {
		t.Fatalf("parsed %d items; want 1", len(items))
	}
	got := items[0]
	if got.MediaID != "42" || got.Slug != "serie-piloto" || got.Title != "Serie piloto" || got.Total != 4 || got.Seen != 2 || !got.Favorite {
		t.Fatalf("unexpected item: %+v", got)
	}
	if got.Aliases["en"] != "Pilot Series" {
		t.Fatalf("unexpected aliases: %+v", got.Aliases)
	}
}
