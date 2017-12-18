<template>
  <div class="cat-search">
      <sidebar></sidebar>
      <section class="books">
          <header><h3 class="is-size-3">{{$route.params.category}}</h3></header>
          <div class="book-list">
              <book-thumb class="list-item" v-for="book in books" :key="book.isbn" :book="book"></book-thumb>
          </div>
      </section>
  </div>
</template>

<style lang="scss" scoped>
.cat-search {
    display: flex;
    .books {
        padding: 20px;
        header {
            margin-bottom: 20px;
        }
        .book-list {
            display: flex;
            .list-item {
                margin-right: 20px;
            }
        }

    }
}
</style>


<script>
import Sidebar from '../../../components/Sidebar/Sidebar.vue'
import BookThumb from '../../../components/BookThumb/BookThumb.vue'
export default {
  components: {
      Sidebar,
      BookThumb
  },
  data() {
      return {
          books: []
      }
  },
  created() {
      this.$http.get(`http://localhost:8080/api/category?category=${this.$route.params.category}`).then(resp => {
          this.books = resp.body;
      })
  }
}
</script>
