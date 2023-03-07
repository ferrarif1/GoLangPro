package work

import(
	"fmt"
)

func TestGOTO1()  {
	fmt.Println("")
	var i int  = 4
LB:
    i +=3
	if i>200{
		return
	}
	goto LB
}

func ifGoTo(){
	var i int = 2
	if i%2 ==0 {
		goto L1
	}else{
		goto L2
	}
L1:
    i+= 4
L2: 
    i*=3
}
