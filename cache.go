package cache

type Cache struct {
	store map[string]interface{}
}

func New() Cache {
	return Cache{
		store: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.store[key] = value
}

func (c Cache) Get(key string) interface{} {
	return c.store[key]
}

func (c Cache) Delete(key string) {
	delete(c.store, key)
}
