package main
import (
	"fmt"
	"os"
	"rtraversal/util"
)

type Robot struct{
	p int
	q int
	pos string
}

var dict1 = map[string]string {
	"NL": "W",
	"NR": "E",
	"SL": "E",
	"SR": "W",
	"EL": "N",
	"ER": "S",
	"WL": "S",
	"WR": "N",        
}

var dict2 = map[string]int {
	"NX" : 0,
	"NY" : 1,
	"EX" : 1,
	"EY" : 0,
	"WX" : -1,
	"WY" : 0,
	"SX" : 0,
	"SY" : -1,
}

func (r Robot) move(cmd string)(int, int, string) {
	if cmd == "M" {
		dir_x := r.pos+"X"
		dir_y := r.pos+"Y"
		dx := dict2[dir_x] 
		dy := dict2[dir_y] 
	
			r.p += dx 
			r.q += dy		
	}else{
		r.pos = dict1[r.pos+cmd]
	}
	return r.p, r.q, r.pos
	
  }

func main() {
	
	
	var rectDimension string
	fmt.Println("Enter rectangle dimension")
	fmt.Scanln(&rectDimension)
	if !util.LengthCheck(rectDimension, "Dimension"){
		fmt.Println("Invalid Dimension")
		os.Exit(1)
	}
	M, N := util.ValidateRectDimension(rectDimension)
	
	var robotPosition string
	fmt.Println("Enter robot position")
	fmt.Scanln(&robotPosition)
	if !util.LengthCheck(robotPosition, "Position"){
		fmt.Println("Invalid Robot Position")
		os.Exit(1)
	}
	x, y, pos := util.ValidateRobotPosition(robotPosition)
	
	var commands string
	fmt.Println("Enter comands")
	fmt.Scanln(&commands)
	if !util.LengthCheck(robotPosition, "Commands"){
		fmt.Println("Invalid Robot Position")
		os.Exit(1)
	}
	if !util.ValidateCommands(commands){
		fmt.Println("Enter valid commands")
		os.Exit(1)
	}

	
	plane := make([][]int, M)
	for i := 0; i < M; i++ {
		plane[i] = make([]int, N)
	}
	

	

	r  := Robot{x, y, pos}
	for i, c := range commands {
		fmt.Println(i, " => ", string(c))
		
		x1, y1, pos1 := r.move(string(c))
		fmt.Println(x1,y1, pos1)

		stopper := 0
		if x1!=M && y1 != N {
			stopper = plane[x1][y1]
		}
		if x1 > M || x1 < 0 ||y1 >N || y1 < 0 || stopper== 1 {
			break
		}else{
			r.p, r.q, r.pos = x1, y1, pos1 
			if r.pos != pos1{
			plane[x1][y1] = 1
			}
		}
	}
	
	fmt.Println(r.p, r.q, r.pos)



}


