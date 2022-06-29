package receiverfunction

/**
 * @author Dev Problems(A Sarang Kumar Tak)
 * @YoutubeChannel <a href="https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ">...</a>
 */

type Person struct {
	Name string
	Age  int
}

//SetName value-based receiver function
func (p Person) SetName(name string) {
	println("address of person object in RF", &p)
	p.Name = name
}

//SetAge pointer-based receiver function
func (p *Person) SetAge(age int) {
	println("address of person object in RF", p)
	p.Age = age
}
