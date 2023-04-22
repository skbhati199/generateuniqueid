# GENERATE UniqueId


```shell
go mod init github.com/my/repo
```

Then install generateuniqueid helo redis:

```shell
go get github.com/skbhati199/generateuniqueid
```

## Quickstart

```go

import (
	"fmt"
	generatePkg "github.com/skbhati199/generateUnqiueId/pkg"
)

func main() {
	generateId, err := generatePkg.NewGeneratePkg()
	if err != nil {
		_ = fmt.Errorf("fail to generate id %v", err)
	}

	fmt.Printf("Transaction Id : %v", generateId)
}
```