package main

import "fmt"

type Data struct {
	Status string
    Id int
}


func CloneDanModifikasi(dataAsli *Data) Data {
    // 1. Buat salinan dari data asli
    fmt.Println(dataAsli.Id)
    // 2. Lakukan modifikasi pada salinan
    newData.Status = "Diproses"
    // 3. Kembalikan salinan baru
    return newData
}

func main() {
    fmt.Println(Data{Status: "Pending", Id: 1})
    fmt.Println(&Data{Status: "Pending", Id: 1})
	a := CloneDanModifikasi(&Data{Status: "Pending", Id: 1})
	fmt.Println(a)
	fmt.Println(&a)
}