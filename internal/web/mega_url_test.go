package web

import "testing"

func TestMegaHandleAndKeyFormats(t *testing.T) {
	const key = "aqdlvbSxZHOAfqyvCQsg_JRuZ6MyTvweYHNjAdDFOy8"
	cases := []struct {
		name string
		raw  string
	}{
		{name: "modern", raw: "https://mega.nz/file/CYkU3SrJ#" + key},
		{name: "legacy", raw: "https://mega.nz/file/!CYkU3SrJ!" + key},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			handle, gotKey, err := megaHandleAndKey(tc.raw)
			if err != nil {
				t.Fatalf("megaHandleAndKey: %v", err)
			}
			if handle != "CYkU3SrJ" {
				t.Fatalf("handle = %q, want CYkU3SrJ", handle)
			}
			if gotKey != key {
				t.Fatalf("key mismatch")
			}
		})
	}
}

func TestRedactMegaURLLegacy(t *testing.T) {
	const raw = "https://mega.nz/file/!CYkU3SrJ!aqdlvbSxZHOAfqyvCQsg_JRuZ6MyTvweYHNjAdDFOy8"
	got := redactMegaURL(raw)
	if got == raw {
		t.Fatal("legacy Mega key was not redacted")
	}
	if want := "https://mega.nz/file/!CYkU3SrJ!<redacted:43>"; got != want {
		t.Fatalf("redacted URL = %q, want %q", got, want)
	}
}
