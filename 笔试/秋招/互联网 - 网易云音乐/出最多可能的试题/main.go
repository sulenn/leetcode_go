package main

import (
	"fmt"
	"io"
)

func main() {
	var E, EM, M, MH, H int
	for {
		_, err := fmt.Scan(&E, &EM, &M, &MH, &H)
		if err == io.EOF {
			break
		}
		result := 0
		if E <= M && E <= H {
			result += E
			E = 0
			M -= E
			H -= E
		}
		if M <= E && M <= H {
			result += M
			M = 0
			E -= M
			H -= M
		}
		if H <= E && H <= M {
			result += H
			H = 0
			E -= H
			M -= H
		}
		for EM > 0 && MH > 0 {
			if E == 0 {
				EM--
			}
			if H == 0 {
				MH--
			}
			if M == 0 {
				if EM > MH {
					EM--
				} else {
					MH--
				}
			}
			result++
		}
		for EM > 0 && H > 0 {
			if EM > 0 && E == 0 {
				EM--
			}
			if EM > 0 && M == 0 {
				EM--
			}
			result++
		}
		for MH != 0 && E != 0 {
			if MH > 0 && M == 0 {
				MH--
			}
			if MH > 0 && H == 0 {
				MH--
			}
			result++
		}
		fmt.Println(result)
	}
}
