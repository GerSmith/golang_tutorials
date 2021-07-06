package main

import "fmt"

func main() {
	var (
		num    int // 000
		fstDig int
		sndDig int
		thdDig int
	)

	fmt.Scan(&num)

	fstDig = num / 100 % 10
	sndDig = num / 10 % 10
	thdDig = num % 10

	//fmt.Println(fstDig, sndDig, thdDig)
	if fstDig != sndDig && fstDig != thdDig && sndDig != thdDig {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
