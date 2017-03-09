# trfile
Time rotate files

### Usage
File name format is [Golang time format.](<https://golang.org/src/time/format.go>)

```
package main

import (
	"log"
	"os"

	"github.com/nissy/trfile"
)

var fileNameFormat = "logs/2006/01/02/halo.log" // Golang time format

func main() {
	w, err := trfile.NewFormat(fileNameFormat)

	if err != nil {
		os.Exit(1)
	}

	log.SetOutput(w)
	log.Println("Halo World!") // Now time => logs/2016/09/13/halo.log
}
```