<template>
    <div class="book-page">
        <sidebar></sidebar>
        <div class="book">
            <header>
                <h3 class="is-size-3">{{book.name}}</h3>
            </header>
            <div class="info-wrapper">
                <div class="book-cover">
                    <img :src="book.cover" alt="">
                </div>
                <ul class="book-info">
                    <li>Автор(ы): {{book.authors.join(", ")}}</li>
                    <li v-if="book.original">Оригинальное название: {{book.original}}</li>
                    <li v-if="book.category">Категория: {{book.category}}</li>
                    <li>isbn: {{book.isbn}}</li>
                    <li v-if="book.published">Год: {{book.published}}</li>
                    <li v-if="book.description">Описание: {{book.description}}</li>
                    <li>
                        <div class="level">
                            <div class="level-left">
                                <div class="level-item">
                                    <router-link :to="'/book/' + book.isbn + '/page/1'" class="button">Читать</router-link>
                                </div>
                                <div class="level-item" v-if="inFavs">
                                    <a @click="deleteFromFavs" class="button">Удалить из избранного</a>
                                </div>
                                <div class="level-item" v-else>
                                    <a @click="toFavs" class="button">Добавить в избранное</a>
                                </div>
                                <div class="level-item">
                                    <router-link :to="'/book/' + book.isbn + '/edit'" class="button">Редактировать</router-link>
                                </div>
                            </div>

                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.level {
  padding-top: 20px;
}
.book-page {
  display: flex;
}
.book {
  padding: 20px;
  header {
    margin-bottom: 20px;
  }
  .info-wrapper {
    display: flex;
    .book-info {
      padding-left: 20px;
    }
    .book-cover {
      height: 400px;
      width: 300px;
      border-radius: 5px;
      img {
        height: 100%;
        width: auto;
      }
    }
  }
}
</style>


<script>
import Sidebar from "../../components/Sidebar/Sidebar.vue";
export default {
  components: {
    Sidebar
  },
  data() {
    return {
      inFavs: false,
      book: {}
    };
  },
  methods: {
    toFavs() {
      this.$http
        .post(`http://localhost:8080/api/user/favs/${this.book.isbn}`)
        .then(resp => (this.inFavs = true));
    },
    deleteFromFavs() {
      this.$http
        .delete(`http://localhost:8080/api/user/favs/${this.book.isbn}`)
        .then(resp => (this.inFavs = false));
    },
    loadData() {
      this.$http
        .get(`http://localhost:8080/api/book/${this.$route.params.isbn}`)
        .then(resp => {
          this.book = resp.body;
          this.$http
            .get(`http://localhost:8080/api/user/favs/${this.book.isbn}`)
            .then(resp => (this.inFavs = true))
            .catch(err => (this.inFavs = false));
        });
    }
  },
  props: ["isbn"],
  created() {
    this.loadData();
  }
};
</script>

