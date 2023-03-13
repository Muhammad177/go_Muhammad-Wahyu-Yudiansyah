package main

import "fmt"

type user struct {
	id       int
	username string //Rubah menjadi string
	password string //Rubah menjadi string
}

type userservice struct { //selanjutnya dalam penerapan code clean yaitu code yang mudah di pahami oleh orang lain pada masa datang tidak diterapkan karna code clean berarti kejelasan di struct baguan user service di gunakan var t untuk slice user tidak jelas apa maksud dari t
	tabel []user
}

func (data userservice) getallusers() []user { //disini juga penerapan clean code kurang jelas sebagai inisialisasi pada parameternya untuk memanggil struct userservice
	return data.tabel
}

func (data userservice) getuserbyid(id int) user { //disini juga penerapan clean code kurang jelas sebagai inisialisasi pada parameternya untuk memanggil struct userservice
	for _, no := range data.tabel {
		if id == no.id {//penjelasan looping pemanggilan menjadi tidak jelas
			return no
		}
	}

	return user{}
}

func main() { //Banyak kekurang pada kode ini yang sangat jelas adalah tidak ditemukannya func main yang mengakibatkan program mengalami undeclarted atau tidak memberikan hasil

	data := userservice{
		tabel: []user{
			{id: 1, username: "user1", password: "pass1"},
			{id: 2, username: "user2", password: "pass2"},
			{id: 3, username: "user3", password: "pass3"},
		},
	}

	users := data.getallusers()

	fmt.Println("All users:")
	for _, all := range users {
		fmt.Printf("ID: %d, Username: %s, Password: %s\n", all.id, all.username, all.password)
	}

	id := 2
	user := data.getuserbyid(id)
	fmt.Printf("User dengan ID %d:\n", id)
	fmt.Printf("ID: %d, Username: %s, Password: %s\n", user.id, user.username, user.password)
}
