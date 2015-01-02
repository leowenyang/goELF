package goELF

import (
  "fmt"
)

type ELFHeader struct {
  e_ident     [10]byte
  e_type      int16
  e_machine   int16
  e_version   int32
  e_entry     int32
  e_phoff     int32
  e_shoff     int32
  e_flags     int32
  e_ehsize    int16
  e_phentsize int16
  e_phnum     int16
  e_shentsize int16
  e_shunum    int16
  e_shstrndx  int16
}

func NewELFHeader(data []byte) * ELFHeader {
  var tmpData int32
  offset := 0
  h := &ELFHeader{}
  // e_ident
  for i:= 0; i < 10; i++ {
    h.e_ident[i] = data[i]
  }
  offset += 10
  fmt.Println("ident is ", h.e_ident)

  // e_type
  tmpData = int32(data[offset+1]  << 8)
  tmpData += int32(data[offset+1])
  h.e_type = int16(tmpData)
  offset += 2
  fmt.Println("tye is ", h.e_type)

  // e_machine_
  //tmpData = data[offset+1]  << 8
  //tmpData += data[offset+1]
  //h.e_type = tmpData
  //offset += 2

  return h;

}
