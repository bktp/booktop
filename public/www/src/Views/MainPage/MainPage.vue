<template>
  <div class="mainpage">
    <sidebar></sidebar>
    <main>
      <header>
        <h3 class="is-size-3">Популярные книги</h3>
      </header>
      <section class="books">
      <book-thumb class="books-item" :key="book.isbn" v-for="book in books" :book="book"></book-thumb>
      </section>
    </main>
  </div>
</template>

<style lang="scss" scoped>
.mainpage {
  display: flex;
  main {
    padding: 20px;
    header {
      margin-bottom: 20px;
    }
  }
}
.books {
  display: flex;
  flex-wrap: wrap;
  &-item {
    margin-right: 20px;
  }
}
</style>


<script>
import Sidebar from "../../components/Sidebar/Sidebar.vue";
import BookThumb from "../../components/BookThumb/BookThumb.vue";

export default {
  data() {
    return {
      books: []
    }
  },
  components: {
    Sidebar,
    BookThumb
  },
  methods: {
    fetchBooks() {
      this.$http.get("http://localhost:8080/api/book/search/popular").then(resp => {
        this.books = resp.body;
        console.dir(resp);
        });
    }
  },
  created() {
    this.fetchBooks()
  }
};
</script>

