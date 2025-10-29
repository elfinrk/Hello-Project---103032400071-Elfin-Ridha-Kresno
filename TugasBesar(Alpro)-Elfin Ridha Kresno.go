// Edited by Elfin Ridha Kresno - 103032400071

package main

import "fmt"

type dataPolusi struct {
	lokasi string
	waktu string
	aqi int
	sumber string
	kategori string
}

var data []dataPolusi

func main() {
	for {
		fmt.Println("Menu: ")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Lihat Semua Data")
		fmt.Println("3. Edit Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Cari Data Polusi Berdasarkan Nama kota")
		fmt.Println("6. Urutkan berdasarkan AQI")
		fmt.Println("7. Urutkan bedasarkan Waktu")
		fmt.Println("8. AQI Tertinggi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		var p int
		fmt.Scan(&p)

		switch p {
		case 1: tambah()
		case 2: lihat()
		case 3: edit()
		case 4: hapus()
		case 5: cariKota()
		case 6: sortAQI()
		case 7: sortWaktu()
		case 8: cariTertinggi()
		case 0: return
		}
	}
}

func tambah() {
	var tambah dataPolusi

	fmt.Print("Lokasi: ")
	fmt.Scan(&tambah.lokasi)
	fmt.Print("Waktu (DD-MM-YYYY): ")
	fmt.Scan(&tambah.waktu)
	fmt.Print("Besar AQI: ")
	fmt.Scan(&tambah.aqi)
	fmt.Print("Sumber: ")
	fmt.Scan(&tambah.sumber)

	tambah.kategori = kategori(tambah.aqi)

	data = append(data, tambah)
}

func lihat() {
	if len(data) == 0 {
		fmt.Println("Data kosong, Harap isi terlebih dahulu.")
		return
	}

	fmt.Printf("%-3s | %-10s | %-20s | %-5s | %-12s | %-10s\n", "No", "Lokasi", "Waktu", "AQI", "Sumber", "Kategori")
	cetak(0)
}

func cetak(i int) {
	if i == len(data) {
		return
	}
	fmt.Printf("%-3d | %-10s | %-20s | %-5d | %-12s | %-10s\n", i+1, data[i].lokasi, data[i].waktu, data[i].aqi, data[i].sumber, data[i].kategori)
	cetak(i + 1)
}


func kategori(aqi int)string {
	switch {
		case aqi <= 50: 
		return "Baik"
		case aqi <= 100: 
		return "Sedang"
		case aqi <= 150:
		return "Tidak Sehat"
		default: 
		return "Berbahaya"
	}
}

func edit() {
	var lokasi, waktu string

	if len(data) == 0 {
		fmt.Println("Belum ada data")
		return
	}

	fmt.Print("Lokasi yang ingin di edit: ")
	fmt.Scan(&lokasi)
	fmt.Print("Waktu (DD-MM-YYY): ")
	fmt.Scan(&waktu)

	for i := 0; i < len(data); i++ {
		if data[i].lokasi == lokasi && data[i].waktu == waktu {
		   fmt.Print("AQI Baru: ")
		   fmt.Scan(&data[i].aqi)
		   fmt.Print("Sumber: ")
		   fmt.Scan(&data[i].sumber)

		   data[i].kategori = kategori(data[i].aqi)

		   fmt.Println("Data Berhasil di Edit")
		   return
		}
	}
	fmt.Println("Data tidak ditemukan")
}

func hapus() {
	var lokasi, waktu string

	if len(data) == 0 {
		fmt.Println("Belum ada data")
	}

	fmt.Print("Lokasi yang ingin dihapus: ")
	fmt.Scan(&lokasi)
	fmt.Print("Waktu data (DD-MM-YYYY): ")
	fmt.Scan(&waktu)

	for i := 0; i < len(data); i++ {
		if data[i].lokasi == lokasi && data[i].waktu == waktu {
			for j := i; j < len(data)-1; j++ {
				data[j] = data[j+1]
			}
			data = data[:len(data)-1]

			fmt.Println("Data berhasil dihapus")
			return
		}
	}
	fmt.Println("Data tidak ditemukan")
}

func cariKota() { 
	if len(data) == 0 {
		fmt.Println("Data kosong, silakan tambahkan data terlebih dahulu.")
		return
	}

	var kotaCari string

	fmt.Print("Masukkan nama kota yang ingin dicari: ")
	fmt.Scan(&kotaCari)

	var hasil []dataPolusi

	for _, d := range data {
		if d.lokasi == kotaCari {
			hasil = append(hasil, d)
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Data tidak ditemukan untuk kota", kotaCari)
		return
	}

	N := len(hasil)
	pass := 1

	for pass < N {
		idx := pass - 1
		i := pass

		for i < N {
			if hasil[i].waktu < hasil[idx].waktu {
				idx = i
			}
			i++
		}

		hasil[pass-1], hasil[idx] = hasil[idx], hasil[pass-1]

		pass++
	}

	fmt.Println("Data polusi di kota", kotaCari, "setelah diurutkan berdasarkan waktu:")
	fmt.Printf("%-3s | %-10s | %-20s | %-5s | %-12s | %-10s\n", "No", "Lokasi", "Waktu", "AQI", "Sumber", "Kategori")

	for i, d := range hasil {
		fmt.Printf("%-3d | %-10s | %-20s | %-5d | %-12s | %-10s\n", i+1, d.lokasi, d.waktu, d.aqi, d.sumber, d.kategori)
	}
}

func sortAQI() { 
	if len(data) == 0 {
		fmt.Println("Data kosong, silakan tambahkan data terlebih dahulu.")
		return
	}

	var i, idx, pass int
	var temp dataPolusi
	N := len(data)

	pass = 1

	for pass < N {
		idx = pass - 1
		i = pass

		for i < N {
			if data[i].aqi < data[idx].aqi { 
				idx = i
			}
			i = i + 1
		}

		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp

		pass = pass + 1
	}

	fmt.Println("Data berhasil diurutkan berdasarkan AQI:")
	lihat()

}

func sortWaktu() {
	if len(data) == 0 {
		fmt.Println("Data kosong, silakan tambahkan data terlebih dahulu.")
		return
	}

	var i, pass int
	var temp dataPolusi
	var d1, m1, y1, d2, m2, y2 int
	N := len(data)

	pass = 1
	for pass <= N-1 {
		i = pass
		temp = data[pass]

		fmt.Sscanf(temp.waktu, "%d-%d-%d", &d1, &m1, &y1)

		for i > 0 {
			fmt.Sscanf(data[i-1].waktu, "%d-%d-%d", &d2, &m2, &y2)

			if y1 < y2 || (y1 == y2 && m1 < m2) || (y1 == y2 && m1 == m2 && d1 < d2) {
				data[i] = data[i-1]
				i = i - 1
			} else {
				break
			}
		}
		data[i] = temp
		pass = pass + 1
	}

	fmt.Println("Data berhasil diurutkan berdasarkan waktu:")
	lihat()
}

func cariTertinggi() { 
	if len(data) == 0 {
		fmt.Println("Data kosong, silakan tambahkan data terlebih dahulu.")
		return
	}

	var aqiList []int
	for _, d := range data {
		aqiList = append(aqiList, d.aqi)
	}

	n := len(aqiList)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if aqiList[j] < aqiList[minIdx] {
				minIdx = j
			}
		}
		aqiList[i], aqiList[minIdx] = aqiList[minIdx], aqiList[i]
	}

	maxAQI := aqiList[n-1]

	left, right := 0, n-1
	found := false

	for left <= right {
		mid := (left + right) / 2
		if aqiList[mid] == maxAQI {
			found = true
			break
		} else if maxAQI < aqiList[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if found {
		fmt.Println("Data dengan AQI tertinggi:")
		for _, d := range data {
			if d.aqi == maxAQI {
				fmt.Printf("Lokasi   : %s\n", d.lokasi)
				fmt.Printf("Waktu    : %s\n", d.waktu)
				fmt.Printf("AQI      : %d\n", d.aqi)
				fmt.Printf("Sumber   : %s\n", d.sumber)
				fmt.Printf("Kategori : %s\n", d.kategori)
				fmt.Println()
			}
		}
	} else {
		fmt.Println("AQI tertinggi tidak ditemukan.")
	}
}