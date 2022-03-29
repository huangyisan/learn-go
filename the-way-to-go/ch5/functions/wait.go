package functions

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err != nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying…", err)
		// 2 4 8 16 ....s
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
