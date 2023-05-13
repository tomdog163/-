package day01

import "fmt"

//type user struct {
//	name     string
//	password string
//}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func Example11() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)
	fmt.Println(checkPassword(a, "haha"))
	fmt.Println(checkPassword2(&a, "haha"))

	e := user{name: "wang", password: "1024"}
	e.resetPassword("2048")
	fmt.Println(e.checkPassword("2048"))
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}
