package any

type Routers struct {
	GetMap  map[string]interface{}
	PostMap map[string]interface{}
}

func (r *Routers) Get(path string, method interface{}) {
	if r.GetMap == nil {
		GetMap := map[string]interface{}{}
		GetMap[path] = method
		r.GetMap = GetMap
	}
	r.GetMap[path] = method
}

func (r *Routers) Post(path string, method interface{}) {
	if r.PostMap == nil {
		PostMap := map[string]interface{}{}
		PostMap[path] = method
		r.PostMap = PostMap
	}
	r.PostMap[path] = method
}
