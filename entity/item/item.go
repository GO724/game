package item

type Item struct {
	Name        string
	PlaceHolder string
	Quantity    int
}

func New(name, placeHolder string, quantity int) *Item {
	return &Item{
		Name:        name,
		PlaceHolder: placeHolder,
		Quantity:    quantity,
	}
}

func (i *Item) Increase() {
	i.Quantity++
}

func (i *Item) Decrease() bool {
	if i.Quantity <= 0 {
		return false
	}
	i.Quantity--
	return true
}

func (i *Item) Empty() bool {
	if i == nil {
		return true
	}
	return i.Quantity <= 0
}

func (i *Item) Have() bool {
	if i == nil {
		return false
	}
	return i.Quantity > 0
}
