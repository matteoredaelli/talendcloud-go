# talendcloud-go
Minimal talendcloud API wrapper
=======

## Usage

```go
package main

import (
	"fmt"
	"github.com/matteoredaelli/talendcloud-go"
)

func main() {
	client := talendcloud.NewClient("https://api.eu.cloud.talend.com/tmc/v2.6", "MY_SILLY_TALEND_TOKEN")
	conn := &client

	message, err := conn.Get("runtimes/remote-engine-clusters?_s=workspace.environment.name==PRD", nil)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)
}

```
