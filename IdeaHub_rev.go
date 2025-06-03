package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type ideStartUp struct {
	Judul, Kategori string
	vote            int
	Waktu           time.Time
}

var data []ideStartUp

func Tambah_IdeBaru(bacaData *bufio.Reader) {
	fmt.Print("Masukkan Judul Ide Baru : ")
	Judul, _ := bacaData.ReadString('\n')
	fmt.Print("Masukkan Kategori Judul Ide : ")
	Kategori, _ := bacaData.ReadString('\n')

	Judul = strings.TrimSpace(Judul)
	Kategori = strings.TrimSpace(Kategori)
	if Judul == "" || Kategori == "" {
		fmt.Println("Judul dan kategori yang dimasukkan tidak boleh kosong ")
		return
	}

	data = append(data, ideStartUp{Judul: Judul, Kategori: Kategori, Waktu: time.Now()})
	fmt.Println("Ide baru berhasil ditambahkan")
}

func Lihat_DataIde() {
	if len(data) == 0 {
		fmt.Println("Belum ada data Ide yang tersedia")
		return
	}
	for i, ide := range data {
		fmt.Printf("%d. %s [%s] (%d upvotes) - %s\n", i+1, ide.Judul, ide.Kategori, ide.vote, ide.Waktu.Format("2006-01-02"))
	}
}

func Vote_Ide(bacaData *bufio.Reader) {
	Lihat_DataIde()
	if len(data) == 0 {
		return
	}
	fmt.Print("Pilih ide yang ingin di vote : ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > len(data) {
		fmt.Println("Data yang anda masukkan tidak tersedia ")
		return
	}
	data[nomor-1].vote++
	fmt.Println("Vote berhasil diberikan")
}

func CariIde_Sequential(bacaData *bufio.Reader) {
	fmt.Print("Masukkan kata kunci : ")
	cariSequential, _ := bacaData.ReadString('\n')
	cariSequential = strings.TrimSpace(strings.ToLower(cariSequential))

	ketemu := false
	for _, ide := range data {
		if strings.Contains(strings.ToLower(ide.Judul), cariSequential) {
			fmt.Printf("[SEQUENTIAL] %s - %s (%d upvotes)\n", ide.Judul, ide.Kategori, ide.vote)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Ide dengan kata kunci yang anda cari tidak ditemukan")
	}
}

func CariIde_Binary(bacaData *bufio.Reader) {
	if len(data) == 0 {
		fmt.Println("Belum ada data Ide yang tersedia")
		return
	}

	sort.Slice(data, func(i_1, i_2 int) bool {
		return data[i_1].Judul < data[i_2].Judul
	})

	fmt.Print("Masukkan judul ide secara lengkap: ")
	binaryKey, _ := bacaData.ReadString('\n')
	binaryKey = strings.TrimSpace(binaryKey)

	batas_kiri, batas_kanan := 0, len(data)-1
	for batas_kiri <= batas_kanan {
		tengah := (batas_kiri + batas_kanan) / 2
		if data[tengah].Judul == binaryKey {
			fmt.Printf("[BINARY] %s - %s (%d upvotes)\n", data[tengah].Judul, data[tengah].Kategori, data[tengah].vote)
			return
		} else if data[tengah].Judul < binaryKey {
			batas_kiri = tengah + 1
		} else {
			batas_kanan = tengah - 1
		}
	}
	fmt.Println("Ide yang anda maksud tidak tersedia")
}

func UrutIde_Vote_Selection() {
	for i := 0; i < len(data)-1; i++ {
		i_max := i
		for j := i + 1; j < len(data); j++ {
			if data[j].vote > data[i_max].vote {
				i_max = j
			}
		}
		data[i], data[i_max] = data[i_max], data[i]
	}
	fmt.Println("Ide diurutkan dari vote tertinggi : ")
	Lihat_DataIde()
}

func UrutIde_Tanggal_Insertion() {
	for i := 1; i < len(data); i++ {
		InsertionKey := data[i]
		j := i - 1
		for j >= 0 && data[j].Waktu.After(InsertionKey.Waktu) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = InsertionKey
	}
	fmt.Println("Ide diurutkan berdasarkan tanggal dari terlama ke terbaru : ")
	Lihat_DataIde()
}

func DataIde_Populer(bacaData *bufio.Reader) {
	var hari int
	fmt.Print("Masukkan jumlah rentang hari terakhir: ")
	fmt.Scanln(&hari)

	batasHari := time.Now().AddDate(0, 0, -hari)
	ketemu := false
	for _, ide := range data {
		if ide.Waktu.After(batasHari) {
			fmt.Printf("[POPULER] %s - %s (%d upvotes)\n", ide.Judul, ide.Kategori, ide.vote)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Tidak terdapat ide pada rentang waktu tersebut")
	}
}

func Ubah_DataIde(bacaData *bufio.Reader) {
	Lihat_DataIde()
	if len(data) == 0 {
		return
	}
	fmt.Print("Masukkan nomor ide yang ingin diubah : ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > len(data) {
		fmt.Println("Nomor yang anda masukkan tidak tersedia pada daftar")
		return
	}

	fmt.Print("Masukkan judul ide baru : ")
	Judul, _ := bacaData.ReadString('\n')
	fmt.Print("Masukkan kategori ide baru : ")
	Kategori, _ := bacaData.ReadString('\n')

	data[nomor-1].Judul = strings.TrimSpace(Judul)
	data[nomor-1].Kategori = strings.TrimSpace(Kategori)
	fmt.Println("Data Ide berhasil diubah")
}

func Hapus_DataIde(bacaData *bufio.Reader) {
	Lihat_DataIde()
	if len(data) == 0 {
		return
	}
	fmt.Print("Masukkan nomor ide yang ingin dihapus: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > len(data) {
		fmt.Println("Nomor yang anda masukkan tidak tersedia pada daftar")
		return
	}
	data = append(data[:nomor-1], data[nomor:]...)
	fmt.Println("Ide berhasil dihapus")
}

func main() {
	bacaData := bufio.NewReader(os.Stdin)
	data = []ideStartUp{
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
