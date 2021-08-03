package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Print("Hello World!")
	// Инициализируем gorilla/mux роутер
	r := mux.NewRouter()

	// Страница по умолчанию для нашего сайта это простой html.
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// Наше API состоит из трех роутов
	// /status - нужен для проверки работоспособности нашего API
	// /products - возвращаем набор продуктов,
	// по которым мы можем оставить отзыв
	// /products/{slug}/feedback - отображает фидбек пользователя по продукту
	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	// Статику (картинки, скрипти, стили) будем раздавать
	// по определенному роуту /static/{file}
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))

	// Наше приложение запускается на 3000 порту.
	// Для запуска мы указываем порт и наш роутер
	http.ListenAndServe(":3000", r)

}

// Необходимо реализовать хендлер NotImplemented.
// Этот хендлер просто возвращает сообщение "Not Implemented"
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
