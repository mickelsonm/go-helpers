Thesaurus Helper
---
This is a pretty basic helper for looking up synonyms for a word. By the API design, it separates words by noun and verbs. For my use case, I needed Combined functionality, so I went ahead and added that in.

Example Usage:
--
```go
package main

import (
	"fmt"

	"github.com/mickelsonm/go-helpers/thesaurus"
)

func main() {
	//key has to be obtained here: https://words.bighugelabs.com/getkey.php
	thesaurus := &thesaurus.BigHugh{APIKey: "<your key>"}

	syns, err := thesaurus.Synonyms("<word>")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(syns.Combined)
}

```