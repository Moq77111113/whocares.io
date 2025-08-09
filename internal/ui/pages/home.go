package pages

import (
	"fmt"

	"github.com/moq77111113/whocares/internal/ui"
	"github.com/moq77111113/whocares/internal/ui/layouts"
	"github.com/moq77111113/whocares/internal/ui/models"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Home(r *ui.Request, c *models.Counter) Node {

	r.Title = r.Config.Base.Title
	r.Metatags.Description = r.Config.Base.Description
	r.OpenGraph.Title = fmt.Sprintf("%s %s - %s", c.Count, c.Message, r.Config.Base.Title)
	r.OpenGraph.Description = c.Subtext
	r.OpenGraph.ImageURL = c.OgImageURL

	return layouts.BaseLayout(r,
		Main(
			Class("min-h-svh flex items-center justify-center px-8 py-16"),
			Div(
				ID("counter-container"),
				Class("text-center w-full max-w-4xl lg:max-w-full mx-auto"),
				Attr("hx-get", "/counter"),
				Attr("hx-trigger", "every 8s"),
				Attr("hx-swap", "innerHTML"),
				CounterContent(c.Count, c.Message, c.Subtext),
			),
		),
	)
}

func CounterContent(count, message, subtext string) Node {
	return Div(
		Class("space-y-8 sm:space-y-12"),

		Div(
			Class("space-y-6"),
			H2(
				Class("text-6xl sm:text-4xl lg:text-[12rem] xl:text-[14rem] font-black tracking-tighter leading-none"),
				Attr("x-data", "counterAnim($el.dataset.count)"),
				Attr("data-count", count),
				Attr("x-init", "start()"),
				Attr("x-text", "display"),
				Text(count),
			),
		),
		Div(
			Class("space-y-4 max-w-3xl mx-auto"),
			H3(
				Class("text-2xl sm:text-3xl lg:text-4xl font-bold leading-tight"),
				Text(message),
			),
		),
		Div(
			Class("max-w-2xl mx-auto"),
			P(
				Class("text-base sm:text-lg opacity-70 leading-relaxed"),
				Text(subtext),
			),
		),
	)
}
