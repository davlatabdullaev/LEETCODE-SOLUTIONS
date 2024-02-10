package addbinary

import (
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
)

func AddBinary(a string, b string) string {

	// a1, err := strconv.Atoi(a)
	// if err != nil {
	// 	log.Println("error while convert a string to int")
	// }

	b1, err := strconv.Atoi(b)
	if err != nil {
		log.Println("error while convert b string to int")
	}

	w := 1

	aa := make([]byte, 8)
	binary.LittleEndian.PutUint32(aa, uint32(w))
    fmt.Println(aa)

	bb := make([]byte, 8)
	binary.LittleEndian.PutUint32(bb, uint32(b1))
    fmt.Println(bb)

   return ""
}
