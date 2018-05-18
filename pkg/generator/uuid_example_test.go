package generator

func ExampleNewUUIDv4() {
	generator := NewUUIDv4()
	generator.Generate("foo")
}
