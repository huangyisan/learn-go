package main

import "learn/functions"

func main() {
	functions.CallchangeSlice()
}

SELECT table_schema, ROUND(SUM(data_length+index_length)/1024/1024/1024,2) "size in GB" FROM information_schema.