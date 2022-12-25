package homework

type Mapa map[string]interface{}

func New() Mapa {
	return Mapa{}
}

func (m Mapa) Set(name string, value interface{}) {
	m[name] = value
}

func (m Mapa) Get(name string) interface{} {
	return m[name]
}

func (m Mapa) Delete(name string) {
	delete(m, name)
}
