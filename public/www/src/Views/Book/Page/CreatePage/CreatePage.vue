<template>
    <div class="create-page">
        <sidebar></sidebar>
        <section class="create-form">
            <header>
                <h3 class="is-size-3">Добавление страницы</h3>
            </header>
            <form @submit.prevent="addPage">
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
                        <input type="submit" class="button" value="Добавить страницу">
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
      page: {
          pagenum: 0
      }
    };
  },
  methods: {
    addPage() {
      this.$http
        .post(
          `http://localhost:8080/api/book/${this.page.isbn}/pages`,
          JSON.stringify(this.page)
        )
        .then(resp => {
          this.page.pagenum++;
          this.page.text = "";
        });
    }
  },
  created() {
    this.page.isbn = this.$route.params.isbn;
    this.$http
      .get(`http://localhost:8080/api/book/${this.page.isbn}/pages`)
      .then(resp => (this.page.pagenum = Number(resp.body) + 1));
  }
};
</script>

