package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/bmaupin/go-epub"
)

type Section struct {
	Title      string
	Paragraphs []string
}

func main() {
	file, err := os.Open("./81439.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	str := string(all)

	chapters := strings.Split(str, "------------")

	e := epub.NewEpub("从红月开始")
	e.SetAuthor("黑山老鬼")
	imgPath, _ := e.AddImage("./from.png", "")
	e.SetCover(imgPath, "")

	for _, v := range chapters {
		sec := splitParagraphs(v)
		content := makeSectionContent(sec)
		e.AddSection(content, sec.Title, "", "")
	}

	e.Write("out.epub")
}

func splitParagraphs(content string) Section {
	lines := strings.Split(content, "\n")

	var section Section

	section.Title = lines[0]

	i := 0
	for _, v := range lines {
		trimmed := strings.TrimSpace(v)
		if len(trimmed) > 0 {
			if i == 0 {
				section.Title = trimmed
			} else {
				section.Paragraphs = append(section.Paragraphs, trimmed)
			}
			i += 1
		}
	}
	return section
}

func makeSectionContent(sec Section) string {
	content := `<h3>` + sec.Title + `</h3>`
	for _, v := range sec.Paragraphs {
		content = content + `<p>` + v + `</p>`
	}

	return content
}
