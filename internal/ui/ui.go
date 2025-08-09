package ui

import (
	"fmt"
	"time"
)

// cacheBuster stores the current time as a cache buster for static files.
var cacheBuster = fmt.Sprint(time.Now().Unix())

// PublicFile generates a relative URL to a public file. Matches the router's /public static mount.
func PublicFile(filepath string) string {
	return fmt.Sprintf("/%s/%s", "public", filepath)
}

// StaticFile generates a relative URL to a static file including a cache-buster query parameter.
// In this project, static assets are served under /public.
func StaticFile(r *Request, filepath string) string {

	return fmt.Sprintf("/%s/%s?v=%s", r.Config.Static.PublicDir, filepath, cacheBuster)
}
