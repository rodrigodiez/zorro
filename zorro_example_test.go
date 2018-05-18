package zorro

func ExampleNew() {
	zorro := New(NewUUIDv4Generator(), NewInMemoryStorage())
	zorro.Mask("foo")
}
