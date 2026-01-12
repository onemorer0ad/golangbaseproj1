package coffee

import (
	"fmt"
	"strconv"
)

type Stats struct{
	Orders int
	Revenue int
}

func (s *Stats) AddOrder(price int) {
	s.Orders += 1
	s.Revenue += price
}

func (s *Stats) GetStats(){
	fmt.Println(strconv.Itoa(s.Orders) + "- orders, revenue-  " + strconv.Itoa(s.Revenue))
}