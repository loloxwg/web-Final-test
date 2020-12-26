package main

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers,Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}

type Pokers []Poker

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PrintPoker()," ")
	}
	fmt.Println()
}










func main()  {

	a:=[]int{3,5,-1,6,35,56}
	sort.Ints(a)
	fmt.Println(a)


}

