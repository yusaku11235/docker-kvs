package main

import (
	"fmt"
	"net/http"
	
)



func main() {
	http.HandleFunc("/a",handler1)
	http.HandleFunc("/b",handler2)

	http.ListenAndServe(":30",nil)
}

func handler1 (w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"hello")
}

func handler2 (w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"www")
}


// func main() {
// 	c:=make(chan int)
// 	for i:=0;i<5;i++ {
// 		go sleepyGopher(i,c)
// 	}

// 	for i:=0;i<5;i++ {
// 		gopherID:=<-c
// 		fmt.Println("gopher",gopherID,"はスリープを終えました")
// 	}
// }

// func sleepyGopher(id int, c chan int)  {
// 	time.Sleep(3*time.Second)
// 	fmt.Println("...",id,"snore...")
// 	c<-id
// }




// func dump(world string,slice []string ) {
// 	fmt.Printf("%v; 長さ:%v 容量%v %v\n",world,len(slice),cap(slice),slice)
// }

// type budget float64
// type rate float64
// type year float64

// func benStock (m budget, r rate, y year)  {
	
// 	for i := 1; i <= int(y); i++ {
// 		number:=math.Pow(float64(1+r/100),float64(i))
// 		total:=m*budget(number)
// 		fmt.Printf("%v年目:%.1f\n",i,total)
// 	}
// }

// type hope float64

// func initPay (r rate, y year, h hope) {
// 	var m budget 
// 	var total float64
// 	number:=math.Pow(float64(1+r/100),float64(y))
// 	for m = 1; total < float64(h); m+=1 {
// 		total:=m*budget(number)
// 		if total >= budget(h) {
// 			benefit:=total-m
// 			fmt.Printf("初期費用:%.1f, 合計:%.1f, 利益:%.1f\n",m,total,benefit)
// 			break
// 		}
// 	}
	
// }

// func main() {
// 	// }
// 	// dump("planets",planets)
// 	// planets2:=append(planets,"aaa","bbb")
// 	// dump("planets2",planets2)
// 	// planets3:=append(planets2,"3","4","5")
// 	// dump("planets3",planets3)

// 	benStock(61400000,5,10)
// 	initPay(5,20,100000000)
// }

// // type stardater interface {
// // 	YearDay() int
// // 	Hour() int
// // }

// // func stardate(s stardater) float64 {
// // 	d:=float64(s.YearDay())
// // 	h:=float64(s.Hour())/24
// // 	return 1000+d+h
// // }

// // type sol int

// // func (s sol) YearDay() int {
// // 	return int(s%668)
// // }

// // func (s sol) Hour() int {
// // 	return int(s/24)
// // }

// // func main() {
// // 	t:=time.Now()
// // 	fmt.Printf("%.1f\n",stardate(t))
// // 	s:=sol(1444)
// // 	fmt.Printf("%.1f\n",stardate(s))
// // }


