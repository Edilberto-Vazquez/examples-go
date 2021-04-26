package mypackage

type Pc struct {
	ram   int
	disk  int
	brand string
}

func New(rm, dsk int, brn string) Pc {
	myPC := Pc{ram: rm, disk: dsk, brand: brn}
	return myPC
}

func (myPC *Pc) DuuplicateRAM() {
	myPC.ram = myPC.ram * 2
}
