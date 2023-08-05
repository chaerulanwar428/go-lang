package main

import "fmt"

func main() {
	var mounts = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"Setptember",
		"Oktober",
		"November",
		"Desember"
	}


	var slice1 = mounts[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// jika arraynya diubah maka hasil slicenya juga akan berubah
	// maka hati-hati dalam mengubah array
	// mounts[5] = "diubah"
	// fmt.Println(slice1)

	//begitu pun sebaliknya,slice berubah arraynya pun ikut berubah
	//slice[0] = "Mei Lagi"
	//fmt.Println(mouths)

	var slice2 = mounths[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2,"chaerul")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(mouths)

	newSlice := make([]string,2,5)

	newSlice[0] = "chaerul"
	newSlice[1] = "anwar"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string,len(newSlice),cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)
}