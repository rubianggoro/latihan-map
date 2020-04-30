package main

import "fmt"

func main() {
	var namaBarang = make(map[int]string)
	var hargaBarang = make(map[int]int)
	var stockBarang = make(map[int]int)

	namaBarang[0] = "Indomie"
	hargaBarang[0] = 1500
	stockBarang[0] = 9

	namaBarang[1] = "Pop Mie"
	hargaBarang[1] = 1500
	stockBarang[1] = 19

	namaBarang[2] = "Sedapp"
	hargaBarang[2] = 1500
	stockBarang[2] = 5

	for index, value := range stockBarang {
		if value < 10 {
			for index1, value1 := range namaBarang {
				for index2, value2 := range hargaBarang {
					if index == index1 && index == index2 {
						fmt.Println(value1, value2)
					}
				}
			}
		}
	}

}

// listBarang := make(map[int]item)
// listBarang[0] = item{namaProduk: "Mie Gelas", hargaProduk: 1000, stokProduk: 8}
// listBarang[1] = item{namaProduk: "Pop Mie", hargaProduk: 1300, stokProduk: 18}
// listBarang[2] = item{namaProduk: "Indomie", hargaProduk: 900, stokProduk: 9}
// listBarang[3] = item{namaProduk: "Sedap", hargaProduk: 2300, stokProduk: 5}
// listBarang[4] = item{namaProduk: "Sarimi", hargaProduk: 500, stokProduk: 24}

// fmt.Println("SEMUA BARANG :")
// for _, value1 := range listBarang {
// 	fmt.Println("\nNamaBarang :", value1.namaProduk, "\nHargaBarang :", value1.hargaProduk, "\nStockBarang :", value1.stokProduk)
// }
// fmt.Println("\nBarang yang Stocknya Di bawah 10 :")
// for _, value2 := range listBarang {
// 	if value2.stokProduk < 10 {
// 		fmt.Println("Nama Barang:", value2.namaProduk)
// 		fmt.Println("Stock :", value2.stokProduk)
// 	}

// }
