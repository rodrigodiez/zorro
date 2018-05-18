package zorro

func ExampleNewUUIDv4Generator() {
	generator := NewUUIDv4Generator()
	generator.Generate("foo")
}
