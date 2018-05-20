package uuid

func ExampleNewV4() {
	generator := NewV4()
	generator.Generate("foo")
}
