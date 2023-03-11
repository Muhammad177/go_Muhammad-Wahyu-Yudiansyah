package main

type kendaraan struct { //lebih baik pada struct jika tidak spasi setiap yang seharusnya spasi di beri huruf besar
	totalroda       int
	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahkecepatan(10) //jelaskan 10 itu apa
}

func (m *mobil) tambahkecepatan(kecepatanbaru int) { //pada parameter var lebih baik menggunakan kata yang lebih jelas agar mudah di ingat untuk di panggil
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru //tentukan satuan
}

func main() { //lebih baik pada struct jika tidak spasi setiap yang seharusnya spasi di beri huruf besar
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	mobillamban := mobil{}
	mobillamban.berjalan()
}
