package main

import (
	"bufio"   //digunakan untuk membaca input dari user agar bisa membaca semua data tanpa terpengaruh spasi
	"fmt"     //untuk input output
	"os"      //digunakan bersamaan dengan bufio agar tau harus membaca apa
	"sort"    //digunakan untuk mengurutkan data sesuai abjad
	"strings" //digunakan untuk menghapus spasi dan mengubahnya menjadi huruf kecil
	"time"    //digunakan untuk menampilkan tanggal
)

type ideStartUp struct { //ini adalah template atau formulirkosong yang akan diisi
	Judul, Kategori string
	vote            int
	Waktu           time.Time
}

var data []ideStartUp // variabel data digunakan untuk menyimpan data yang telah diisi dari template ideStartUp

func Tambah_IdeBaru(bacaData *bufio.Reader) { //func untuk menambahkan isi data untuk ideStartup yang akan disimpan di variabel data, bacaData adalah alat untuk membaca input dari user yang telah kita buat di func main sebelumnya,
	fmt.Print("Masukkan Judul Ide Baru : ")
	Judul, _ := bacaData.ReadString('\n')       //bacaData adalah alat yang telah kita buat di func main untuk membaca inputan dari user dan Readstring digunakan karena inputan yang diminta adalah string, lalu isi dari data yang telah diinputkan akan disimpan di dalam variabel Judul
	fmt.Print("Masukkan Kategori Judul Ide : ") //judul adalah nilai yang akan diterima sedangkan _ adalah sebuah error yang akan langsung dibuang karena tidak akan digunakan
	Kategori, _ := bacaData.ReadString('\n')

	Judul = strings.TrimSpace(Judul) //digunakan untuk menghapus spasi diawal dan diakhir string Judul serta menghapus \n
	Kategori = strings.TrimSpace(Kategori)
	if Judul == "" || Kategori == "" { //program memeriksa apakah inputan untuk judul atau kategori kosong atau hanya berisi spasi
		fmt.Println("Judul dan kategori yang dimasukkan tidak boleh kosong ")
		return //program akan mengembalikan ke pilihan menu
	}

	data = append(data, ideStartUp{Judul: Judul, Kategori: Kategori, Waktu: time.Now()}) //memasukkan data yang diisi user kedalam template ideStartUp, waktu yang digunakan oleh time.Now berasal dari waktu internal laptop kita
	fmt.Println("Ide baru berhasil ditambahkan")                                         //append untuk menambahkan data baru kedalam variabel data yang isinya dari array ideStartUp
}

func Lihat_DataIde() { //func untuk menampilkan data Ide yang tersedia
	if len(data) == 0 { //program akan memeriksa apakah data kosong atau tidak, jika kosong maka akan dikembalikan ke menu pilihan jika tidak program akan lanjut menampilkan data yang tersedia
		fmt.Println("Belum ada data Ide yang tersedia")
		return
	}
	for i, ide := range data { //i: index data, ide: data yang akan ditampilkan, range data: isi dari variabel data yang akan ditampilkan
		fmt.Printf("%d. %s [%s] (%d upvotes) - %s\n", i+1, ide.Judul, ide.Kategori, ide.vote, ide.Waktu.Format("2006-01-02"))
	} //menampilkan daftar ide dengan format yang mudah untuk dibaca, kenapa i+1? karena index di go mulai dari 0 sedangkan urutan data yang ditampilkan dimulai dari 1
}

func Vote_Ide(bacaData *bufio.Reader) { //func untuk memberikan vote kepada ide yang telah ditampilkan
	Lihat_DataIde()     //menampilkan data ide yang tersedia
	if len(data) == 0 { //memeriksa apakah data kosong atau tidak
		return
	}
	fmt.Print("Pilih ide yang ingin di vote : ")
	var nomor int
	fmt.Scanln(&nomor)                  //memilih ide pada nomor berapa yang akan di vote
	if nomor < 1 || nomor > len(data) { //memeriksa nomor ide yang dipilih, nomor tidak boleh kurang dari 1 dan lebih dari jumlah data yang ada
		fmt.Println("Data yang anda masukkan tidak tersedia ") //jika nomor tidak tersedia maka akan di kembalikan ke menu pilihan diawal
		return
	}
	data[nomor-1].vote++ //program akan menambahkan vote (vote++) pada ide yang dipilih, yaitu nomor pilihan user -1 untuk memperoleh index dari data yang ingin di vote
	fmt.Println("Vote berhasil diberikan")
}

func CariIde_Sequential(bacaData *bufio.Reader) { //func untuk mencari ide dengan sequential
	fmt.Print("Masukkan kata kunci : ")                                 //meminta inputan dari user berupa kata kunci yang ingin dicari (tidak perlu kalimat lengkap)
	cariSequential, _ := bacaData.ReadString('\n')                      //membaca inputan dari user dan menyimpannya di variabel cariSequential
	cariSequential = strings.TrimSpace(strings.ToLower(cariSequential)) //menghapus spasi di awal dan akhir kalimat serta merubahnya menjadi huruf kecil semua

	ketemu := false            //deklarasikan ketemu menjadi false di awal
	for _, ide := range data { //mengambil data dari variabel data namun yang kita perlu hanya ide nya saja tidak dengan index nya, jadi kita abaikan index nya dengan menggunakan _
		if strings.Contains(strings.ToLower(ide.Judul), cariSequential) { //string.ToLower mengubah judul menjadi huruf kecil semua, strings.Contains digunakan untuk mencari kata kunci di dalam judul ide yang disimpan di variabel cariSequential
			fmt.Printf("[SEQUENTIAL] %s - %s (%d upvotes)\n", ide.Judul, ide.Kategori, ide.vote) //menampilkan judul ide yang telah ditemukan
			ketemu = true
		}
	}
	if !ketemu { //jika ide tidak terdapat pada daftar maka akan menampilkan pesan error
		fmt.Println("Ide dengan kata kunci yang anda cari tidak ditemukan")
	}
}

func CariIde_Binary(bacaData *bufio.Reader) { //func untuk mencari ide dengan binary
	if len(data) == 0 { //memeriksa apakah data kosong atau tidak
		fmt.Println("Belum ada data Ide yang tersedia")
		return
	}

	sort.Slice(data, func(i_1, i_2 int) bool { //mengurutkan slice data i_1 dan i_2 adalah index dari 2 elemen pada data yang akan dibandingkan
		return data[i_1].Judul < data[i_2].Judul //mengurutkan data secara ascending dari A-Z
	})

	fmt.Print("Masukkan judul ide secara lengkap: ")
	binaryKey, _ := bacaData.ReadString('\n') //membaca inputan dari user dan menyimpannya di variabel binary
	binaryKey = strings.TrimSpace(binaryKey)  //menghapus spasi di awal dan akhir kalimat serta \n

	batas_kiri, batas_kanan := 0, len(data)-1 //batas kiri adalah index 0, dan batas kanan kita set 0 diawal, serta len(data)-1 adalah jumlah index pada data
	for batas_kiri <= batas_kanan {
		tengah := (batas_kiri + batas_kanan) / 2 //mencari nilai tengah
		if data[tengah].Judul == binaryKey {     //konfirmasi apakah nilai tengah tersebut sama dengan kata kunci yang kita masukkan
			fmt.Printf("[BINARY] %s - %s (%d upvotes)\n", data[tengah].Judul, data[tengah].Kategori, data[tengah].vote)
			return
		} else if data[tengah].Judul < binaryKey { //hal yang dicari lebih besar daripada data tengah, maka akan ditambah satu untuk mencari di bagian kanan begitupun sebaliknya
			batas_kiri = tengah + 1
		} else {
			batas_kanan = tengah - 1
		}
	}
	fmt.Println("Ide yang anda maksud tidak tersedia")
}

func UrutIde_Vote_Selection() { //func untuk mengurutkan ide menggunakan selection sort
	for i := 0; i < len(data)-1; i++ {
		i_max := i                           //asumsikan i_max awal adalah i atau i dimasukkan ke variabel bernama i_max
		for j := i + 1; j < len(data); j++ { //looping untuk mencari ide dengan vote tertinggi
			if data[j].vote > data[i_max].vote { //jika ingin terurut dari kecil ke besar tinggal ubah menjadi <
				i_max = j //ganti isi dari i_max dengan j jika data[j].vote lebih besar dari data[i_max].vote
			}
		}
		data[i], data[i_max] = data[i_max], data[i] //tukar nilai
	}
	fmt.Println("Ide diurutkan dari vote tertinggi : ")
	Lihat_DataIde()
}

func UrutIde_Tanggal_Insertion() { //func untuk mengurutkan ide menggunakan insertion sort
	for i := 1; i < len(data); i++ {
		InsertionKey := data[i]
		j := i - 1
		for j >= 0 && data[j].Waktu.After(InsertionKey.Waktu) { //memeriksa apakah kartu yang telah diurutkan lebih baru dari data selanjutnya, jika ingin dari terbaru hanya ubah menjadi Before
			data[j+1] = data[j] //merubah posisi jika data yang telah diurutkan lebih baru dari data selanjutnya
			j--
		}
		data[j+1] = InsertionKey //memasukkan isi insertionKey kedalam index data yang kosong
	}
	fmt.Println("Ide diurutkan berdasarkan tanggal dari terlama ke terbaru : ")
	Lihat_DataIde()
}

func DataIde_Populer(bacaData *bufio.Reader) { //func untuk menampilkan ide populer bacaData
	var hari int
	fmt.Print("Masukkan jumlah rentang hari terakhir: ")
	fmt.Scanln(&hari) //user menginputkan rentang hari yang ingin dicari

	batasHari := time.Now().AddDate(0, 0, -hari) //mengurangi hari dari tanggal sekarang dan dimasukkan ke variabel batasHari
	ketemu := false
	for _, ide := range data { //memeriksa setiap ide tanpa menggunakan index
		if ide.Waktu.After(batasHari) { //memeriksa apakah ide tersebut lebih baru dari batas hari yang telah ditentukan
			fmt.Printf("[POPULER] %s - %s (%d upvotes)\n", ide.Judul, ide.Kategori, ide.vote)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak terdapat ide pada rentang waktu tersebut")
	}
}

func Ubah_DataIde(bacaData *bufio.Reader) { //func untuk mengubah data ide
	Lihat_DataIde() //menampilkan data ide
	if len(data) == 0 {
		return
	}
	fmt.Print("Masukkan nomor ide yang ingin diubah : ") //inputkan nomor ide yang ingin diubah
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > len(data) { //memeriksa apakah nomor yang diinputkan ada pada daftar atau tidak
		fmt.Println("Nomor yang anda masukkan tidak tersedia pada daftar")
		return
	}

	fmt.Print("Masukkan judul ide baru : ")
	Judul, _ := bacaData.ReadString('\n') //inputkan judul baru dan akan disimpan ke variabel judul
	fmt.Print("Masukkan kategori ide baru : ")
	Kategori, _ := bacaData.ReadString('\n')

	data[nomor-1].Judul = strings.TrimSpace(Judul)       //merubah isi dari judul lama dengan judul baru serta menghapus spasi dan \n
	data[nomor-1].Kategori = strings.TrimSpace(Kategori) //sama hanya untuk kategori
	fmt.Println("Data Ide berhasil diubah")
}

func Hapus_DataIde(bacaData *bufio.Reader) { //func untuk menghapus data ide
	Lihat_DataIde()     //melihat data yang tersedia
	if len(data) == 0 { //memeriksa apakah data kosong
		return
	}
	fmt.Print("Masukkan nomor ide yang ingin dihapus: ") //inputkan nomor ide yang ingin dihapus
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > len(data) { //memeriksa apakah nomor yang diinputkan ada pada daftar atau tidak
		fmt.Println("Nomor yang anda masukkan tidak tersedia pada daftar")
		return
	}
	data = append(data[:nomor-1], data[nomor:]...) //menghapus ide yang dipilih dan menggantikannya dengan isi data yang ada di index berikutnya
	fmt.Println("Ide berhasil dihapus")
}

func main() {
	bacaData := bufio.NewReader(os.Stdin) //variabel bacaData digunakan untuk menyimpan isi dari newreader yang dimana bufio.newreader berfungsi untuk membaca dan menampung input dari user dan os.stdin adalah sumber datanya
	data = []ideStartUp{                  //data dummy
		{"Ruang Guru", "Pendidikan", 4, time.Now().AddDate(0, 0, -6)},
		{"Tokopedia", "E-Commerce", 7, time.Now().AddDate(0, 0, -4)},
		{"Shopee", "E-Commerce", 8, time.Now().AddDate(0, 0, -9)},
		{"Home Assistant", "Teknologi", 10, time.Now().AddDate(0, 0, -5)},
		{"Peduli Umat", "Sosial", 5, time.Now().AddDate(0, 0, -2)},
	}

	for {
		fmt.Println("\n==============  IdeaHub Startup  ==============\n")
		fmt.Println("1. Tambahkan Ide ")
		fmt.Println("2. Lihat Daftar Ide ")
		fmt.Println("3. Vote Ide ")
		fmt.Println("4. Cari Ide (gunakan kata kunci) ")
		fmt.Println("5. Cari Ide (gunakan judul ide lengkap) ")
		fmt.Println("6. Tampilkan Ide Terurut (berdasarkan vote) ")
		fmt.Println("7. Tampilkan Ide Terurut (berdasarkan waktu) ")
		fmt.Println("8. Tampilkan Ide populer ")
		fmt.Println("9. Ubah Data Ide ")
		fmt.Println("10. Hapus Data Ide ")
		fmt.Println("11. Keluar Aplikasi ")
		fmt.Println("\n================== Pilih Menu ==================\n")

		fmt.Print("Pilih Menu (nomor) => ")

		var menuPilih int
		fmt.Scanln(&menuPilih)
		switch menuPilih {
		case 1:
			fmt.Println("\n================================================\n")
			Tambah_IdeBaru(bacaData)
		case 2:
			fmt.Println("\n================================================\n")
			Lihat_DataIde()
		case 3:
			fmt.Println("\n================================================\n")
			Vote_Ide(bacaData)
		case 4:
			fmt.Println("\n================================================\n")
			CariIde_Sequential(bacaData)
		case 5:
			fmt.Println("\n================================================\n")
			CariIde_Binary(bacaData)
		case 6:
			fmt.Println("\n================================================\n")
			UrutIde_Vote_Selection()
		case 7:
			fmt.Println("\n================================================\n")
			UrutIde_Tanggal_Insertion()
		case 8:
			fmt.Println("\n================================================\n")
			DataIde_Populer(bacaData)
		case 9:
			fmt.Println("\n================================================\n")
			Ubah_DataIde(bacaData)
		case 10:
			fmt.Println("\n================================================\n")
			Hapus_DataIde(bacaData)
		case 11:
			fmt.Println("Terimakasih Telah Menuangkan Ide Anda di IdeaHub Startup ")
			return
		default:
			fmt.Println("Pilihan yang Anda Masukkan Tidak Ada di Menu ")
		}
	}
}
