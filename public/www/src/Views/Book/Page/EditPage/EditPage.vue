<template>
    <div class="create-page">
        <sidebar></sidebar>
        <section class="create-form">
            <header>
                <h3 class="is-size-3">Добавление страницы</h3>
            </header>
            <div class="field">
                <label class="label">Выбор страницы: </label>
                <div class="select">
                    <select @change="choosePage" v-model="pagenum">
                        <option v-for="num in pages" :key="num">
                            {{num}}
                        </option>
                    </select>
                </div>
            </div>
            <form @submit.prevent="editPage" v-show="page.pagenum > 0">
                <div class="field">
                    <p class="is-size-6">Страница № {{page.pagenum}}</p>
                </div>
                <div class="field">
                    <label class="label">Текст страницы</label>
                    <div class="control">
                        <textarea v-model="page.text" class="textarea" placeholder="Описание"></textarea>
                    </div>
                </div>
                <div class="field">
                    <div class="control">
                        <input type="submit" class="button" value="Сохранить страницу">
                    </div>
                </div>
            </form>
        </section>
    </div>
</template>

<style lang="scss" scoped>
.create-page {
  display: flex;
  .create-form {
    flex: 1;
    padding: 20px;
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
import Sidebar from "../../../../components/Sidebar/Sidebar.vue";
export default {
  components: {
    Sidebar
  },
  data() {
    return {
      pages: 0,
      page: {
        pagenum: 0
      }
    };
  },
  methods: {
    editPage() {
      this.$http.put(
        `http://localhost:8080/api/book/${this.page.isbn}/pages/${this.page
          .pagenum}`,
        JSON.stringify(this.page)
      );
    },
    choosePage() {
      this.$http
        .get(
          `http://localhost:8080/api/book/${this.page.isbn}/pages/${this
            .pagenum}`
        )
        .then(resp => {
          this.page = resp.body;
        });
    }
  },
  created() {
    this.page.isbn = this.$route.params.isbn;
    this.$http
      .get(`http://localhost:8080/api/book/${this.page.isbn}/pages`)
      .then(resp => (this.pages = Number(resp.body)));
  }
};
</script>

