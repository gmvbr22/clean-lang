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
	Expose         bool
	Name           string
	AnnotationList []*NAnnotation
}

func NewClass(name string, expose bool, p *NPackage) *NClass {
	c := new(NClass)
	c.Name = name
	c.Expose = expose
	p.ClassList = append(p.ClassList, c)
	return c
}

type NInterface struct {
	Expose bool
	Name   string
}

func NewInterface(name string, expose bool, p *NPackage) *NInterface {
	i := new(NInterface)
	i.Name = name
	i.Expose = expose
	p.InterfaceList = append(p.InterfaceList, i)
	return i
}

type NAnnotation struct {
	Name   string
	Values []*KeyValue
}

type KeyValue struct {
	Name  string
	Value interface{}
}

func NewAnnotation(name string) *NAnnotation {
	i := new(NAnnotation)
	i.Name = name
	i.Values = []*KeyValue{}
	return i
}

func NewKeyValue(name string, value interface{}) *KeyValue {
	i := new(KeyValue)
	i.Name = name
	i.Value = value
	return i
}
