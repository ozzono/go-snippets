package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	result := "teste"
	code := `
	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println("It worked")
	}
	`
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://play.golang.org/`),
		chromedp.WaitVisible(`pre`),
		chromedp.SetValue("#code", code, chromedp.ByID),
		chromedp.Click(`#run`, chromedp.NodeVisible, chromedp.ByID),
		chromedp.WaitVisible(`
		#output >
			 pre >
				 span.system`),
		chromedp.Text(`#output`, &result, chromedp.ByID),
		chromedp.Stop(),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result %q", result)
}
