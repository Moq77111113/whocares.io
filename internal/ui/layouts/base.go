package layouts

import (
	"github.com/moq77111113/whocares/internal/ui"
	"github.com/moq77111113/whocares/internal/ui/components"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// BaseLayout renders the base HTML layout using a Request, like Pagoda.
func BaseLayout(r *ui.Request, content g.Node) g.Node {
	return h.Doctype(
		h.HTML(
			h.Lang("en"),
			h.Head(
				components.Meta(r),
				components.CSS(r),
				components.JS(),
			),
			h.Body(
				h.Class("min-h-svh bg-main text-text font-mono cursor-pointer"),
				g.Attr("onclick", "document.documentElement.classList.toggle('dark')"),

				h.Div(
					h.Class("absolute top-6 left-6 z-10"),
					h.P(
						h.Class("text-sm font-medium text-muted"),
						g.Text(r.Config.Base.Title),
					),
				),
				content,

				h.Div(
					h.Class("absolute bottom-6 right-6 z-10"),
					h.P(
						h.Class("text-xs text-accent opacity-40 italic hover:opacity-70 transition-opacity"),
						g.Text("this wasn't requested either"),
					),
				),
			),
		),
	)
}
