package main 
import "fmt"

// mencari produk berdasarkan ukuran / warna 

const DATAMAX int = 1000
type pakaian struct{
	id, stok int 
	nama, warna, ukuran string 
}
type datapakaian [DATAMAX]pakaian
var daftarToko datapakaian
var jumlahData int = 0

func menu_utama(){
	fmt.Println("+-----------------------------+")
	fmt.Println("|     Selamat Datang Di       |")
	fmt.Println("|  Aplikasi Manajemen Fashion |")
	fmt.Println("+-----------------------------+")
	fmt.Printf("| %-27s |\n", "[1] Daftar Data Pakaian") 
	fmt.Printf("| %-27s |\n", "[2] Tambah  Data Pakaian") 
	fmt.Printf("| %-27s |\n", "[3] Edit Data Pakaian")
	fmt.Printf("| %-27s |\n", "[4] Hapus Data Pakaian")
	fmt.Printf("| %-27s |\n", "[5] Search Data Pakaian")  
	fmt.Printf("| %-27s |\n", "[6] Sortir Data Pakaian") 
	fmt.Printf("| %-27s |\n", "[0] Exit")
	fmt.Println("+-----------------------------+")
	fmt.Print("Pilih [1 - 6]?")
}

//cetak Data
func cetakData(A datapakaian, awal int, akhir int){
	var i int 
	
	fmt.Println("\n+-------+----------------------+-----------------+------------+-------+")
	fmt.Printf("| %-5s | %-20s | %-15s | %-10s | %-5s |\n", "ID", "Nama Pakaian", "Warna", "Ukuran", "Stok")
	fmt.Println("+-------+----------------------+-----------------+------------+-------+")
	
	for i = awal; i <= akhir; i++ {
		fmt.Printf("| %-5d | %-20s | %-15s | %-10s | %-5d |\n", 
			A[i].id, A[i].nama, A[i].warna, A[i].ukuran, A[i].stok)
	}
	fmt.Println("+-------+----------------------+-----------------+------------+-------+\n")
}

//kay
//menu tambah data
func menu_tambahdata(){
	
}

//kay
//menu edit data 
func menu_editdata(){
	
}

//kay
//menu hapus data
func menu_hapusdata(){
	
}

func inisialisasiData(){
	daftarToko[0] = pakaian{id: 1, nama: "Kaos Polos", warna: "Merah", ukuran: "M", stok: 15}
	daftarToko[1] = pakaian{id: 2, nama: "Kemeja Flanel", warna: "Biru", ukuran: "L", stok: 10}
	daftarToko[2] = pakaian{id: 3, nama: "Celana Chino", warna: "Hitam", ukuran: "L", stok: 20}
	daftarToko[3] = pakaian{id: 4, nama: "Jaket Bomber", warna: "Hijau", ukuran: "XL", stok: 8}
	daftarToko[4] = pakaian{id: 5, nama: "Sweater Crewneck", warna: "Putih", ukuran: "S", stok: 12}
	daftarToko[5] = pakaian{id: 6, nama: "Kaos Polos", warna: "Hitam", ukuran: "XL", stok: 25}
	daftarToko[6] = pakaian{id: 7, nama: "Kemeja Kerja", warna: "Putih", ukuran: "M", stok: 18}
	daftarToko[7] = pakaian{id: 8, nama: "Celana Jeans", warna: "Biru", ukuran: "M", stok: 14}
	daftarToko[8] = pakaian{id: 9, nama: "Hoodie Oversize", warna: "Abu-abu", ukuran: "XL", stok: 7}
	daftarToko[9] = pakaian{id: 10, nama: "Blouse Wanita", warna: "Merah Muda", ukuran: "S", stok: 11}
	daftarToko[10] = pakaian{id: 11, nama: "Rok Panjang", warna: "Cokelat", ukuran: "M", stok: 9}
	daftarToko[11] = pakaian{id: 12, nama: "Kaos Raglan", warna: "Merah", ukuran: "L", stok: 22}
	daftarToko[12] = pakaian{id: 13, nama: "Kemeja Batik", warna: "Cokelat", ukuran: "XL", stok: 5}
	daftarToko[13] = pakaian{id: 14, nama: "Celana Pendek", warna: "Hijau", ukuran: "M", stok: 30}
	daftarToko[14] = pakaian{id: 15, nama: "Jaket Parka", warna: "Hitam", ukuran: "L", stok: 6}
	daftarToko[15] = pakaian{id: 16, nama: "Cardigan Knitted", warna: "Krem", ukuran: "S", stok: 13}
	daftarToko[16] = pakaian{id: 17, nama: "Kaos V-Neck", warna: "Putih", ukuran: "M", stok: 17}
	daftarToko[17] = pakaian{id: 18, nama: "Tunik Dress", warna: "Hijau Mint", ukuran: "L", stok: 10}
	daftarToko[18] = pakaian{id: 19, nama: "Celana Jogger", warna: "Abu-abu", ukuran: "XL", stok: 16}
	daftarToko[19] = pakaian{id: 20, nama: "Rompi Vest", warna: "Hitam", ukuran: "M", stok: 12}
	daftarToko[20] = pakaian{id: 21, nama: "Kaos Polo", warna: "Biru", ukuran: "L", stok: 19}
	daftarToko[21] = pakaian{id: 22, nama: "Kemeja Denim", warna: "Biru", ukuran: "XL", stok: 9}
	daftarToko[22] = pakaian{id: 23, nama: "Celana Cargo", warna: "Hijau", ukuran: "L", stok: 14}
	daftarToko[23] = pakaian{id: 24, nama: "Jaket Denim", warna: "Biru", ukuran: "M", stok: 11}
	daftarToko[24] = pakaian{id: 25, nama: "Sweater Hoodie", warna: "Hitam", ukuran: "S", stok: 15}
	daftarToko[25] = pakaian{id: 26, nama: "Kaos Polos", warna: "Kuning", ukuran: "M", stok: 21}
	daftarToko[26] = pakaian{id: 27, nama: "Kemeja Formal", warna: "Hitam", ukuran: "L", stok: 13}
	daftarToko[27] = pakaian{id: 28, nama: "Celana Formal", warna: "Hitam", ukuran: "M", stok: 17}
	daftarToko[28] = pakaian{id: 29, nama: "Coach Jacket", warna: "Merah", ukuran: "XL", stok: 8}
	daftarToko[29] = pakaian{id: 30, nama: "Blouse Silk", warna: "Putih", ukuran: "S", stok: 10}
	daftarToko[30] = pakaian{id: 31, nama: "Rok Plisket", warna: "Hijau", ukuran: "M", stok: 12}
	daftarToko[31] = pakaian{id: 32, nama: "Kaos Strip", warna: "Hitam", ukuran: "L", stok: 25}
	daftarToko[32] = pakaian{id: 33, nama: "Kemeja Pantai", warna: "Kuning", ukuran: "XL", stok: 14}
	daftarToko[33] = pakaian{id: 34, nama: "Celana Boxer", warna: "Merah", ukuran: "S", stok: 40}
	daftarToko[34] = pakaian{id: 35, nama: "Windbreaker", warna: "Biru", ukuran: "L", stok: 7}
	daftarToko[35] = pakaian{id: 36, nama: "Outer Lace", warna: "Putih", ukuran: "M", stok: 9}
	daftarToko[36] = pakaian{id: 37, nama: "Kaos Oversize", warna: "Ungu", ukuran: "XL", stok: 18}
	daftarToko[37] = pakaian{id: 38, nama: "Midi Dress", warna: "Merah Muda", ukuran: "L", stok: 6}
	daftarToko[38] = pakaian{id: 39, nama: "Sweatpants", warna: "Hitam", ukuran: "M", stok: 22}
	daftarToko[39] = pakaian{id: 40, nama: "Leather Jacket", warna: "Hitam", ukuran: "XL", stok: 4}
	daftarToko[40] = pakaian{id: 41, nama: "Kaos Polos", warna: "Orange", ukuran: "S", stok: 16}
	daftarToko[41] = pakaian{id: 42, nama: "Kemeja Linen", warna: "Krem", ukuran: "L", stok: 11}
	daftarToko[42] = pakaian{id: 43, nama: "Celana Kulot", warna: "Cokelat", ukuran: "M", stok: 15}
	daftarToko[43] = pakaian{id: 44, nama: "Varsity Jacket", warna: "Biru", ukuran: "XL", stok: 5}
	daftarToko[44] = pakaian{id: 45, nama: "Turtleneck", warna: "Hitam", ukuran: "S", stok: 13}
	daftarToko[45] = pakaian{id: 46, nama: "Kaos Grafik", warna: "Putih", ukuran: "M", stok: 20}
	daftarToko[46] = pakaian{id: 47, nama: "Kemeja Pendek", warna: "Hijau", ukuran: "L", stok: 12}
	daftarToko[47] = pakaian{id: 48, nama: "Legging Sport", warna: "Hitam", ukuran: "M", stok: 25}
	daftarToko[48] = pakaian{id: 49, nama: "Blazer Formal", warna: "Abu-abu", ukuran: "L", stok: 8}
	daftarToko[49] = pakaian{id: 50, nama: "Maxi Dress", warna: "Biru", ukuran: "XL", stok: 7}
	daftarToko[50] = pakaian{id: 51, nama: "Kaos Singlet", warna: "Putih", ukuran: "M", stok: 35}
	daftarToko[51] = pakaian{id: 52, nama: "Kemeja Kerja", warna: "Biru Muda", ukuran: "S", stok: 16}
	daftarToko[52] = pakaian{id: 53, nama: "Celana Tartan", warna: "Merah", ukuran: "L", stok: 11}
	daftarToko[53] = pakaian{id: 54, nama: "Anorak Jacket", warna: "Hijau", ukuran: "XL", stok: 6}
	daftarToko[54] = pakaian{id: 55, nama: "Cardigan Polos", warna: "Hitam", ukuran: "M", stok: 14}
	daftarToko[55] = pakaian{id: 56, nama: "Kaos Polos", warna: "Abu-abu", ukuran: "L", stok: 28}
	daftarToko[56] = pakaian{id: 57, nama: "Kemeja Shanghai", warna: "Putih", ukuran: "XL", stok: 10}
	daftarToko[57] = pakaian{id: 58, nama: "Celana Denim", warna: "Hitam", ukuran: "M", stok: 19}
	daftarToko[58] = pakaian{id: 59, nama: "Fleece Jacket", warna: "Cokelat", ukuran: "L", stok: 8}
	daftarToko[59] = pakaian{id: 60, nama: "Crop Top", warna: "Merah Muda", ukuran: "S", stok: 15}
	daftarToko[60] = pakaian{id: 61, nama: "Kaos Raglan", warna: "Biru", ukuran: "M", stok: 17}
	daftarToko[61] = pakaian{id: 62, nama: "Kemeja Flannel", warna: "Merah", ukuran: "L", stok: 12}
	daftarToko[62] = pakaian{id: 63, nama: "Celana Chino", warna: "Krem", ukuran: "XL", stok: 13}
	daftarToko[63] = pakaian{id: 64, nama: "Puffer Jacket", warna: "Hitam", ukuran: "L", stok: 5}
	daftarToko[64] = pakaian{id: 65, nama: "Sweater Rajut", warna: "Maroon", ukuran: "M", stok: 11}
	daftarToko[65] = pakaian{id: 66, nama: "Kaos Polo", warna: "Hijau", ukuran: "XS", stok: 20}
	daftarToko[66] = pakaian{id: 67, nama: "Kemeja Kotak", warna: "Biru", ukuran: "XL", stok: 14}
	daftarToko[67] = pakaian{id: 68, nama: "Celana Pendek", warna: "Hitam", ukuran: "L", stok: 32}
	daftarToko[68] = pakaian{id: 69, nama: "Track Jacket", warna: "Merah", ukuran: "M", stok: 9}
	daftarToko[69] = pakaian{id: 70, nama: "Blouse Katun", warna: "Kuning", ukuran: "S", stok: 12}
	daftarToko[70] = pakaian{id: 71, nama: "Rok Mini", warna: "Hitam", ukuran: "XS", stok: 10}
	daftarToko[71] = pakaian{id: 72, nama: "Kaos V-Neck", warna: "Abu-abu", ukuran: "L", stok: 24}
	daftarToko[72] = pakaian{id: 73, nama: "Kemeja Satin", warna: "Emas", ukuran: "M", stok: 7}
	daftarToko[73] = pakaian{id: 74, nama: "Celana Slimfit", warna: "Abu-abu", ukuran: "XL", stok: 15}
	daftarToko[74] = pakaian{id: 75, nama: "Biker Jacket", warna: "Cokelat", ukuran: "L", stok: 4}
	daftarToko[75] = pakaian{id: 76, nama: "Hoodie Polos", warna: "Hijau", ukuran: "M", stok: 16}
	daftarToko[76] = pakaian{id: 77, nama: "Kaos Polos", warna: "Navy", ukuran: "XL", stok: 30}
	daftarToko[77] = pakaian{id: 78, nama: "Kemeja Oxford", warna: "Biru Muda", ukuran: "L", stok: 13}
	daftarToko[78] = pakaian{id: 79, nama: "Celana Cargo", warna: "Hitam", ukuran: "M", stok: 18}
	daftarToko[79] = pakaian{id: 80, nama: "Long Coat", warna: "Krem", ukuran: "XXL", stok: 5}
	daftarToko[80] = pakaian{id: 81, nama: "Tshirt Stripe", warna: "Merah", ukuran: "S", stok: 22}
	daftarToko[81] = pakaian{id: 82, nama: "Kemeja Formal", warna: "Putih", ukuran: "XXXL", stok: 20}
	daftarToko[82] = pakaian{id: 83, nama: "Celana Jogger", warna: "Hitam", ukuran: "L", stok: 17}
	daftarToko[83] = pakaian{id: 84, nama: "Raincoat", warna: "Kuning", ukuran: "L", stok: 8}
	daftarToko[84] = pakaian{id: 85, nama: "Sweater Vneck", warna: "Biru", ukuran: "M", stok: 12}
	daftarToko[85] = pakaian{id: 86, nama: "Kaos Oversize", warna: "Putih", ukuran: "S", stok: 19}
	daftarToko[86] = pakaian{id: 87, nama: "Kemeja Batik", warna: "Merah", ukuran: "M", stok: 10}
	daftarToko[87] = pakaian{id: 88, nama: "Celana Kulot", warna: "Hitam", ukuran: "XL", stok: 14}
	daftarToko[88] = pakaian{id: 89, nama: "Bomber Jacket", warna: "Maroon", ukuran: "L", stok: 7}
	daftarToko[89] = pakaian{id: 90, nama: "Tunik Polos", warna: "Pink", ukuran: "M", stok: 11}
	daftarToko[90] = pakaian{id: 91, nama: "Kaos Polos", warna: "Tosca", ukuran: "L", stok: 15}
	daftarToko[91] = pakaian{id: 92, nama: "Kemeja Pria", warna: "Abu-abu", ukuran: "M", stok: 16}
	daftarToko[92] = pakaian{id: 93, nama: "Celana Sirwal", warna: "Hitam", ukuran: "XXL", stok: 12}
	daftarToko[93] = pakaian{id: 94, nama: "Harrington", warna: "Navy", ukuran: "XXXL", stok: 6}
	daftarToko[94] = pakaian{id: 95, nama: "Cardigan Rajut", warna: "Abu-abu", ukuran: "S", stok: 13}
	daftarToko[95] = pakaian{id: 96, nama: "Kaos Distro", warna: "Hitam", ukuran: "M", stok: 25}
	daftarToko[96] = pakaian{id: 97, nama: "Kemeja Tactical", warna: "Hijau", ukuran: "XXL", stok: 9}
	daftarToko[97] = pakaian{id: 98, nama: "Celana Tactical", warna: "Hijau", ukuran: "L", stok: 11}
	daftarToko[98] = pakaian{id: 99, nama: "Jaket Hoodie", warna: "Putih", ukuran: "M", stok: 14}
	daftarToko[99] = pakaian{id: 100, nama: "Blouse Casual", warna: "Biru", ukuran: "XS", stok: 12}
	
	jumlahData = 100
}

func daftarpakaian(){
	var i int 
	
	fmt.Println("\n+-------+----------------------+-----------------+------------+-------+")
	fmt.Printf("| %-5s | %-20s | %-15s | %-10s | %-5s |\n", "ID", "Nama Pakaian", "Warna", "Ukuran", "Stok")
	fmt.Println("+-------+----------------------+-----------------+------------+-------+")
	
	for i = 0; i < jumlahData; i++ {
		fmt.Printf("| %-5d | %-20s | %-15s | %-10s | %-5d |\n", 
			daftarToko[i].id, 
			daftarToko[i].nama, 
			daftarToko[i].warna, 
			daftarToko[i].ukuran, 
			daftarToko[i].stok)
	}
	fmt.Println("+-------+----------------------+-----------------+------------+-------+\n")
}

//menu searching
func menu_searching() {
	var pilih, pilih2, pilih3, pilih4 int
	var katakunci string
	
	clearScreen()
	
	fmt.Println("+------------------------------------------+")
	fmt.Println("|              Anda Berada Di              |")
	fmt.Println("|          Menu Pencarian Pakaian          |")
	fmt.Println("+------------------------------------------+")
	fmt.Printf("| %-40s |\n", "[1] Sequential Search") 
	fmt.Printf("| %-40s |\n", "[2] Binary Search") 
	fmt.Printf("| %-40s |\n", "[0] Menu Utama") 
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilih [1/2/0]? ")
	fmt.Scan(&pilih)
	
	if pilih == 1 {
		clearScreen() 
		fmt.Println("+------------------------------------------+")
		fmt.Println("|              Anda Berada Di              |")
		fmt.Println("|           Menu Sequential Search         |") 
		fmt.Println("+------------------------------------------+")
		fmt.Printf("| %-40s |\n", "[1] Ukuran") 
		fmt.Printf("| %-40s |\n", "[2] Warna") 
		fmt.Printf("| %-40s |\n", "[0] Menu Searching") 
		fmt.Println("+------------------------------------------+")
		fmt.Print("Pilih [1/2/0]? ")
		fmt.Scan(&pilih2)
		
		if pilih2 == 1 {
			fmt.Print("Ukuran yang dicari: ")
			fmt.Scan(&katakunci)
			//Kayla
			// Panggil fungsi pencarian sequential ukuran di sini
			
		} else if pilih2 == 2 {
			fmt.Print("Warna yang dicari: ")
			fmt.Scan(&katakunci)
			//Kayla
			// Panggil fungsi pencarian sequential warna di sini
			
		}

	} else if pilih == 2 {
		clearScreen() 
		fmt.Println("+------------------------------------------+")
		fmt.Println("|              Anda Berada Di              |")
		fmt.Println("|             Menu Binary Search           |")
		fmt.Println("+------------------------------------------+")
		fmt.Printf("| %-40s |\n", "[1] Ukuran") 
		fmt.Printf("| %-40s |\n", "[2] Warna") 
		fmt.Printf("| %-40s |\n", "[0] Menu Searching") 
		fmt.Println("+------------------------------------------+")
		fmt.Print("Pilih [1/2/0]? ")
		fmt.Scan(&pilih2)
		
		if pilih2 == 1 {
			clearScreen() 
			fmt.Println("+------------------------------------------+")
			fmt.Println("|              Anda Berada Di              |")
			fmt.Println("|         Menu Ukuran Binary Search        |")
			fmt.Println("+------------------------------------------+")
			fmt.Printf("| %-40s |\n", "[1] Ascending") 
			fmt.Printf("| %-40s |\n", "[2] Descending") 
			fmt.Printf("| %-40s |\n", "[0] Menu Searching") 
			fmt.Println("+------------------------------------------+")
			fmt.Print("Pilih [1/2/0]? ")
			fmt.Scan(&pilih3)
			
			if pilih3 == 1 {
				fmt.Print("Ukuran yang dicari (Binary Asc): ")
				fmt.Scan(&katakunci)
				ukuranBinarySearchAsc(daftarToko, jumlahData, katakunci)
			} else if pilih3 == 2 {
				fmt.Print("Ukuran yang dicari (Binary Desc): ")
				fmt.Scan(&katakunci)
				ukuranBinarySearchDesc(daftarToko, jumlahData, katakunci)
			} else if pilih3 == 0 {
				clearScreen()
				menu_searching()
			}
			
		} else if pilih2 == 2 {
			clearScreen() 
			fmt.Println("+------------------------------------------+")
			fmt.Println("|              Anda Berada Di              |")
			fmt.Println("|          Menu Warna Binary Search        |")
			fmt.Println("+------------------------------------------+")
			fmt.Printf("| %-40s |\n", "[1] Ascending") 
			fmt.Printf("| %-40s |\n", "[2] Descending") 
			fmt.Printf("| %-40s |\n", "[0] Menu Searching") 
			fmt.Println("+------------------------------------------+")
			fmt.Print("Pilih [1/2/0]? ")
			fmt.Scan(&pilih4)
			
			if pilih4 == 1 {
				fmt.Print("Warna yang dicari: ")
				fmt.Scan(&katakunci)
				warnaBinarySearchAsc(daftarToko, jumlahData, katakunci)
			} else if pilih4 == 2 {
				fmt.Print("Warna yang dicari: ")
				fmt.Scan(&katakunci)
				warnaBinarySearchDesc(daftarToko, jumlahData, katakunci)
			} else if pilih4 == 0 {
				clearScreen()
				menu_searching()
			}
		} else if pilih2 == 0 {
			clearScreen()
			menu_searching()
		}
	} else if pilih == 0 {
		clearScreen()
		menu_utama()
	}
}

//ubah huruf jadi abjad
func bobotUkuran(ukuran string) int {
	if ukuran == "XS" {
		return 1
	} else if ukuran == "S" {
		return 2
	} else if ukuran == "M" {
		return 3
	} else if ukuran == "L" {
		return 4
	} else if ukuran == "XL" {
		return 5
	} else if ukuran == "XXL" {
		return 6
	} else if ukuran == "XXXL" {
		return 7
	}
	return 0 // Jika ada ukuran di luar itu
}

//Kay
//ukuran sequential search
func ukuranSequentialSearch(A *datapakaian, n int, x string){
	
}

//Kay
//warna sequential search
func warnaSequentialSearch(A *datapakaian, n int, x string){
	
}

//ukuran binary search asc
func ukuranBinarySearchAsc(A *datapakaian, n int, x string){
	var left, right, mid int 
	var idx int 
	var batasKanan, batasKiri int
	
	idx = -1 
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right)/2
		if (*A)[mid].ukuran == x {
			idx = mid 
		} else if bobotUkuran(x) > bobotUkuran((*A)[mid].ukuran) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	if idx != -1 {
		
		batasKiri = idx 
		batasKanan = idx
		
		// Cari tahu seberapa jauh data yang sama di sebelah kiri tanpa break
		for batasKiri > 0 && (*A)[batasKiri-1].ukuran == x {
			batasKiri--
		}
		
		// Cari tahu seberapa jauh data yang sama di sebelah kanan tanpa break
		for batasKanan < n-1 && (*A)[batasKanan+1].ukuran == x {
			batasKanan++
		}
		
		cetakData(A, batasKiri, batasKanan)
		
	} else {
		fmt.Println("\nData dengan ukuran", x, "tidak ditemukan.\n")
	}
}

//ukuran binary search desc
func ukuranBinarySearchDesc(A *datapakaian, n int, x string){
	var left, right, mid int 
	var idx int 
	var batasKanan, batasKiri int
	
	idx = -1 
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right)/2
		if (*A)[mid].ukuran == x {
			idx = mid 
		} else if bobotUkuran(x) < bobotUkuran((*A)[mid].ukuran) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if idx != -1 {
		
		batasKiri = idx 
		batasKanan = idx
		
		// Cari tahu seberapa jauh data yang sama di sebelah kiri tanpa break
		for batasKiri > 0 && (*A)[batasKiri-1].ukuran == x {
			batasKiri--
		}
		
		// Cari tahu seberapa jauh data yang sama di sebelah kanan tanpa break
		for batasKanan < n-1 && (*A)[batasKanan+1].ukuran == x {
			batasKanan++
		}
		
		cetakData(A, batasKiri, batasKanan)
		
	} else {
		fmt.Println("\nData dengan ukuran", x, "tidak ditemukan.\n")
	}
}

//warna binary search asc 
func warnaBinarySearchAsc(A *datapakaian, n int, x string){
	var left, right, mid int 
	var idx int 
	var batasKanan, batasKiri int
	
	idx = -1 
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right)/2
		if (*A)[mid].warna == x {
			idx = mid 
		} else if x > (*A)[mid].warna {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if idx != -1 {
		
		batasKiri = idx 
		batasKanan = idx
		
		// Cari tahu seberapa jauh data yang sama di sebelah kiri tanpa break
		for batasKiri > 0 && (*A)[batasKiri-1].warna == x {
			batasKiri--
		}
		
		// Cari tahu seberapa jauh data yang sama di sebelah kanan tanpa break
		for batasKanan < n-1 && (*A)[batasKanan+1].warna == x {
			batasKanan++
		}
		
		cetakData(A, batasKiri, batasKanan)
		
	} else {
		fmt.Println("\nData dengan ukuran", x, "tidak ditemukan.\n")
	}
}

//warna binary search desc
func warnaBinarySearchDesc(A *datapakaian, n int, x string){
	var left, right, mid int 
	var idx int 
	var batasKanan, batasKiri int
	
	idx = -1 
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right)/2
		if (*A)[mid].warna == x {
			idx = mid 
		} else if x < (*A)[mid].warna {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if idx != -1 {
		
		batasKiri = idx 
		batasKanan = idx
		
		// Cari tahu seberapa jauh data yang sama di sebelah kiri tanpa break
		for batasKiri > 0 && (*A)[batasKiri-1].warna == x {
			batasKiri--
		}
		
		// Cari tahu seberapa jauh data yang sama di sebelah kanan tanpa break
		for batasKanan < n-1 && (*A)[batasKanan+1].warna == x {
			batasKanan++
		}
		
		cetakData(A, batasKiri, batasKanan)
		
	} else {
		fmt.Println("\nData dengan ukuran", x, "tidak ditemukan.\n")
	}
}

//menu sorting 
func menuSorting(){
	var pilih, pilih2, pilih3, pilih4, pilih21 int
	
	clearScreen()
	
	fmt.Println("+------------------------------------------+")
	fmt.Println("|              Anda Berada Di              |")
	fmt.Println("|           Menu Sorting Pakaian           |")
	fmt.Println("+------------------------------------------+")
	fmt.Printf("| %-40s |\n", "[1] Selection Sort") 
	fmt.Printf("| %-40s |\n", "[2] Insertion Sort") 
	fmt.Printf("| %-40s |\n", "[0] Menu Utama") 
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilih [1/2/0]? ")
	fmt.Scan(&pilih)
	
	if pilih == 1 {
		clearScreen() 
		fmt.Println("+------------------------------------------+")
		fmt.Println("|              Anda Berada Di              |")
		fmt.Println("|            Menu Selection Sort           |") 
		fmt.Println("+------------------------------------------+")
		fmt.Printf("| %-40s |\n", "[1] Ukuran") 
		fmt.Printf("| %-40s |\n", "[2] Warna") 
		fmt.Printf("| %-40s |\n", "[0] Menu Sorting") 
		fmt.Println("+------------------------------------------+")
		fmt.Print("Pilih [1/2/0]? ")
		fmt.Scan(&pilih2)
		
		if pilih2 == 1 {	
			clearScreen()
			fmt.Println("+------------------------------------------+")
			fmt.Println("|              Anda Berada Di              |")
			fmt.Println("|        Menu Ukuran Selection Sort        |") 
			fmt.Println("+------------------------------------------+")
			fmt.Printf("| %-40s |\n", "[1] Asending") 
			fmt.Printf("| %-40s |\n", "[2] Decending") 
			fmt.Printf("| %-40s |\n", "[0] Menu Sorting") 
			fmt.Println("+------------------------------------------+")
			fmt.Print("Pilih [1/2/0]? ")
			fmt.Scan(&pilih21)
			
			if pilih21 == 1 {
				ukuranSelecSortAsc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih21 == 2{
				ukuranSelecSortDesc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih21 == 0 {
				clearScreen()
				menuSorting()
			}
			
		} else if pilih2 == 2 {
			clearScreen()
			fmt.Println("+------------------------------------------+")
			fmt.Println("|              Anda Berada Di              |")
			fmt.Println("|        Menu Warna Selection Sort        |") 
			fmt.Println("+------------------------------------------+")
			fmt.Printf("| %-40s |\n", "[1] Asending") 
			fmt.Printf("| %-40s |\n", "[2] Decending") 
			fmt.Printf("| %-40s |\n", "[0] Menu Sorting") 
			fmt.Println("+------------------------------------------+")
			fmt.Print("Pilih [1/2/0]? ")
			fmt.Scan(&pilih21)
			
			if pilih21 == 1 {
				warnaSelecSortAsc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih21 == 2{
				warnaSelecSortDesc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih21 == 0 {
				clearScreen()
				menuSorting()
			}
		}
	} else if pilih == 2 {
		clearScreen() 
		fmt.Println("+------------------------------------------+")
		fmt.Println("|              Anda Berada Di              |")
		fmt.Println("|            Menu Insertion Sort           |")
		fmt.Println("+------------------------------------------+")
		fmt.Printf("| %-40s |\n", "[1] Ukuran") 
		fmt.Printf("| %-40s |\n", "[2] Warna") 
		fmt.Printf("| %-40s |\n", "[0] Menu Searching") 
		fmt.Println("+------------------------------------------+")
		fmt.Print("Pilih [1/2/0]? ")
		fmt.Scan(&pilih2)
		
		if pilih2 == 1 {
			clearScreen() 
			fmt.Println("+------------------------------------------+")
			fmt.Println("|              Anda Berada Di              |")
			fmt.Println("|         Menu Ukuran Insertion Sort       |")
			fmt.Println("+------------------------------------------+")
			fmt.Printf("| %-40s |\n", "[1] Ascending") 
			fmt.Printf("| %-40s |\n", "[2] Descending") 
			fmt.Printf("| %-40s |\n", "[0] Menu Sorting") 
			fmt.Println("+------------------------------------------+")
			fmt.Print("Pilih [1/2/0]? ")
			fmt.Scan(&pilih3)
			
			if pilih3 == 1 {
				ukuranInsertAsc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih3 == 2 {
				ukuranInsertDesc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih3 == 0 {
				clearScreen()
				menuSorting()
			}
			
		} else if pilih2 == 2 {
			clearScreen() 
			fmt.Println("+------------------------------------------+")
			fmt.Println("|              Anda Berada Di              |")
			fmt.Println("|         Menu Warna Insertion Sort        |")
			fmt.Println("+------------------------------------------+")
			fmt.Printf("| %-40s |\n", "[1] Ascending") 
			fmt.Printf("| %-40s |\n", "[2] Descending") 
			fmt.Printf("| %-40s |\n", "[0] Menu Sorting") 
			fmt.Println("+------------------------------------------+")
			fmt.Print("Pilih [1/2/0]? ")
			fmt.Scan(&pilih4)
			
			if pilih4 == 1 {
				warnaInsertAsc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih4 == 2 {
				warnaInsertDesc(&daftarToko, jumlahData)
				cetakData(daftarToko, 0, jumlahData - 1)
			} else if pilih4 == 0 {
				clearScreen()
				menuSorting()
			}
		} else if pilih2 == 0 {
			clearScreen()
			menuSorting()
		}
	} else if pilih == 0 {
		clearScreen()
		menu_utama()
	}
}

//Muti
//ukuran selection sort asc 
func ukuranSelecSortAsc(A *datapakaian, n int){
	var i, idx, pass int 
	var temp pakaian 
	
	pass = 0
	
	for pass < n - 1 {
		idx = pass
		i = pass + 1
		
		for i < n {
			//biar terurut XS-S-M-L-XL
			if bobotUkuran((*A)[i].ukuran) < bobotUkuran((*A)[idx].ukuran) {
				idx = i 
			}
			i = i + 1
		}
		//swap
		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp 
		
		pass = pass + 1
	}
}

//ukuran selection sort desc 
func ukuranSelecSortDesc(A *datapakaian, n int){
	var i, idx, pass int 
	var temp pakaian 
	
	pass = 0
	
	for pass < n - 1 {
		idx = pass
		i = pass + 1
		
		for i < n {
			//biar terurut XS-S-M-L-XL
			if bobotUkuran((*A)[i].ukuran) > bobotUkuran((*A)[idx].ukuran) {
				idx = i 
			}
			i = i + 1
		}
		//swap
		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp 
		
		pass = pass + 1
	}
}

//warna selection sort asc 
func warnaSelecSortAsc(A *datapakaian, n int){
	var i, idx, pass int 
	var temp pakaian 
	
	pass = 0
	
	for pass < n - 1 {
		idx = pass
		i = pass + 1
		
		for i < n {
			
			if (*A)[i].warna < (*A)[idx].warna {
				idx = i 
			}
			i = i + 1
		}
		//swap
		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp 
		
		pass = pass + 1
	}
}

//warna selection desc  
func warnaSelecSortDesc(A *datapakaian, n int){
	var i, idx, pass int 
	var temp pakaian 
	
	pass = 0
	
	for pass < n - 1 {
		idx = pass
		i = pass + 1
		
		for i < n {
			
			if (*A)[i].warna > (*A)[idx].warna {
				idx = i 
			}
			i = i + 1
		}
		//swap
		temp = (*A)[pass]
		(*A)[pass] = (*A)[idx]
		(*A)[idx] = temp 
		
		pass = pass + 1
	}
}

//Kay
//ukuran insertion sort asc 
func ukuranInsertAsc(A *datapakaian, n int){
	var i, pass int
	var temp pakaian
	
	pass = 1
	
	for pass < n {
		temp = (*A)[pass]
		i = pass - 1
		
		for i >= 0 && bobotUkuran((*A)[i].ukuran) > bobotUkuran(temp.ukuran) {
			(*A)[i+1] = (*A)[i]
			i--
		}
		
		(*A)[i+1] = temp
		pass++
		
	}
}

//ukuran insertion sort desc
func ukuranInsertDesc(A *datapakaian, n int){
	var i, pass int
	var temp pakaian
	
	pass = 1
	
	for pass < n {
		temp = (*A)[pass]
		i = pass - 1
		
		for i >= 0 && bobotUkuran((*A)[i].ukuran) < bobotUkuran(temp.ukuran) {
			(*A)[i+1] = (*A)[i]
			i--
		}
		
		(*A)[i+1] = temp
		pass++
	}
}

//warna insertion asc 
func warnaInsertAsc(A *datapakaian, n int){
	var i, pass int
	var temp pakaian
	
	pass = 1
	
	for pass < n {
		temp = (*A)[pass]
		i = pass - 1
		
		for i >= 0 && (*A)[i].warna > temp.warna {
			(*A)[i+1] = (*A)[i]
			i--
		}
		(*A)[i+1] = temp
		pass++
	}
}

//warna insertion desc 
func warnaInsertDesc(A *datapakaian, n int){
	var i, pass int
	var temp pakaian
	
	pass = 1
	
	for pass < n {
		temp = (*A)[pass]
		i = pass - 1
		
		for i >= 0 && (*A)[i].warna < temp.warna {
			(*A)[i+1] = (*A)[i]
			i--
		}
		(*A)[i+1] = temp
		pass++
	}
}
}


//fungsi clear screen
func clearScreen(){
	
}

//menu exit 
func menuexit(){
	
}

func main(){
	var pilih int

	inisialisasiData()
	
	for {
		menu_utama()
		fmt.Scan(&pilih)
		switch pilih {
			case 1 : daftarpakaian()
			case 2 : menu_tambahdata()
			case 3 : menu_editdata()
			case 4 : menu_hapusdata()
			case 5 : menu_searching()
			case 6 : menuSorting()
			case 0 : menuexit()
		}
	}
}
