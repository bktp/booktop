<template>
    <div class="favs-page">
        <sidebar></sidebar>
        <div class="favs">
            <header><h3 class="is-size-3">Избранное</h3></header>
            <section class="books">
                <book-thumb class="books-item" v-for="book in books" :key="book.isbn" :book="book"></book-thumb>
            </section>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.favs-page {
    display: flex;
    .favs {
        padding: 20px;
        header {
            margin-bottom: 20px;
        }
        .books {
            display: flex;
            &-item {
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
    fetchBooks() {
      this.$http
        .get("http://localhost:8080/api/user/favs")
        .then(resp => (this.books = resp.body));
    }
  },
  created() {
      this.fetchBooks()
  }
};
</script>

