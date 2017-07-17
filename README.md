Client for newsapi.org api written in Golang

Getting articles:

```go
package main

import(
	"fmt"
	"github.com/bas24/newsapi"
)

func main(){
	result, err := newsapi.GetArticles("YOUR_API_KEY")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range result{
		fmt.Println(v)
	}
}

```

If you want to get the topic for feed source, just:

```go
	topic := newsapi.GetTopic("bbc-news")
```