package main

import (
	"net/http"

	"github.com/needsomesleeptd/annotarer-tech-ui/tech_ui/menus"
)

//use-case:
// Все:
// Проверка документа от нейросети (результат -- папка с файлами (проверенные))
// Загрузка документа

// Контроллер:
// Загрузка разметки
// Получение всех размеченных пользователем фото
// Получение всех типов, созданных пользователем
// Получение вообще всех типов
// Удаление разметки
//
// Админ
// Удаление типов разметки (каскадное)
// Смена роли пользователя
func main() {
	menu := menus.NewMenu()
	menu.RunMenu(http.DefaultClient)
}