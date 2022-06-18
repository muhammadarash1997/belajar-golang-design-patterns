/// Product B ///

// struct/class short akan diwarisi oleh child class, tetapi karena golang bukan OOP
// maka inheritance bisa diakali dengan cara child class memiliki field berupa object
// dari parent class agar child class bisa memiliki property-property milik parent class
// sebenarnya ini belum cukup, kita harus menggunakan bantuan interface yang mana parent class
// harus mengimplementasi interface agar child class bisa dikenali dengan tipe parent
// classnya. Di sini iShort adalah interface dari parent class (short), jika nanti ada
// yang mau mewarisi/inhereting short class maka dia juga akan bertipe interface iShort.

// Intinya kita harus menggunakan model Inheritance 2 yang mana menggunakan bantuan interface.
// Jika ingin tau apa itu Inheritance 1 atau Inheritance 2 maka lihat di folder belajar-golang.

package main

type iShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) setSize(size int) {
	s.size = size
}

func (s *short) getLogo() string {
	return s.logo
}

func (s *short) getSize() int {
	return s.size
}