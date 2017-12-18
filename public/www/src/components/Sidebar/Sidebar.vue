<template>
    <div class="sidebar">
        <aside class="menu">
            <header class="logo">
                <book-top></book-top>
            </header>
            <form @submit.prevent="search">
                <div class="field has-addons">
                    <div class="control">
                        <input v-model="pattern" class="input" type="text">
                    </div>
                    <div class="control">
                        <input type="submit" value="Поиск" class="button is-info">
                    </div>
                </div>
            </form>
            <template v-if="user.role == 'admin'">
                <p class="menu-label">
                    Меню администратора
                </p>
                <ul class="menu-list">
                    <li>
                        <router-link to="/categories">Управление категориями</router-link>
                    </li>
                    <li>
                        <router-link to="/book/create">Добавление книги</router-link>
                    </li>
                </ul>
            </template>
            <template v-if="user.role == 'admin' || user.role == 'user'">
                <p class="menu-label">
                    Меню пользователя
                </p>
                <ul class="menu-list">
                    <li>
                        <router-link to="/favs">Избранное</router-link>
                    </li>
                </ul>
            </template>
            <p class="menu-label">
                Категории
            </p>
            <ul class="menu-list">
                <li v-for="cat in categories" :key="cat.id">
                    <router-link :to="'/category/' + cat.name">{{cat.name}}</router-link>
                </li>
            </ul>
            <br>
            <a @click="logout" class="button">Выйти</a>
        </aside>
    </div>
</template>

<style lang="scss" scoped>
.sidebar {
  width: 320px;
  box-sizing: border-box;
  padding: 20px;
  .logo {
    margin-bottom: 20px;
  }
}
</style>


<script>
import BookTop from "../BookTop/BookTop.vue";
export default {
  components: {
    BookTop
  },
  data() {
    return {
      pattern: "",
      user: {},
      categories: []
    };
  },
  created() {
    this.$http
      .get("http://localhost:8080/api/user/role")
      .then(resp => (this.user = resp.body))
      .catch(err => console.log(err));
    this.getCategories();
  },
  methods: {
    getCategories() {
      console.log("worked");
      this.$http
        .get("http://localhost:8080/api/categories")
        .then(resp => (this.categories = resp.body))
        .catch(err => console.log(err));
    },
    logout() {
      this.$http
        .get("http://localhost:8080/api/logout")
        .then(() => this.$router.push("/"))
        .catch(err => console.log(err));
    },
    search() {
      this.$router.push(`/search/${this.pattern}`);
    }
  }
};
</script>

