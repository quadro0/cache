# Cache

This program stores cache and could be used like in example:

```go
func main() {
	cache := cache.New()
	cache.Set("userId", 42)

	userId := cache.Get("userId")
	fmt.Println(userId) //output: 42

	cache.Delete("userId")
	userId = cache.Get("userId")
	fmt.Println(userId) //output: <nil>

	cache.Set("name", "Mykyta")
	name := cache.Get("name")
	fmt.Println(name) //output: Mykyta
}

```
