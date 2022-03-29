package functions

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func eofError() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}
		fmt.Println(r)
	}

}
