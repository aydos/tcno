# tcno
Türkiye Cumhuriyeti Kimlik Numarası Kütüphanesi

## Kullanım
```go
package main

import (
  "fmt"
  "github.com/aydos/tcno"
)

func main() {
  t := tcno.Random()
  fmt.Printf("%v %T\n", t, t)     // int türünde
  if tcno.IsValid(t) {
    fmt.Println(t, "geçerli bir TC kimlik numarasıdır.")
  } else {
    fmt.Println(t, "geçerli değil.")
  }
  g := tcno.Generate(123456789)
  if tcno.IsValid(g) {
    fmt.Println(g)
  }
  h := 12345678
  if tcno.IsValid(tcno.Generate(h))==false {
    fmt.Println(h, "hatalı")
  }
}
```
