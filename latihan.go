package main
import "fmt"

func main(){
	var n, i int
	fmt.Scan(&n)
	i = jumlahDigit(n)
	fmt.Print(i)
}
func jumlahDigit(n int)int{
	var hasil int
	hasil = 0
	if n < 9{
		return n
	}else{
		hasil = n%10
		return hasil+jumlahDigit(n/10)
	}
}