<template>
    <div class="search-page">
        <sidebar></sidebar>
        <section class="search-results">
            <header>
                <h3 class="is-size-3">Результаты поиска</h3>
            </header>
            <div class="books">
                <book-thumb class="book-item" :key="book.isbn" v-for="book in books" :book="book"></book-thumb>
            </div>
        </section>
    </div>
</template>

<style lang="scss" scoped>
.search-page {
  display: flex;
  .search-results {
    padding: 20px;
    header {
      margin-bottom: 20px;
    }
    .books {
      display: flex;
      .book-item {
        margin-right: 20px;
      }
    }
  }
}
</style>


<script>
import Sidebar from "../../components/Sidebar/Sidebar.vue";
import BookThumb from "../../components/BookThumb/BookThumb.vue";

export default {
  components: {
    Sidebar,
    BookThumb
  },
  data() {
    return {
      books: []
    };
  },
  methods: {
    fetchBooks(query) {
      this.$http
        .get(`http://localhost:8080/api/book/search?pattern='${query}'`)
        .then(resp => (this.books = resp.body));
    }
  },
  created() {
    this.fetchBooks(this.$route.params.query);
  },
  beforeRouteUpdate(to, from, next) {
    this.books = [];
    this.fetchBooks(to.params.query);
    next();
  }
};
</script>
