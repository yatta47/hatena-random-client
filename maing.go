package main

import (
    "fmt"
    gofeed "github.com/mmcdole/gofeed"
    _ "reflect"
    "sort"
    "strconv"
    "time"
    "math/rand"
)

type HatenaItem struct {
    Title  string
    Link   string
    Hatebu int
}

type HatenaItems []HatenaItem

func (b HatenaItems) Len() int {
    return len(b)
}

func (b HatenaItems) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

func (b HatenaItems) Less(i, j int) bool {
    return b[i].Hatebu > b[j].Hatebu
}

func main() {
    url := "http://b.hatena.ne.jp/hotentry/it.rss"

    fp := gofeed.NewParser()
    feed, _ := fp.ParseURL(url)

    items := feed.Items

    var hatenaitems HatenaItems

    for _, item := range items {
        extension := item.Extensions
        hatebu := extension["hatena"]["bookmarkcount"]
        var count int
        count, _ = strconv.Atoi(hatebu[0].Value)

        var hatenaitem HatenaItem = HatenaItem{item.Title, item.Link, count}
        hatenaitems = append(hatenaitems, hatenaitem)

    }
    sort.Sort(hatenaitems)

    article := choice(hatenaitems)
    fmt.Printf("%d - %s - %s\n", article.Hatebu, article.Title, article.Link)

}

func choice(s HatenaItems) HatenaItem {
    rand.Seed(time.Now().UnixNano())
    i := rand.Intn(len(s))

    return s[i]
}
