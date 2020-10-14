package queue

var x []string

// List список элементов очереди.
func List() []string {
	return x
}

// Push добавит новый элемент в очередь.
func Push(str string) {
	x = append(x, str)
}

// Pop вернет первый добавленный в очередь элемент.
func Pop() string {
	numOfElements := len(x)
	// Когда очередт будет пустой, он вернет пустую строку.
	if numOfElements == 0 {
		return ""
	}

	popElem := x[0]
	x = x[1:]
	return popElem
}
