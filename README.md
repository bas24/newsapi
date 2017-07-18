Client for <a href="https://www.newsapi.org">newsapi.org</a> api written in Golang.

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
package main

import(
	"fmt"
	"github.com/bas24/newsapi"
)

func main(){
	topic := newsapi.GetTopic("bbc-news")
	fmt.Println(topic)
}
```
