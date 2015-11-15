title: Idiomatic not Idiotic CLIs
author:
  name: Austin Riendeau
  twitter: apriendeau
  url: http://www.apriendeau.com
output: presentation.html
theme: apriendeau/simple-theme

--

# Idiomatic not Idiotic CLIs

--

### Mistakes...

we all make em, embrace them, learn from them.

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
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "boom",
		Short: "make an explosive entrance",
		SilenceErrors: true,
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			println("make an explosive entrance")
			return nil
		},
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err) // <- all errors come here
	}
}
```

--

### New Fad. New Problem... microservices.

Where does all this logic live?

--

## Each API comes with its own client lib of course.

Contents:

* Models
* Version
* Library Commands for the API calls
  * Example: a `.Login()` command for our `/login` api

We have a total of 6 services, each now owning their own lib.

--

### Common Printer


