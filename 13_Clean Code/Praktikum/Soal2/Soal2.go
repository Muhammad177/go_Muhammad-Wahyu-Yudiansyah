package main

type Kendaraan struct { //lebih baik pada struct jika tidak spasi setiap yang seharusnya spasi di beri huruf besar
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) berjalan() {
	m.tambahkecepatan(10) //jelaskan 10 itu apa
}

func (m *Mobil) tambahkecepatan(KecepatanBaru int) { //pada parameter var lebih baik menggunakan kata yang lebih jelas agar mudah di ingat untuk di panggil
	m.KecepatanPerJam = m.KecepatanPerJam + KecepatanBaru //tentukan satuan
}

func main() { //lebih baik pada struct jika tidak spasi setiap yang seharusnya spasi di beri huruf besar
	MobilCepat := Mobil{}
	MobilCepat.berjalan()
	MobilCepat.berjalan()
	MobilCepat.berjalan()

	MobilLamban := Mobil{}
	MobilLamban.berjalan()
}
