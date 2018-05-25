package list

import "errors"

const DefaultCapacity = 10

type List struct {
	size     int
	elements []interface{}
}

func New() *List {
	return &List{elements: make([]interface{}, DefaultCapacity)}
}

func (this *List) Add(element interface{}) {
	if this.size == len(this.elements) {
		this.increaseCapacity()
	}
	this.elements[this.size] = element
	this.size += 1
}

func (this *List) DeleteAt(index int) error {
	if index < 0 || index > this.size {
		return errors.New("Index out of bounds")
	}
	this.elements = append(this.elements[:index], this.elements[index+1:]...)
	this.size -= 1

	return nil
}

func (this *List) Get(index int) (interface{}, error) {
	if index < 0 || index > this.size {
		return nil, errors.New("Index out of bounds")
	}
	return this.elements[index], nil
}

func (this *List) increaseCapacity() {
	this.elements = append(this.elements, make([]interface{}, DefaultCapacity)...)
}

func (this *List) Size() int { return this.size }
