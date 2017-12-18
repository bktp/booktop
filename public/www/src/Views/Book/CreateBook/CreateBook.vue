<template>
    <div class="create-book">
        <sidebar></sidebar>
        <section class="form">
            <header>
                <h3 class="is-size-3">Добавление книги</h3>
            </header>
            <form @submit.prevent="addBook">
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

                <div class="field">
                    <div class="control">
                        <input type="submit" class="button" value="Добавить книгу">
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
      addBook() {
          if (this.book.authors) this.book.authors = this.book.authors.split(",");
          this.$http.post("http://localhost:8080/api/book", JSON.stringify(this.book)).then(resp => {
              console.log(resp);
              this.book = {};
          });
      }
  },
  created() {
    this.$http
      .get("http://localhost:8080/api/categories")
      .then(resp => (this.cats = resp.body));
  }
};
</script>
