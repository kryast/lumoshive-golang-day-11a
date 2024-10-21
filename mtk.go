package MTK

// func Tambah(a int, b int) int {
// 	return a + b
// }

// func Kurang(a int, b int) int {
// 	return a - b
// }

// func Kali(a int, b int) int {
// 	return a * b
// }

// func Bagi(a int, b int) int {
// 	if b == 0 {
// 		return 0
// 	}
// 	return a / b
// }

type Data struct {
	Id int
}

var DataAkun []Data

func GetId(length int, ch chan<- []Data) {
	for i := 0; i < length; i++ {
		DataAkun = append(DataAkun, Data{Id: i})
	}
	ch <- DataAkun

}
