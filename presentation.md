title: Idiomatic not Idiotic CLIs
output: index.html

--

# Idiomatic not Idiotic CLIs

--

### Mistakes...

we all make em, embrace them, learn from them.

![facepalm](https://media.giphy.com/media/p60L4kNY1szqU/giphy.gif)

--

### Major Errors

Most Libraries get this wrong (IMO) because:

1. No good Common Error Handling
2. Error does not propagated

--
What we started with:

```go
package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "boom"
  app.Usage = "make an explosive entrance"
  app.Action = func(c *cli.Context) {
    println("boom! I say!")
  }

  app.Run(os.Args)
}
```

--

What we should do is:

```go
package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "boom",
	Short:         "make an explosive entrance",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(c *cobra.Command, args []string) error {
		// oh no! lets fake a error
		return errors.New("BOOM!! I SAY!")
	},
}

func main() {
	if cmd, err := RootCmd.ExecuteC(); err != nil {
		fmt.Println(cmd.Usage())
		log.Fatal(err) // <- all errors come here
	}
}
```

--

### Note on Cobra

It will print the error and usage by default, to disable this use
these boolean flags:

* SilenceErrors
* SilenceUsage


This also helps with testing because you can call

```go
func TestBoom(t *testing.T) {
	err := RootCmd.RunE()
	if err != nil {
		t.Errorf(err) // boom! easy enough to check
	}
}
```

--

### New Fad. New Problem... microservices.

Where does all this logic live?

#### Each API comes with its own client lib of course.

Contents:

* Models
* Version
* Library Commands for the API calls
  * i.e. a `.Login()` command for our `/login` api

We have a total of 6 services, each now owning their own lib.

--

### Common Printer and Styles:

1. Have your configuration or a package do the formatting.

```go
func scribe(w io.Writer, v interface{}, str string) error {
	// example uses our config package
	switch config.Format {
	case "json"a:
		if err := writeJSON(v); err != nil {
			return err
		}
	case "shell":
		fmt.Fprintln(w, str)
	default:
		return errors.New("Not an acceptable format (json|shell)")
	}
}

func writeJson(v interface{}) error {
	...
}
```

--

### Repos I recommend and suggest

* github.com/spf13/cobra

* github.com/kelseyhightower/envconfig

* github.com/spf13/viper

--

<div class="author">
<span class="name">Austin Riendeau</span><br/>
<div class="social-container"><div class="twitter"><span class="flaticon-twitter"></span></div>
<div class="name">@apriendeau</div>
</div><br/>
<div class="social-container"><div class="github"><span class="flaticon-cat6"></span></div>
<div class="name">github.com/apriendeau</div></div><br/>
</div>
