package lang_types

type NPackage struct {
	Name          string
	ClassList     []*NClass
	InterfaceList []*NInterface
}

func NewPackage(name string) *NPackage {
	p := new(NPackage)
	p.Name = name
	p.ClassList = []*NClass{}
	return p
}

type NClass struct {
	Name string
}

func NewClass(name string, p *NPackage) *NClass {
	c := new(NClass)
	c.Name = name
	p.ClassList = append(p.ClassList, c)
	return c
}

type NInterface struct {
	Name string
}

func NewInterface(name string, p *NPackage) *NInterface {
	c := new(NInterface)
	c.Name = name
	p.InterfaceList = append(p.InterfaceList, c)
	return c
}
