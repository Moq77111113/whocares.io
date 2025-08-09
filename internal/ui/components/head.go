package components

import (
	"fmt"
	"strings"

	"github.com/moq77111113/whocares/internal/ui"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// JS returns script tags for HTMX and Alpine.js (deferred), similar to Pagoda.
func JS() g.Node {
	return g.Group{
		h.Script(h.Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js"), h.Defer()),
		h.Script(h.Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"), h.Defer()),
		h.Script(g.Raw(counterJs))}
}

func CSS(r *ui.Request) g.Node {
	return h.Link(

		h.Href(ui.StaticFile(r, "css/main.css")),
		h.Rel("stylesheet"),
		h.Type("text/css"),
	)
}

// Meta renders common meta tags and dynamic OG/Twitter tags from the Request.
func Meta(r *ui.Request) g.Node {

	fmt.Println(r.OpenGraph.Title, r.OpenGraph.Description, r.OpenGraph.ImageURL)
	return g.Group{
		h.Meta(h.Charset("utf-8")),
		h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
		h.TitleEl(g.Text(r.Title)),
		g.If(r.Metatags.Description != "", h.Meta(h.Name("description"), h.Content(r.Metatags.Description))),
		g.If(len(r.Metatags.Keywords) > 0, h.Meta(h.Name("keywords"), h.Content(strings.Join(r.Metatags.Keywords, ", ")))),

		// OpenGraph
		g.If(r.OpenGraph.Title != "", h.Meta(g.Attr("property", "og:title"), h.Content(r.OpenGraph.Title))),
		g.If(r.OpenGraph.Description != "", h.Meta(g.Attr("property", "og:description"), h.Content(r.OpenGraph.Description))),
		h.Meta(g.Attr("property", "og:type"), h.Content("website")),
		g.If(r.OpenGraph.ImageURL != "", h.Meta(g.Attr("property", "og:image"), h.Content(r.OpenGraph.ImageURL))),

		// Twitter
		h.Meta(h.Name("twitter:card"), h.Content("summary_large_image")),
		g.If(r.OpenGraph.Title != "", h.Meta(h.Name("twitter:title"), h.Content(r.OpenGraph.Title))),
		g.If(r.OpenGraph.Description != "", h.Meta(h.Name("twitter:description"), h.Content(r.OpenGraph.Description))),
		g.If(r.OpenGraph.ImageURL != "", h.Meta(h.Name("twitter:image"), h.Content(r.OpenGraph.ImageURL))),
	}
}
