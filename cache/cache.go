package cache

var caches = make(map[string]interface{})

func AddCache(key string, val interface{}){
	caches[key] = val
}
func RemoveCache(key string){
	delete(caches, key)
}

func GetCache(key string) interface{}{
	return caches[key]
}

