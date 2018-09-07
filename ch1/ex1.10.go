// Fetchall fetches URLs in parallel and reports their times and sizes.
package main
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strconv"
)
func main() {
	start := time.Now()
	f, _ := os.Create("record")
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		tmp := <- ch
		tmp = tmp + "\n"
		f.WriteString(tmp)
		fmt.Println(tmp) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	f.WriteString(strconv.FormatFloat(time.Since(start).Seconds(),'f', 2, 64) + "s elapsed\n")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}