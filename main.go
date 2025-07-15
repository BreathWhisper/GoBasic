package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type stringMap = map[string]string

func main() {
	bookmarks := stringMap{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Приложение для закладок")

Menu:
	for {
		chooseOption, err := getMenu()
		if err != nil {
			fmt.Println(err)
			continue
		}

		reader.ReadString('\n')

		switch chooseOption {
		case 1:
			showBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks, reader)
		case 3:
			deleteBookmark(bookmarks, reader)
		case 4:
			break Menu
		}
	}
}

func getMenu() (int, error) {
	var chooseOption int

	fmt.Println("Выберете пункт меню")
	fmt.Println("1. Показать закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выйти")

	fmt.Scan(&chooseOption)

	if chooseOption > 0 {
		return chooseOption, nil
	}
	return 0, errors.New("не верный ввод, попробуйте 1,2,3,4")
}

func addBookmark(bookmarks stringMap, reader *bufio.Reader) {

	fmt.Println("Введите название записки: ")
	caption, _ := reader.ReadString('\n')
	caption = strings.TrimRight(caption, "\r\n")

	fmt.Println("Введите текст записки: ")
	bookmarkText, _ := reader.ReadString('\n')
	bookmarkText = strings.TrimRight(bookmarkText, "\r\n")

	bookmarks[caption] = bookmarkText
}

func deleteBookmark(bookmarks stringMap, reader *bufio.Reader) {
	fmt.Println("Введите название записки которую хотите удалить: ")
	caption, _ := reader.ReadString('\n')
	caption = strings.TrimRight(caption, "\r\n")

	delete(bookmarks, caption)
}

func showBookmarks(bookmarks stringMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Нет закладок")
	}

	for key, value := range bookmarks {
		fmt.Println(key, value)
	}
}
