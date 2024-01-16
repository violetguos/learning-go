package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"golang.org/x/net/html"
)

var raw = ` <!DOCTYPE html> <html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML images are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
</body> </html>`

func helper(curr *html.Node, numWords int, numImg int) (int, int) {
    
    if curr.Type == html.ElementNode && curr.Data == "img" {
        for _, a := range curr.Attr {
            if a.Key == "src" {
                numImg++
                
            }
        }
    } else if curr.Type == html.TextNode {
            
            numWords+=len(strings.Fields(curr.Data))
        
    }
    for next := curr.FirstChild; next != nil; next = next.NextSibling {
        fmt.Printf("curr %s, next %s\n", curr.Data, next.Data)
        numWords, numImg = helper(next, numWords, numImg)
    }
    return numWords, numImg

}

func countWordsAndImages(doc *html.Node) (int, int) {
	
    // count all nodes in the current level
    
	
	return helper(doc, 0, 0)
}

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}
	words, pics := countWordsAndImages(doc)
	fmt.Printf("%d words and %d images\n", words, pics)
}
