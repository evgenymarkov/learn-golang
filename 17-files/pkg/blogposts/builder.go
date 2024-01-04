package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleField       = "Title"
	descriptionField = "Description"
	tagsField        = "Tags"
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	post := Post{
		Title:       readMetaField(scanner, titleField),
		Description: readMetaField(scanner, descriptionField),
		Tags:        strings.Split(readMetaField(scanner, tagsField), ", "),
		Body:        readBody(scanner),
	}

	return post, nil
}

func readMetaField(scanner *bufio.Scanner, field string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), field+": ")
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore separator

	buffer := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}

	return strings.TrimSpace(buffer.String())
}
