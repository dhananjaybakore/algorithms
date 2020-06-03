package main

import (
	"log"
	. "os"
)

func main() {

	f, err := OpenFile("/Users/dhananjaybakore/work/sftpdownload/AU_VSI_cross_ref_file_08_05142020.csv", O_APPEND|O_CREATE|O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 100000; i++ {
		if _, err := f.Write([]byte("491678xxxxxx6058,2002,491678xxxxxx6058,2502,200000000000000174,200000000000000174,E,VIC,ORDER\n")); err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write([]byte("411111XXXXXX1111,2002,545454XXXXXX5454,2502,200000000000000000,200000000000000000,E,MCC,ORDER\n")); err != nil {
			log.Fatal(err)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
