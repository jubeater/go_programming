// Fetchall fetches URLs in parallel and reports their times and sizes.
package main
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch, 
		//it can also make sure we will output all the result from our goroutine
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url,"http://") {
			url = strings.Join([]string{"http://",url},"")
	}
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

//from the official definition of the method http.Get() and goroutine, 
//i guess if have no response ,
//it will output other result first, even maybe the no response will be time out, then output the err info,
//and only after get the all the goroutine info back, the main program would stop stuck at the 
//fmt.Println(<-ch), without this line, our program will exit with no info fetched
