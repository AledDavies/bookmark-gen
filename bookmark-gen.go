package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	info, err := os.Stdin.Stat()

	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("ERROR: The command is intended to work with pipes.")
		fmt.Println("Usage: cat links.txt | bookmark-gen")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	count := 0

	// There is probably a cleaner way to do this with templates...
	fmt.Println("<!DOCTYPE NETSCAPE-Bookmark-file-1>")
	fmt.Println("<HTML>")
	fmt.Println("<META HTTP-EQUIV=\"Content-Type\" CONTENT=\"text/html; charset=UTF-8\">")
	fmt.Println("<Title>Bookmarks</Title>")
	fmt.Println("<H1>Bookmarks</H1>")
	fmt.Println("<DT>")
	fmt.Println("<H3 FOLDED>New Bookmarks</H3>")
	fmt.Println("<DL>")

	for {
		input, err := reader.ReadString('\n')
		count++

		// Only print lines with at least 1 char
		if len(input) > 0 {
			fmt.Println(formatLine(input, count))
		}

		if err != nil && err == io.EOF {
			break
		}
	}

	fmt.Println("</DL>")
	fmt.Println("</DT>")
	fmt.Println("</HTML>")
}

// Link line template
const linkTemplate = "<DT><A HREF=\"%s\">Link #%d</DT>" // Do not attempt to balance the <A> tag

// Formats the link into the expected line format
func formatLine(input string, count int) string {
	return fmt.Sprintf(linkTemplate, strings.TrimSuffix(input, "\n"), count)
}
