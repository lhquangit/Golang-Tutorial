
// - This package define the fixed data
// - In the real projects, we define and connect to database here


package data

import "learn-golang/pkg/dto"

var Todos []dto.Todo

func init() {
	Todos = []dto.Todo{
		{ID: 1, Name: "Monday", Content: "Learn Maths"},
		{ID: 2, Name: "Tuesday", Content: "Learn Literature"},
		{ID: 3, Name: "Wednesday", Content: "Learn Physics"},
		{ID: 4, Name: "Thursday", Content: "Learn Chemistry"},
		{ID: 5, Name: "Friday", Content: "Learn English"},
	}
}

