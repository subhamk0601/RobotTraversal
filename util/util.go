package util

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)


func ValidateRectDimension(inp string)(int, int){
	m,n :=0,0
	for i, c := range inp {
		x, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("Invalid Rectangle dimension")
			os.Exit(1)
		}
		if i == 0 {
			m = x
		}else{
			n = x
		}

	}
	return m,n
		
}

func ValidateRobotPosition(inp string)(int, int, string){
	p,q,r :=0,0,""
	for i, c := range inp {
		if i==0 || i ==1{
		x, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Println("Invalid Robot position")
			os.Exit(1)
		}
		if i == 0 {
			p = x
		}else{
			q = x
		}
	
		}else {
			r = string(c)
		}

	}
	return p,q,r
		
}

func ValidateCommands(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func LengthCheck(inp string, inpType string ) bool{
		if inpType == "Dimension"{
			if len(inp) != 2 {
				return false
			}
		}
		if inpType == "Commands"{
			if len(inp) <=0  {
				return false
			}
		}
		if inpType == "Position"{
			if len(inp) !=3  {
				return false
			}
		}
		return true
	}