package slides

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	template = `
<!DOCTYPE html>
<html>
<head>
<title>%s</title>
<meta charset="utf-8">
<style>
%s
</style>
</head>
<body>
<textarea id="source">
%s
</textarea>
<script src="https://gnab.github.io/remark/downloads/remark-latest.min.js"></script>
<script> var slideshow = remark.create(); </script>
</body>
</html>
`
	styles = `
@import url(https://fonts.googleapis.com/css?family=Yanone+Kaffeesatz);
@import url(https://fonts.googleapis.com/css?family=Droid+Serif:400,700,400italic);
@import url(https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700,400italic);

body {
    font-family: 'Droid Serif';
}
h1, h2, h3 {
    font-family: 'Yanone Kaffeesatz';
    font-weight: normal;
}
.remark-code, .remark-inline-code {
    font-family: 'Ubuntu Mono';
}
table {
    margin: 25px;
}
table td {
    padding: 5px;
}
table td i {
    color: red;
}
blockquote {
  margin-top: 10px;
  margin-bottom: 10px;
  margin-left: 50px;
  padding-left: 15px;
  border-left: 3px solid #ccc;
} 
`
)

// Handler returns a handler that will respond with the slide deck.
func Handler(title, sourceFile string) http.Handler {
	source, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, template, title, styles, source)
	})
}
