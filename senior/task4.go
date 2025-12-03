package senior

import "net/http"

//Собственный минимальный router
//
//**Задача:**
//Реализовать простой роутер:
//
//- поддерживает пути вида `/users/{id}` и `/articles/{slug}`;
//
//- извлекает path-параметры и кладёт в `context` или `map[string]string`.

type HandlerFunc func(w http.ResponseWriter, r *http.Request, params map[string]string)

type Router struct {
	routes []route
}

type route struct {
	method  string
	pattern string
	handler HandlerFunc
}

// TODO: метод Handle(method, pattern, handler) и ServeHTTP
