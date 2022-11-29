package arithmetic

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arithmetic Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arithmetic Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arithmetic Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arithmetic Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
