package arithmetic

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arithmetic Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (arithmetic Adapter) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (arithmetic Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (arithmetic Adapter) Division(a, b int32) (val int32, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	return a / b, nil
}
