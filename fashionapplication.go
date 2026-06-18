package main

import (
	"fmt"
)

// mencari produk berdasarkan ukuran / warna

const DATAMAX int = 1000

type pakaian struct {
	id, stok            int
	nama, warna, ukuran string
}
type datapakaian [DATAMAX]pakaian

var daftarToko datapakaian
var jumlahData int = 0

func main() {
	var pilih int
	var konfirmasiExit bool

	icon()

	inisialisasiData()

	for {
		menu_utama()
		pilih = inputInt()
		switch pilih {
		case 1:
			daftarpakaian()
		case 2:
			menutambahData()
		case 3:
			menuEditData()
		case 4:
			menuhapusData()
		case 5:
			menu_searching()
		case 6:
			menuSorting()
		case 7:
			menu_rekomendasi()
		case 0:
			konfirmasiExit = menuExit() //tanya user
			if konfirmasiExit {
				return //kalo true, aplikasi bakalan keluar
			}
			//kalo false, aplikasi bakalan balik ke menu utama
		default:
			fmt.Println("Pilihan tidak valid, silahkan coba lagi")
		}
	}
}

func icon() {
	/*
		I.S.: Layar terminal siap menerima tampilan baru.
		F.S.: Layar terminal menampilkan logo raksasa "KAYRA" di dalam kotak tabel,
		dan program dalam kondisi tertahan (jeda) sampai user mengetik "ENTER".
	*/
	var jeda int

	clearScreen()

	fmt.Println("+-----------------------------------------------------+")
	fmt.Println("|                     WELCOME TO                      |")
	fmt.Println("+-----------------------------------------------------+")
	fmt.Printf("| %-51s |\n", " ")
	fmt.Printf("| %-51s |\n", "##    ##     ###    ##     ## ########     ###    ")
	fmt.Printf("| %-51s |\n", "##   ##     ## ##    ##   ##  ##     ##   ## ##   ")
	fmt.Printf("| %-51s |\n", "##  ##     ##   ##    ####    ##     ##  ##   ##  ")
	fmt.Printf("| %-51s |\n", "#####     ##     ##    ##     ########  ##     ## ")
	fmt.Printf("| %-51s |\n", "##  ##    #########    ##     ##   ##   ######### ")
	fmt.Printf("| %-51s |\n", "##   ##   ##     ##    ##     ##    ##  ##     ## ")
	fmt.Printf("| %-51s |\n", "##    ##  ##     ##    ##     ##     ## ##     ## ")
	fmt.Printf("| %-51s |\n", " ")
	fmt.Println("+-----------------------------------------------------+")
	fmt.Println()
	fmt.Println("          [ Ketik 1 Untuk Masuk Aplikasi ]        ")

	fmt.Scanln(&jeda)
	for jeda != 1 {
		//kalo inputnya selain 1 dia akan mengulang kembali
		fmt.Println("Input tidak valid silahkan coba lagi!")
		fmt.Scanln(&jeda)
	}
	clearScreen()
}

func menu_utama() {
	/*
		I.S.: Layar terminal dalam keadaan bersih.
		F.S.: Menampilkan pilihan menu utama [1 - 7] beserta
		opsi [0] Exit ke layar terminal, siap menunggu input pilihan dari user.
	*/

	fmt.Println("+--------------------------------------+")
	fmt.Println("|          Selamat Datang Di           |")
	fmt.Println("|       Aplikasi Manajemen Fashion     |")
	fmt.Println("+--------------------------------------+")
	fmt.Printf("| %-36s |\n", "[1] Daftar Data Pakaian")
	fmt.Printf("| %-36s |\n", "[2] Tambah Data Pakaian")
	fmt.Printf("| %-36s |\n", "[3] Edit Data Pakaian")
	fmt.Printf("| %-36s |\n", "[4] Hapus Data Pakaian")
	fmt.Printf("| %-36s |\n", "[5] Search Data Pakaian")
	fmt.Printf("| %-36s |\n", "[6] Sortir Data Pakaian")
	fmt.Printf("| %-36s |\n", "[7] Rekomendasi Fashion")
	fmt.Printf("| %-36s |\n", "[0] Exit")
	fmt.Println("+--------------------------------------+")
	fmt.Print("Pilih [0 - 7]?")
}
func inputInt() int {
	/*
		fungsi mengembalikan input integer yang sudah divalidasi.
		Input harus berupa angka dan tidak boleh negatif.
		Jika input salah, pengguna diminta mengulang.
	*/
	var err error
	var value int

	for {
		_, err = fmt.Scan(&value)

		if err != nil {
			fmt.Println("Input harus berupa angka!")
			var dummy string
			fmt.Scanln(&dummy)

		} else if value < 0 {
			fmt.Println("Input tidak boleh negatif!")

		} else {
			return value
		}
	}
}

// cetak Data
func cetakData(A datapakaian, awal int, akhir int) {
	/*
		I.S.: Terdefinisi sebuah array A yang berisi data pakaian, serta batas indeks awal dan akhir.
		F.S.: Data pakaian yang berada di antara indeks awal sampai akhir tercetak rapi ke layar dalam format tabel.
		Array A tidak mengalami perubahan posisi atau isi data.
	*/
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

// menu tambah data
func menutambahData() {
	/*
		IS : Menampilkan menu untuk tambah data.
		FS : Menampilkan opsi untuk memulai penambahan. Jika user input 1 (true), maka
		sistem akan melanjutkan eksekusi dengan memanggil prosedur tambahPakaian.
		Jika user input 0 (false) maka sistem akan mengembalikan ke menu utama.
	*/

	clearScreen()
	fmt.Println("+------------------------------------------+")
	fmt.Println("|              Anda Berada Di              |")
	fmt.Println("|         Menu Tambah Data Pakaian         |")
	fmt.Println("+------------------------------------------+")
	fmt.Printf("| %-40s |\n", "[1] Tambah Data")
	fmt.Printf("| %-40s |\n", "[0] Menu Utama")
	fmt.Println("+------------------------------------------+")

	var pilih int
	fmt.Print("Pilih [0 / 1]? ")
	pilih = inputInt()

	if pilih == 1 {
		// Langsung panggil fungsi gabungan kita
		tambahPakaian(&daftarToko)
	} else if pilih != 0 && pilih != 1 {
		fmt.Println("Pilihan tidak valid silahkan coba lagi!")
		menutambahData()
	}
	// Jika pilih 0, otomatis keluar dari fungsi ini dan balik ke main loop
}

func tambahPakaian(data *datapakaian) {
	/*
		IS : Terdefinisi alamat memori pointer array data yang menampung rekaman data pakaian saat ini,
		serta variabel global jumlahData yang menyimpan total data yang sudah tersimpan sebelum
		operasi penambahan dilakukan.
		FS : User menginput elemen baru sebanyak banyakData. Setiap data baru akan mendapatkan ID unik yang
		berurutan dan nilai variabel jumlahData akan bertambah otomatis sebanyak data yang berhasil ditambah.
	*/

	clearScreen()
	fmt.Println("========================================")
	fmt.Println("          FITUR TAMBAH PAKAIAN          ")
	fmt.Println("========================================")

	var banyakData, i, inputStok int
	var jeda int

	fmt.Print("Mau menambah berapa data pakaian? ")
	fmt.Scan(&banyakData)

	// Validasi jika user iseng menginput angka 0 atau minus
	if banyakData <= 0 {
		fmt.Println("\nJumlah data tidak valid. Batal menambah data.")
		return
	}

	// Melakukan perulangan sebanyak jumlah data yang diinginkan user
	for i = 0; i < banyakData; i++ {
		fmt.Printf("\n--- Menginput Data ke-%d dari %d ---\n", i+1, banyakData)

		// Otomatis generate ID berdasarkan jumlah data saat ini
		data[jumlahData].id = jumlahData + 1

		fmt.Print("Masukkan Nama Pakaian  : ")
		fmt.Scan(&data[jumlahData].nama)

		fmt.Print("Masukkan Warna  : ")
		fmt.Scan(&data[jumlahData].warna)

		fmt.Print("Masukkan Ukuran : ")
		fmt.Scan(&data[jumlahData].ukuran)

		fmt.Print("Masukkan Stok   : ")
		inputStok = inputInt()
		data[jumlahData].stok = inputStok

		// Menaikkan jumlahData global setiap kali 1 baju sukses diinput
		jumlahData++
	}

	fmt.Println("\n========================================")
	fmt.Printf(" Berhasil menambahkan %d data pakaian!\n", banyakData)
	fmt.Println("========================================")

	// Beri jeda Enter agar user bisa melihat pesan suksesnya
	fmt.Print("Ketik 1 untuk kembali...")
	fmt.Scan(&jeda)

	clearScreen()
}

// menu edit data
func menuEditData() {
	/*
		IS : Program berada pada menu utama atau menu sebelumnya, dan data pakaian yang tersimpan
		     dapat diakses untuk dilakukan proses pengeditan.

		FS : Menu Edit Data ditampilkan. Jika user memilih opsi 1 maka program memanggil
		     prosedur editDatabyId untuk mengubah data pakaian berdasarkan ID.
		     Jika user memilih opsi 0 maka tidak dilakukan proses edit dan program kembali
		     ke alur menu sebelumnya.
	*/

	var pilih int

	clearScreen()
	fmt.Println("+------------------------------------------+")
	fmt.Println("|              Anda Berada Di              |")
	fmt.Println("|           Menu Edit Data Pakaian         |")
	fmt.Println("+------------------------------------------+")
	fmt.Printf("| %-40s |\n", "[1] Edit Data")
	fmt.Printf("| %-40s |\n", "[0] Menu Utama")
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilih [0 / 1]? ")
	pilih = inputInt()
	if pilih == 1 {
		editDatabyId(&daftarToko)
	} else if pilih != 0 && pilih != 1 {
		fmt.Println("Pilihan tidak valid silahkan coba lagi!")
		menuEditData()
	}
}

// fungsi untuk mengedit data berdasarkan id
func editDatabyId(data *datapakaian) {
	/*
		IS : Menampilkan menu edit data pakaian yang tersedia dalam sistem.
		FS : Menampilkan opsi untuk melakukan pengeditan data. Jika user input 1 (true),
		     maka sistem akan memanggil prosedur editDatabyId untuk mengubah data
		     pakaian berdasarkan ID. Jika user input 0 (false), maka sistem akan
		     kembali ke menu utama tanpa melakukan perubahan data.
	*/

	var id, stok int
	var nama, warna, ukuran string
	fmt.Print("Masukkan Id     : ")
	fmt.Scan(&id)
	for !findId(data, id) {
		fmt.Println("Not Found!")
		fmt.Print("Masukkan Id   : ")
		fmt.Scan(&id)
	}
	fmt.Print("Masukkan Nama   : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Stok   : ")
	fmt.Scan(&stok)
	fmt.Print("Masukkan Warna  : ")
	fmt.Scan(&warna)
	fmt.Print("Masukkan Ukuran : ")
	fmt.Scan(&ukuran)
	data[id-1].nama = nama
	data[id-1].stok = stok
	data[id-1].warna = warna
	data[id-1].ukuran = ukuran
}

// fungsi untuk mencari id
func findId(data *datapakaian, id int) bool {
	/*
		Mencari data pakaian berdasarkan ID.
		Mengembalikan posisi data jika ditemukan, dan -1 jika tidak ditemukan.

	*/

	var left, right, mid int
	var found bool
	found = false
	left = 0
	right = jumlahData - 1
	for left <= right && !found {
		mid = (left + right) / 2
		if data[mid].id == id {
			found = true
		} else if data[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}

// fungsi untuk mencari indeks berdasarkan id
func findIdxbyId(data *datapakaian, id int) int {
	/*
		Mencari indeks data pakaian berdasarkan ID.
		Mengembalikan indeks jika ditemukan, dan -1 jika tidak ditemukan.
	*/

	var left, right, mid int
	var idxfound int
	idxfound = -1
	left = 0
	right = jumlahData - 1
	for left <= right && idxfound == -1 {
		mid = (left + right) / 2
		if (*data)[mid].id == id {
			idxfound = mid
		} else if (*data)[mid].id > id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idxfound
}

// menu hapus data
func menuhapusData() {
	/*
		IS : Menampilkan menu utama untuk menghapus data.
		FS : Menampilkan opsi untuk memulai penghapusan. Jika user input 1 (true), maka
		sistem akan melanjutkan eksekusi dengan memanggil prosedur hapusDatabyID.
		Jika user input 0 (false) maka sistem akan mengembalikan ke menu utama.
	*/

	var pilih int

	clearScreen()
	fmt.Println("+------------------------------------------+")
	fmt.Println("|              Anda Berada Di              |")
	fmt.Println("|          Menu Hapus Data Pakaian         |")
	fmt.Println("+------------------------------------------+")
	fmt.Printf("| %-40s |\n", "[1] Hapus Data")
	fmt.Printf("| %-40s |\n", "[0] Menu Utama")
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilih [1 / 0]? ")
	pilih = inputInt()

	if pilih == 1 {
		hapusDatabyID(&daftarToko, &jumlahData)
	} else if pilih != 0 && pilih != 1 {
		fmt.Println("Pilihan tidak valid silahkan coba lagi!")
		menuhapusData()
	}
}

func hapusDatabyID(A *datapakaian, n *int) {
	/*
		IS : Terdefinisi array A dan banyaknya data pakaian sebagai n.
		FS : Jika ID target yang diinput user ditemukan di dalam array A dan user mengonfirmasi
		dengan mengetik "iya", maka elemen data pada posisi tersebut akan dihapus,
		seluruh elemen di sebelah kanannya akan digeser satu langkah ke kiri untuk mengisi kekosongan,
		dan nilai penunjuk jumlah data (*n) otomatis dikurangi 1. Jika ID tidak ditemukan atau user membatalkan
		dengan mengetik "tidak".
	*/

	var idTarget, idx, jeda int
	var opsi string

	clearScreen()
	fmt.Println("==========================================")
	fmt.Println("           FITUR HAPUS PAKAIAN            ")
	fmt.Println("==========================================")

	// Input ID yang mau dieksekusi
	fmt.Print("Masukkan ID pakaian yang mau dihapus: ")
	fmt.Scan(&idTarget)

	// Validasi dulu, ID-nya beneran ada gak di array?
	idx = findIdxbyId(A, idTarget)
	if idx == -1 {
		fmt.Printf("\n[Error] Pakaian dengan ID %d tidak ditemukan!\n", idTarget)

		fmt.Print("[Tekan 1 untuk kembali...]")
		fmt.Scan(&jeda)

		menuhapusData()
	}

	// Jika ID ada, baru tampilkan info singkat baju tersebut & konfirmasi
	fmt.Println("\n------------------------------------------")
	fmt.Println("Data ditemukan:")
	fmt.Printf("Nama  : %s\n", (*A)[idx].nama)
	fmt.Printf("Warna : %s\n", (*A)[idx].warna)
	fmt.Printf("Ukuran: %s\n", (*A)[idx].ukuran)
	fmt.Println("------------------------------------------")

	fmt.Println("Apakah anda yakin menghapus data tersebut?")
	fmt.Print("[Ketik iya / tidak]: ")
	fmt.Scan(&opsi)

	if opsi == "iya" {
		if idx == *n-1 {
			// Kasus A: Jika data yang dihapus kebetulan ada di paling ujung belakang array,
			// kita cukup kurangi saja jumlah data totalnya sebanyak 1.
			*n = *n - 1
		} else {
			// Kasus B: Jika data ada di tengah-tengah, geser semua data kanan ke kiri
			i := idx
			for i < *n-1 {
				(*A)[i] = (*A)[i+1]
				i++
			}
			*n = *n - 1 // Kurangi jumlah elemen setelah digeser
		}

		fmt.Println("\n==========================================")
		fmt.Println("   SUKSES: Data pakaian berhasil dihapus! ")
		fmt.Println("==========================================")
	} else {
		fmt.Println("\nPenghapusan data dibatalkan.")
	}

	// Beri jeda enter agar user sempat membaca status sukses/gagalnya
	fmt.Print("\nTekan 1 untuk kembali...")
	fmt.Scan(&jeda)

	clearScreen()
}

func daftarpakaian() {
	/*
		IS : Tersedia data pakaian yang tersimpan dalam array daftarToko dengan jumlah
		     data yang tercatat pada variabel jumlahData.
		FS : Seluruh data pakaian yang tersimpan ditampilkan ke layar dalam bentuk
		     tabel yang berisi ID, nama pakaian, warna, ukuran, dan stok.
	*/

	var i int

	clearScreen()

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

// menu rekomendasi
func menu_rekomendasi() {
	/* Sub program menu_rekomendasi merupakan sub program menu rekomendasi fashion yang dicari berdasarkan
	   kondisi cuaca, warna favorit user, dan acara yang akan dihadiri oleh user.

	   BIKIN MENU : cuasa, warna, acara, gender
	   setelah itu user akan nge input berdasarkan yang ada di menu.
	*/
	var pilih int

	fmt.Println("+------------------------------------------+")
	fmt.Println("|              Anda Berada Di              |")
	fmt.Println("|         Menu Rekomendasi Pakaian         |")
	fmt.Println("+------------------------------------------+")
	fmt.Printf("| %-40s |\n", "[1] Cari Rekomendasi Outfit")
	fmt.Printf("| %-40s |\n", "[0] Menu Utama")
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilih [1 / 0]? ")
	pilih = inputInt()

	if pilih == 1 {
		pilih_rekomendasiPakaian()
	} else if pilih == 0 {
		clearScreen()
	} else {
		fmt.Println("Pilihan tidak valid silahkan coba lagi!")
		menu_rekomendasi()
	}
}
func pilih_rekomendasiPakaian() {
	var pilihanCuaca, pilihanAcara int
	var cuaca, acara string

	fmt.Println("+------------------------------------------+")
	fmt.Println("|              PILIH CUACA                 |")
	fmt.Println("+------------------------------------------+")
	fmt.Println("| 1. Cerah                                 |")
	fmt.Println("| 2. Mendung                               |")
	fmt.Println("| 3. Hujan                                 |")
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihanCuaca)

	switch pilihanCuaca {
	case 1:
		cuaca = "cerah"
	case 2:
		cuaca = "mendung"
	case 3:
		cuaca = "hujan"
	default:
		fmt.Println("Pilihan cuaca tidak valid")
		return
	}

	fmt.Println()

	fmt.Println("+------------------------------------------+")
	fmt.Println("|              PILIH ACARA                 |")
	fmt.Println("+------------------------------------------+")
	fmt.Println("| 1. Formal                                |")
	fmt.Println("| 2. Semi Formal                           |")
	fmt.Println("| 3. Casual                                |")
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihanAcara)

	switch pilihanAcara {
	case 1:
		acara = "formal"
	case 2:
		acara = "semiformal"
	case 3:
		acara = "casual"
	default:
		fmt.Println("Pilihan acara tidak valid")
		return
	}
	rekomendasiPakaian(cuaca, acara)
}
func rekomendasiPakaian(cuaca string, acara string) {
	var i int
	fmt.Println()
	fmt.Println("+------------------------------------------+")
	fmt.Println("|         REKOMENDASI PAKAIAN              |")
	fmt.Println("+------------------------------------------+")

	for i = 0; i < len(daftarToko); i++ {

		switch {

		// ================= CERAH =================

		case cuaca == "cerah" && acara == "formal":
			if daftarToko[i].id == 7 || // Kemeja Kerja
				daftarToko[i].id == 13 || // Kemeja Batik
				daftarToko[i].id == 27 || // Kemeja Formal
				daftarToko[i].id == 42 || // Kemeja Linen
				daftarToko[i].id == 49 || // Blazer Formal
				daftarToko[i].id == 52 || // Kemeja Kerja
				daftarToko[i].id == 57 || // Kemeja Shanghai
				daftarToko[i].id == 78 || // Kemeja Oxford
				daftarToko[i].id == 82 || // Kemeja Formal
				daftarToko[i].id == 87 || // Kemeja Batik
				daftarToko[i].id == 92 { // Kemeja Pria
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		case cuaca == "cerah" && acara == "semiformal":
			if daftarToko[i].id == 2 || // Kemeja Flanel
				daftarToko[i].id == 3 || // Celana Chino
				daftarToko[i].id == 16 || // Cardigan Knitted
				daftarToko[i].id == 21 || // Kaos Polo
				daftarToko[i].id == 31 || // Rok Plisket
				daftarToko[i].id == 55 || // Cardigan Polos
				daftarToko[i].id == 65 || // Sweater Rajut
				daftarToko[i].id == 85 { // Sweater Vneck
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		case cuaca == "cerah" && acara == "casual":
			if daftarToko[i].id == 1 || // Kaos Polos
				daftarToko[i].id == 6 || // Kaos Polos
				daftarToko[i].id == 8 || // Celana Jeans
				daftarToko[i].id == 17 || // Kaos V-Neck
				daftarToko[i].id == 25 || // Sweater Hoodie
				daftarToko[i].id == 26 || // Kaos Polos
				daftarToko[i].id == 37 || // Kaos Oversize
				daftarToko[i].id == 41 || // Kaos Polos
				daftarToko[i].id == 46 || // Kaos Grafik
				daftarToko[i].id == 51 || // Kaos Singlet
				daftarToko[i].id == 56 || // Kaos Polos
				daftarToko[i].id == 61 || // Kaos Raglan
				daftarToko[i].id == 66 || // Kaos Polo
				daftarToko[i].id == 72 || // Kaos V-Neck
				daftarToko[i].id == 77 || // Kaos Polos
				daftarToko[i].id == 81 || // Tshirt Stripe
				daftarToko[i].id == 86 || // Kaos Oversize
				daftarToko[i].id == 91 || // Kaos Polos
				daftarToko[i].id == 96 { // Kaos Distro
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		// ================= MENDUNG =================

		case cuaca == "mendung" && acara == "formal":
			if daftarToko[i].id == 30 || // Blouse Silk
				daftarToko[i].id == 45 || // Turtleneck
				daftarToko[i].id == 95 { // Cardigan Rajut
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		case cuaca == "mendung" && acara == "semiformal":
			if daftarToko[i].id == 5 || // Sweater Crewneck
				daftarToko[i].id == 18 || // Tunik Dress
				daftarToko[i].id == 59 || // Fleece Jacket
				daftarToko[i].id == 60 || // Crop Top
				daftarToko[i].id == 90 { // Tunik Polos
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		case cuaca == "mendung" && acara == "casual":
			if daftarToko[i].id == 9 || // Hoodie Oversize
				daftarToko[i].id == 24 || // Jaket Denim
				daftarToko[i].id == 35 || // Windbreaker
				daftarToko[i].id == 69 || // Track Jacket
				daftarToko[i].id == 76 || // Hoodie Polos
				daftarToko[i].id == 99 { // Jaket Hoodie
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		// ================= HUJAN =================

		case cuaca == "hujan" && acara == "formal":
			if daftarToko[i].id == 15 || // Jaket Parka
				daftarToko[i].id == 64 || // Puffer Jacket
				daftarToko[i].id == 75 || // Biker Jacket
				daftarToko[i].id == 80 { // Long Coat
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		case cuaca == "hujan" && acara == "semiformal":
			if daftarToko[i].id == 4 || // Jaket Bomber
				daftarToko[i].id == 29 || // Coach Jacket
				daftarToko[i].id == 44 || // Varsity Jacket
				daftarToko[i].id == 54 || // Anorak Jacket
				daftarToko[i].id == 89 || // Bomber Jacket
				daftarToko[i].id == 94 { // Harrington
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}

		case cuaca == "hujan" && acara == "casual":
			if daftarToko[i].id == 84 || // Raincoat
				daftarToko[i].id == 88 || // Celana Kulot
				daftarToko[i].id == 93 || // Celana Sirwal
				daftarToko[i].id == 98 || // Celana Tactical
				daftarToko[i].id == 100 { // Blouse Casual
				fmt.Printf("|  %-40s|\n", daftarToko[i].nama)
			}
		}
	}

	fmt.Println("+------------------------------------------+")
}

// ubah huruf jadi abjad
func bobotUkuran(ukuran string) int {
	/*
		Inputan sebuah teks yang merepresentasikan ukuran pakaian (misal: "S", "M", "L").
		Memetakan inputan ukuran tersebut ke dalam skala angka agar bisa dibandingkan secara matematis.
		Dan mengembalikan angka integer dari 1 sampai 7 sesuai tingkatan ukuran baju
		(XS=1, S=2, dst). Jika ukuran tidak dikenali, mengembalikan angka 0.
	*/

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

// menu searching
func menu_searching() {
	/*
		IS : User memilih menu pencarian dari menu utama. Layar terminal dibersihkan.
		FS : Menampilkan pilihan menu algoritma pencarian (Sequential/Binary) dan
		kategori pencarian (Ukuran/Warna), siap menerima input pilihan dari user.
	*/
	var pilih, pilih2, pilih3 int
	var katakunci int
	var katakunci2 string

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
	pilih = inputInt()

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
			fmt.Scan(&katakunci2)
			sequentialSearchbySize(daftarToko, jumlahData, katakunci2)

		} else if pilih2 == 2 {
			fmt.Print("Warna yang dicari: ")
			fmt.Scan(&katakunci2)
			sequentialSearchbyColor(daftarToko, jumlahData, katakunci2)
		} else {
			clearScreen()
			menu_searching()
		}

	} else if pilih == 2 {
		clearScreen()
		fmt.Println("+------------------------------------------+")
		fmt.Println("|              Anda Berada Di              |")
		fmt.Println("|         Menu ID Binary Search        |")
		fmt.Println("+------------------------------------------+")
		fmt.Printf("| %-40s |\n", "[1] Ascending")
		fmt.Printf("| %-40s |\n", "[0] Menu Searching")
		fmt.Println("+------------------------------------------+")
		fmt.Print("Pilih [1/0]? ")
		fmt.Scan(&pilih3)

		if pilih3 == 1 {
			fmt.Print("ID yang dicari: ")
			fmt.Scan(&katakunci)
			idBinarySearch(&daftarToko, jumlahData, katakunci)

		} else if pilih3 == 0 {
			clearScreen()
			menu_searching()
		}
	} else if pilih == 0 {
		clearScreen()
		menu_utama()
	} else {
		fmt.Println("Pilihan tidak valid silahkan coba lagi!")
		menu_searching()
	}
}

// mencari ukuran menggunakan sequential search
func sequentialSearchbySize(data datapakaian, n int, ukuran string) {
	var i int
	i = 0

	fmt.Println("\n+-------+----------------------+-----------------+------------+-------+")
	fmt.Printf("| %-5s | %-20s | %-15s | %-10s | %-5s |\n", "ID", "Nama Pakaian", "Warna", "Ukuran", "Stok")
	fmt.Println("+-------+----------------------+-----------------+------------+-------+")
	for i < n {
		if data[i].ukuran == ukuran {
			fmt.Printf("| %-5d | %-20s | %-15s | %-10s | %-5d |\n",
				daftarToko[i].id,
				daftarToko[i].nama,
				daftarToko[i].warna,
				daftarToko[i].ukuran,
				daftarToko[i].stok)
		}
		i++
	}
	fmt.Println("+-------+----------------------+-----------------+------------+-------+\n")
}

// mencari warna menggunakan sequential search
func sequentialSearchbyColor(data datapakaian, n int, warna string) {
	var i int
	i = 0
	fmt.Println("\n+-------+----------------------+-----------------+------------+-------+")
	fmt.Printf("| %-5s | %-20s | %-15s | %-10s | %-5s |\n", "ID", "Nama Pakaian", "Warna", "Ukuran", "Stok")
	fmt.Println("+-------+----------------------+-----------------+------------+-------+")
	for i < n {
		if data[i].warna == warna {
			fmt.Printf("| %-5d | %-20s | %-15s | %-10s | %-5d |\n",
				daftarToko[i].id,
				daftarToko[i].nama,
				daftarToko[i].warna,
				daftarToko[i].ukuran,
				daftarToko[i].stok)
		}
		i++
	}
	fmt.Println("+-------+----------------------+-----------------+------------+-------+\n")
}

// Binary search by ID
func idBinarySearch(A *datapakaian, n int, x int) {
	/*
		IS : Terdefinisi sebuah array A yang berisi n data pakaian dalam keadaan sudah terurut secara naik (Ascending) berdasarkan nomor ID,
		dan terdefinisi sebuah variabel x sebagai nomor ID pakaian yang ingin dicari.
		FS : Jika nomor ID yang disimpan dalam x ditemukan di dalam array A, maka data pakaian tersebut dicetak rapi ke layar terminal (menggunakan prosedur cetakData).
		Jika nomor ID x tidak ditemukan, maka layar akan menampilkan pesan teks bahwa data tidak ditemukan.
	*/
	var left, right, mid int
	var idx int
	var batasKanan, batasKiri int

	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if (*A)[mid].id == x {
			idx = mid
		} else if x > (*A)[mid].id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if idx != -1 {

		batasKiri = idx
		batasKanan = idx

		// Cari tahu seberapa jauh data yang sama di sebelah kiri
		for batasKiri > 0 && (*A)[batasKiri-1].id == x {
			batasKiri--
		}

		// Cari tahu seberapa jauh data yang sama di sebelah kanan
		for batasKanan < n-1 && (*A)[batasKanan+1].id == x {
			batasKanan++
		}

		cetakData(daftarToko, batasKiri, batasKanan)

	} else {
		fmt.Println("Data dengan ID", x, "tidak ditemukan.")
	}
}

// menu sorting
func menuSorting() {
	/*
		IS : User memilih menu pengurutan dari menu utama. Layar terminal dibersihkan.
		FS : Menampilkan pilihan menu algoritma pengurutan (Selection/Insertion) beserta
		pilihan urutan (Ascending/Descending), siap menerima input pilihan dari user.
	*/

	var pilih, pilih2, pilih3 int

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
	pilih = inputInt()

	if pilih == 1 {
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
		pilih2 = inputInt()

		if pilih2 == 1 {
			ukuranSelecSortAsc(&daftarToko, jumlahData)
			cetakData(daftarToko, 0, jumlahData-1)
		} else if pilih2 == 2 {
			ukuranSelecSortDesc(&daftarToko, jumlahData)
			cetakData(daftarToko, 0, jumlahData-1)
		} else if pilih2 == 0 {
			clearScreen()
			menuSorting()
		}
	} else if pilih == 2 {
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
		fmt.Scan(&pilih3)

		if pilih3 == 1 {
			warnaInsertAsc(&daftarToko, jumlahData)
			cetakData(daftarToko, 0, jumlahData-1)
		} else if pilih3 == 2 {
			warnaInsertDesc(&daftarToko, jumlahData)
			cetakData(daftarToko, 0, jumlahData-1)
		} else if pilih3 == 0 {
			clearScreen()
			menuSorting()
		}
	} else if pilih == 0 {
		clearScreen()
		menu_utama()
	}
}

// Muti
// ukuran selection sort asc
func ukuranSelecSortAsc(A *datapakaian, n int) {
	/*
		IS : Terdefinisi alamat memori array A (menggunakan pointer) berisi n data pakaian yang urutan
		ukurannya masih acak.
		FS : Seluruh data pakaian di dalam array A berhasil diurutkan secara naik (Ascending) berdasarkan
		bobot ukuran pakaian (dari XS sampai XXXL).
	*/
	var i, idx, pass int
	var temp pakaian

	pass = 1

	for pass < n-1 {
		idx = pass - 1
		i = pass

		for i < n {
			//biar terurut XS-S-M-L-XL
			if bobotUkuran((*A)[i].ukuran) < bobotUkuran((*A)[idx].ukuran) {
				idx = i
			}
			i = i + 1
		}
		//swap
		temp = (*A)[pass-1]
		(*A)[pass-1] = (*A)[idx]
		(*A)[idx] = temp

		pass = pass + 1
	}
}

// ukuran selection sort desc
func ukuranSelecSortDesc(A *datapakaian, n int) {
	/*
		IS : Terdefinisi alamat memori array A (menggunakan pointer) berisi n data pakaian yang urutan
		ukurannya masih acak.
		FS : Seluruh data pakaian di dalam array A berhasil diurutkan secara menurun (Descending) berdasarkan
		bobot ukuran pakaian (dari XXXL sampai XS).
	*/
	var i, idx, pass int
	var temp pakaian

	pass = 1

	for pass < n-1 {
		idx = pass - 1
		i = pass

		for i < n {
			//biar terurut XS-S-M-L-XL
			if bobotUkuran((*A)[i].ukuran) > bobotUkuran((*A)[idx].ukuran) {
				idx = i
			}
			i = i + 1
		}
		//swap
		temp = (*A)[pass-1]
		(*A)[pass-1] = (*A)[idx]
		(*A)[idx] = temp

		pass = pass + 1
	}
}

// warna insertion asc
func warnaInsertAsc(A *datapakaian, n int) {
	/*
		IS : Terdefinisi alamat memori array A (menggunakan pointer) berisi n data pakaian yang urutan
		warnanya masih acak.
		FS : Seluruh data pakaian di dalam array A berhasil diurutkan secara naik (Ascending) berdasarkan
		warna pakaian (dari A sampai Z).
	*/
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

// warna insertion desc
func warnaInsertDesc(A *datapakaian, n int) {
	/*
		IS : Terdefinisi alamat memori array A (menggunakan pointer) berisi n data pakaian yang urutan
		warnanya masih acak.
		FS : Seluruh data pakaian di dalam array A berhasil diurutkan secara menurun (Descending) berdasarkan
		warna pakaian (dari Z sampai A).
	*/
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

func menuExit() bool {
	/*
		Function menuExit akan menerima input integer dari user di dalam fungsi.
		Fungsi ini akan menampilkan kotak konfirmasi berlapis ("KAMU YAKIN?!" dan "BENERAN?!") untuk memastikan user ingin keluar,
		serta mencetak logo "THANK YOU" jika user memilih keluar.
		Fungsi ini juga mengembalikan nilai boolean true jika user memilih "Yes" pada semua konfirmasi (sinyal aplikasi boleh ditutup),
		atau mengembalikan false jika user memilih "No" di salah satu konfirmasi (sinyal batal keluar).
	*/
	var input int

	clearScreen()

	fmt.Println("+------------------------------------------+")
	fmt.Println("|              Anda Berada Di              |")
	fmt.Println("|                Menu Exit                 |")
	fmt.Println("+------------------------------------------+")
	fmt.Printf("| %-40s |\n", "KAMU YAKIN ?!")
	fmt.Printf("| %-40s |\n", "[1] Yes")
	fmt.Printf("| %-40s |\n", "[0] No")
	fmt.Println("+------------------------------------------+")
	fmt.Print("Pilih [1 / 0]? ")
	fmt.Scan(&input)

	if input == 1 {
		clearScreen()
		fmt.Println("+------------------------------------------+")
		fmt.Println("|              Anda Berada Di              |")
		fmt.Println("|                Menu Exit                 |")
		fmt.Println("+------------------------------------------+")
		fmt.Printf("| %-40s |\n", "BENERAN ?!")
		fmt.Printf("| %-40s |\n", "[1] Yes")
		fmt.Printf("| %-40s |\n", "[0] No")
		fmt.Println("+------------------------------------------+")
		fmt.Print("Pilih [1 / 0]? ")
		fmt.Scan(&input)

		if input == 1 {
			clearScreen()

			fmt.Println("+-----------------------------------------------------------------------------------+")
			fmt.Printf("| %-79s   |\n", " ")
			fmt.Printf("| %-65s   |\n", "####### ##     ##    ###    ##    ## ##    ##    ##    ##  #######  ##     ##  ")
			fmt.Printf("| %-65s   |\n", "  ##    ##     ##   ## ##   ###   ## ##   ##      ##  ##  ##     ## ##     ##  ")
			fmt.Printf("| %-65s   |\n", "  ##    ##     ##  ##   ##  ####  ## ##  ##        ####   ##     ## ##     ##  ")
			fmt.Printf("| %-65s   |\n", "  ##    ######### ##     ## ## ## ## #####          ##    ##     ## ##     ##  ")
			fmt.Printf("| %-65s   |\n", "  ##    ##     ## ######### ##  #### ##  ##         ##    ##     ## ##     ##  ")
			fmt.Printf("| %-65s   |\n", "  ##    ##     ## ##     ## ##   ### ##   ##        ##    ##     ## ##     ##  ")
			fmt.Printf("| %-65s   |\n", "  ##    ##     ## ##     ## ##    ## ##    ##       ##     #######   #######   ")
			fmt.Printf("| %-79s   |\n", " ")
			fmt.Println("+-----------------------------------------------------------------------------------+")
			fmt.Println("\n          Terima kasih telah menggunakan Aplikasi Manajemen Fashion!            ")
			fmt.Println("                       Program otomatis akan ditutup. ")
			fmt.Println()

			return true
		}
	}

	return false
}

// hapus data sebelum
func clearScreen() {
	/*
		IS : Layar terminal menampilkan riwayat teks atau menu dari proses sebelumnya.
		FS : Layar terminal bersih dari teks lama.
	*/

	fmt.Print("\033[H\033[2J")
}

func inisialisasiData() {
	/*
		IS : Variabel global daftarToko masih kosong (berisi nilai bawaan/default struct) dan variabel jumlahData bernilai 0.
		FS : Array daftarToko telah terisi dengan 100 data pakaian tiruan (dummy data) siap pakai,
		dan variabel jumlahData berubah nilainya menjadi 100.
	*/

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
