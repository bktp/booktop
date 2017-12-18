# booktop

> BookTop online library

## Build Setup

``` bash
#через go get
go get https://github.com/bktp/booktop

#или govendor
govendor get https://github.com/bktp/booktop
```

В директории проекта надо собрать vue-файлы. Сначала перейти в public/www

``` bash
npm install

# сборка
npm run webpack
```

Дамп для postgres в корне. После создания пользователя надо поменять в src/config/config.go данные бд.
