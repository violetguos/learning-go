package main
import ( "bytes"
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

func countWordsAndImages(doc *html.Node) (int, int) { 
	// to do
	

}

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
	fmt.Fprintf(os.Stderr, "parse failed: %s\n", err) os.Exit(-1)
	}
		words, pics := countWordsAndImages(doc)
		fmt.Printf("%d words and %d images\n", words, pics)
}