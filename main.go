package main

import (
	"fmt"
	// "strings"
)

// func main1() {

// 	// 1

// 	dayOfTheWeek := map[string]int{
// 		"Mon":  1,
// 		"Tues": 2,
// 		"Wens": 3,
// 		"Thus": 4,
// 		"Frid": 5,
// 		"Satu": 6,
// 		"Sund": 7,
// 	}

// 	fmt.Println(dayOfTheWeek["Mon"])

// 	fmt.Println(dayOfTheWeek["Sun"])
// мы получили 0 потому что в несущ мапах в го задаётся значение поумалчанию-"0"-

// 2

// }

// func getVisits(m map[string]int, day string) (int, bool) {
// 	v, ok := m[day]
// 	count := 0
// if  ok {
// 	count = v
// } else {
// 	return 0, ok
// }

// return  count,ok
// }

// func main() {
// 	dayOfTheWeek := map[string]int{
// 		"Mon":  1,
// 		"Tues": 2,
// 		"Wens": 3,
// 		"Thus": 4,
// 		"Frid": 5,
// 		"Satu": 6,
// 		"Sund": 7,
// 	}

// 	fmt.Println(getVisits(dayOfTheWeek,"Tues"))

// 	fmt.Println(getVisits(dayOfTheWeek,"Tu"))
// }
// }

// 3

// func main() {
// 	dayOfTheWeek := map[string]int{
// 		"Mon":  1,
// 		"Tues": 2,
// 		"Wens": 3,
// 		"Thus": 4,
// 		"Frid": 5,
// 		"Satu": 6,
// 		"Sund": 7,
// 	}

// // начальная длина
// 	fmt.Println(len(dayOfTheWeek))

// // удаление элимента
// 	delete(dayOfTheWeek,"Mon")
// //после удаления
// fmt.Println(len(dayOfTheWeek))
// }

// 4

// func main5()  {
// 	Ch_r := map[string]string{
// 		"Катар-юрт": "Ваха",
// 		"Урус-мартан": "Джамбулат",
// 		"Ставрапромасольский": "Дени",
// 		"Гер-зель": "Ислам ",
// 	}

// 	for  k, velue :=  range Ch_r  {
// 		fmt.Println(k,velue)
		
// 	}
// }

// 5

// func main()  {
// 	var m map[string]int 
// 	// m["visits"] = 1
// 	// panic: assignment to entry in nil map
// 	m = make(map[string]int)
// 		m["visits"] = 1

// 		fmt.Println(m["visits"])
// }

// 6

// func main()  {

// // Map := map[[]int]string{}

// m:= map[string]string{
// 		"1": "1",
// 		"2": "2",
// 		"3": "3",
// 	}
	
// 	fmt.Println(m["ключ1"])

// }

// 7 

// type Student struct {
//     Name   string
//     Active bool
// }



// func main()  {
// 	m := map[string]Student {
// 		"str1": {Name: "vaha"},
// 		"str2": {Name: "djamb"},
// 	}

// 	fmt.Println(m["str1"])
// }

// 8

// func main()  {
// 	a := map[string]int{"Mon": 10}
	
// 	b := a
// 	b["Mon"] = 99

// 	fmt.Println(a,b)
// 	fmt.Println(clone(a))
// }

// func clone(src map[string]int) map[string]int {
// src = make(map[string]int)
// src["Mon"] = 98
// return src
// }

// 9 

// func WordCount(s string) map[string]int {
// 	myMap := make(map[string]int)

// 	words := strings.Fields(s)
// 	for i, word := range words {
// 		myMap[word] += i
// 	}
// 	return myMap
// }

// func main()  {
// 	fmt.Println(WordCount("go go gopher"))
// }

// 10




func Add(phones map[string]string, name, number string){
phones[name] = number
}


func Get(phones map[string]string, name string) (string, bool){
value,status := phones[name]
fmt.Println("GEEeeeeet",value, status)
return value,status
}

func Remove(phones map[string]string, name string){
delete(phones,"")
}

func main()  {
//

	phones := map[string]string{}
Add(phones, "vaha","+79388890223",)

Add(phones, "mustafa","+79388880970")

// 
Get(phones, "vahja")
for get, getted := range phones{
	fmt.Println(get,getted)
}
fmt.Println(phones)
// 

delete(phones, "vaha")
fmt.Println(phones)

}