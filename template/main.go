package main

import (
	"log"
	"os"

	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/builtin"
	"github.com/open2b/scriggo/native"
)

func main() {

	content := []byte(`
<!DOCTYPE html>
<html>
<head>Hello</haed>
<body>
	{% who := "World" %}
	Hello, {{ replace(who, "World", "世界", 1) }}!
</body>
</html>
	`)

	fsys := scriggo.Files{"index.html": content}

	// Allow to use "replace"
	globals := native.Declarations{
		"replace": builtin.Replace,
	}
	opts := scriggo.BuildOptions{Globals: globals}

	template, err := scriggo.BuildTemplate(fsys, "index.html", &opts)
	if err != nil {
		log.Fatal(err)
	}

	err = template.Run(os.Stdout, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

}
