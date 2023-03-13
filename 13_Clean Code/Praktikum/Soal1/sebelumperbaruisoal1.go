package main

type user struct {
	id       int
	username int//Rubah menjadi string
	password int//Rubah menjadi string
}

type userservice struct {//selanjutnya dalam penerapan code clean yaitu code yang mudah di pahami oleh orang lain pada masa datang tidak diterapkan karna code clean berarti kejelasan di struct baguan user service di gunakan var t untuk slice user tidak jelas apa maksud dari t
	t []user
}

func (u userservice) getallusers() []user {//disini juga penerapan clean code kurang jelas sebagai inisialisasi pada parameternya untuk memanggil struct userservice
	return u.t
}

func (u userservice) getuserbyid(id int) user {//disini juga penerapan clean code kurang jelas sebagai inisialisasi pada parameternya untuk memanggil struct userservice
	for _, r := range u.t {
		if id == r.id {//penjelasan looping pemanggilan menjadi tidak jelas
			return r
		}
	}

	return user{}
}//Banyak kekurang pada kode ini yang sangat jelas adalah tidak ditemukannya func main yang mengakibatkan program mengalami undeclarted atau tidak memberikan hasil