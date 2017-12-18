import Login from './Views/Login/Login.vue'
import MainPage from './Views/MainPage/MainPage.vue'
import Search from './Views/Search/Search.vue'

import Book from './Views/Book/Book.vue'
import CreateBook from './Views/Book/CreateBook/CreateBook.vue'
import EditBook from './Views/Book/EditBook/EditBook.vue'

import Page from './Views/Book/Page/Page.vue'
import CreatePage from './Views/Book/Page/CreatePage/CreatePage.vue'
import EditPage from './Views/Book/Page/EditPage/EditPage.vue'
import Favs from './Views/Favs/Favs.vue'

import Categories from './Views/Categories/Categories.vue'
import CatSearch from './Views/Categories/Search/CatSearch.vue'


const routes = [
    {
        path: "/",
        component: Login
    },
    {
        path: "/categories",
        component: Categories
    },
    {
        path: "/category/:category",
        component: CatSearch
                    },
    {
        path: "/favs",
        component: Favs
    },
    {
        path: "/main",
        component: MainPage
    },
    {
        path: "/search/:query",
        component: Search
    },
    {
        path: "/book/create",
        component: CreateBook
    },
    {
        path: "/book/:isbn",
        component: Book
    },
    {
        path: "/book/:isbn/edit",
        component: EditBook
    },
    {
        path: "/book/:isbn/page/create",
        component: CreatePage
    },
    {
        path: "/book/:isbn/page/edit",
        component: EditPage
    },
    {
        path: "/book/:isbn/page/:page",
        component: Page
    }
];

export default routes;