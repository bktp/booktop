<template>
    <div class="create-book">
        <sidebar></sidebar>
        <section class="form">
            <header>
                <h3 class="is-size-3">Редактирование книги</h3>
            </header>
            <form @submit.prevent="editBook">
                <div class="field">
                    <label class="label">Название</label>
                    <div class="control">
                        <input required v-model="book.name" class="input" type="text" placeholder="Название">
                    </div>
                </div>
                <div class="field">
                    <label class="label">isbn</label>
                    <div class="control">
                        <input required v-model="book.isbn" class="input" type="text" placeholder="1234567891234">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Оригинальное название</label>
                    <div class="control">
                        <input v-model="book.original" class="input" type="text" placeholder="Оригинальное название">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Автор(ы)</label>
                    <div class="control">
                        <input required v-model="book.authors" class="input" type="text" placeholder="John D., Dohn J.">
                    </div>
                </div>
                <div class="field">
                    <label class="label">Год</label>
                    <div class="control">
                        <input v-model="book.published" class="input" type="text" placeholder="2000 до н.э.">
                    </div>
                </div>

                <div class="field">
                    <label class="label">Категория</label>
                    <div class="control">
                        <div class="select">
                            <select v-model="book.category">
                                <option v-for="cat in cats" :key="cat.id">{{cat.name}}</option>
                            </select>
                        </div>
                    </div>
                </div>

                <div class="field">
                    <label class="label">Описание</label>
                    <div class="control">
                        <textarea v-model="book.description" class="textarea" placeholder="Описание"></textarea>
                    </div>
                </div>

                <div class="field">
                    <label class="label">Обложка</label>
                    <div class="control">
                        <input v-model="book.cover" class="input" type="text" placeholder="www.example.com/image.png">
                    </div>
                </div>

                <div class="field is-grouped is-grouped-multiline">
                    <div class="control">
                        <input type="submit" class="button" value="Сохранить изменения">
                    </div>
                    <div class="control">
                        <router-link :to="'/book/' + book.isbn + '/page/create'" class="button">Перейти к добавлению страниц</router-link>
                    </div>
                    <div class="control">
                        <router-link :to="'/book/' + book.isbn + '/page/edit'" class="button">Перейти к редактированию страниц</router-link>
                    </div>
                    <div class="control">
                        <a @click="deleteBook" class="button">Удалить книгу</a>
                    </div>
                </div>
            </form>
        </section>
    </div>
</template>

<style lang="scss" scoped>
.create-book {
  display: flex;
  .form {
    padding: 20px;
    flex: 1;
    header {
      margin-bottom: 20px;
    }
    form {
      max-width: 600px;
    }
  }
}
</style>


<script>
import Sidebar from "../../../components/Sidebar/Sidebar.vue";
export default {
  components: {
    Sidebar
  },
  data() {
    return {
      cats: [],
      book: {}
    };
  },
  methods: {
    editBook() {
      if (this.book.authors) this.book.authors = this.book.authors.split(",");
      this.$http
        .put(
          `http://localhost:8080/api/book/${this.book.isbn}`,
          JSON.stringify(this.book)
        )
        .then(resp => {
          console.log(resp);
        });
    },
    deleteBook() {
        this.$http.delete(`http://localhost:8080/api/book/${this.book.isbn}`).then(() => {
            this.$router.push("/main");
        })
    }
  },
  created() {
    this.$http
      .get("http://localhost:8080/api/categories")
      .then(resp => (this.cats = resp.body));
    this.$http
      .get(`http://localhost:8080/api/book/${this.$route.params.isbn}`)
      .then(resp => {
        this.book = resp.body;
        this.book.authors = this.book.authors.join(", ");
      });
  }
};
</script>
