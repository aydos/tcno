// github.com/aydos/tcno
package tcno

import (
	"math/rand"
	"strconv"
	"time"
)

var oldrandnumber = 123456789

const (
	mintcbase = 100000000
	maxtcbase = 999999999
	mintcno   = 10000000000
	maxtcno   = 99999999999
)

// getC, verilen 9 haneli sayıdan TCNo için kontrol basamakları hesaplar
// Hata kontrolü yapmaz.
func getC(t int) int {
	s := strconv.Itoa(t)
	a1 := int(s[0]) - 48
	a2 := int(s[1]) - 48
	a3 := int(s[2]) - 48
	a4 := int(s[3]) - 48
	a5 := int(s[4]) - 48
	a6 := int(s[5]) - 48
	a7 := int(s[6]) - 48
	a8 := int(s[7]) - 48
	a9 := int(s[8]) - 48
	n := 7*(a1+a3+a5+a7+a9) - (a2 + a4 + a6 + a8)
	if n < 0 {
		n = -n
	}
	c1 := n % 10
	c2 := (a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + c1) % 10
	return c1*10 + c2
}

// IsValid, verilen bir TCNo'nun geçerli olup olmadığını bulur
func IsValid(tcno int) bool {
	if tcno < mintcno || tcno > maxtcno {
		return false
	}
	c := getC(tcno)
	if tcno%100 == c {
		return true
	}
	return false
}

// Generate, ilk 9 hanesi verilen sayıdan TCNo üretir
// Hatalı girdide 0-sıfır döner
func Generate(tcbase int) int {
	if tcbase < mintcbase || tcbase > maxtcbase {
		return 0
	}
	c := getC(tcbase)
	return tcbase*100 + c
}

// Random, rastgele bir TCNo üretir
func Random() int {
	rand.Seed(time.Now().UnixNano() + int64(oldrandnumber))
	tcno := rand.Intn(maxtcbase-mintcbase+1) + mintcbase
	oldrandnumber = tcno
	c := getC(tcno)
	return tcno*100 + c
}
