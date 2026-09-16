package site

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type adCleanupTransport struct {
	base http.RoundTripper
}

var adWrapperRE = regexp.MustCompile(`(?is)<div\b[^>]*>\s*(?:<!--.*?-->\s*)*<iframe\b[^>]*(?:title=["']Ads\s+[0-9]+x[0-9]+["']|src=["'][^"']*adtng\.com[^"']*["'])[^>]*>.*?</iframe>\s*(?:<!--.*?-->\s*)*</div>`)

const adCleanupInjection = `<style id="mirror-ad-cleanup-style">
.banner-container:has(iframe[title^="Ads "]),.banner-container:has(iframe[src*="adtng.com"]),div:has(>iframe[title^="Ads "]),div:has(>iframe[src*="adtng.com"]){display:none!important;height:0!important;min-height:0!important;margin:0!important;padding:0!important;overflow:hidden!important}
</style><script id="mirror-ad-cleanup-script">(function(){function pruneAds(){document.querySelectorAll('iframe[title^="Ads "],iframe[src*="adtng.com"]').forEach(function(frame){var parent=frame.parentElement;if(parent&&parent!==document.body)parent.remove();else frame.remove()})}function start(){pruneAds();new MutationObserver(pruneAds).observe(document.documentElement,{subtree:true,childList:true})}if(document.readyState==='loading')document.addEventListener('DOMContentLoaded',start,{once:true});else start()})();</script>`

func init() {
	base := http.DefaultTransport
	http.DefaultTransport = &adCleanupTransport{base: base}
}

func (t *adCleanupTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}
	resp, err := base.RoundTrip(req)
	if err != nil || resp == nil || resp.Body == nil {
		return resp, err
	}
	if !strings.Contains(strings.ToLower(resp.Header.Get("Content-Type")), "text/html") {
		return resp, nil
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return resp, nil
	}
	_ = resp.Body.Close()
	body = adWrapperRE.ReplaceAll(body, nil)
	lower := bytes.ToLower(body)
	if i := bytes.Index(lower, []byte("</head>")); i >= 0 {
		patched := make([]byte, 0, len(body)+len(adCleanupInjection))
		patched = append(patched, body[:i]...)
		patched = append(patched, adCleanupInjection...)
		patched = append(patched, body[i:]...)
		body = patched
	} else {
		body = append([]byte(adCleanupInjection), body...)
	}
	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))
	resp.Header.Set("Content-Length", stringLength(len(body)))
	return resp, nil
}

func stringLength(n int) string {
	if n == 0 {
		return "0"
	}
	buf := [20]byte{}
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
