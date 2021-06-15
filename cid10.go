package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"unicode"

	"github.com/ronoaldo/bot"
)

type cidChapter struct {
	url   string
	title string
}

type crawl struct {
	bot *bot.Bot
}

func main() {
	flow := newCrawl()
	flow.bot.BaseURL("https://git-scm.com/book/en/v2/GitHub-Contributing-to-a-Project")
	// cid10MainUrl := "/wiki/Classifica%C3%A7%C3%A3o_Estat%C3%ADstica_Internacional_de_Doen%C3%A7as_e_Problemas_Relacionados_com_a_Sa%C3%BAde?oldformat=true"
	page, err := flow.bot.GET("")
	if err != nil {
		log.Printf(".GET err: %v", err)
		return
	}
	fmt.Println(page)
	// tables, err := page.Tables()
	// if err != nil {
	// 	log.Printf("tables() err: %v", err)
	// 	return
	// }
}

func WriteToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func myregexp(exp, text string) []string {
	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(text)
	if len(match) < 1 {
		fmt.Printf("Unable to find match for exp %s\n", exp)
		return []string{}
	}
	return match
}

func cidURLList(rawdata [][]string) []cidChapter {
	output := []cidChapter{}
	for i := range rawdata {
		output = append(output, cidChapter{
			myregexp("href=\"(.*)\".title", fmt.Sprintf("%v", rawdata[i][1]))[1],
			rawdata[i][2]},
		)
	}
	return output
}

func (crawl *crawl) chapters(cidChapter []cidChapter) {
	for i := range cidChapter {
		crawl.chapter(cidChapter[i])
		break
	}
}

func (crawl *crawl) chapter(chapter cidChapter) error {
	fmt.Println(chapter.url)
	page, err := crawl.bot.GET(chapter.url)
	if err != nil {
		return fmt.Errorf("bot.GET err: %v", err)
	}
	tables, err := page.Tables()
	if err != nil {
		return fmt.Errorf("page.Tables err: %v", err)
	}
	for i := range tables {
		fmt.Println(tables[i])
		waitEnter()
	}
	return nil
}

func newCrawl() crawl {
	return crawl{bot.CustomNew(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}})}
}

func waitEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
