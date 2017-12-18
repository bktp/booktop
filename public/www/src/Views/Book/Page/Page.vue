<template>
    <section class="page-wrapper">
        <sidebar></sidebar>
        <div class="page">
            <header>
                <h3 class="is-size-3">Страница {{$route.params.page}}</h3>
            </header>
            <div class="page-text">
                {{page.text}}
            </div>
            <div class="level">
                <div class="level-right">
                    <div class="level-item">
                        <router-link :to="'/book/' + page.isbn + '/page/' + (page.pagenum - 1)">
                          <button class="button" :disabled="page.pagenum <= 1 || pages == 0" >Предыдущая страница</button>
                        </router-link>
                    </div>
                    <div class="level-item">
                        <router-link :to="'/book/' + page.isbn + '/page/' + (page.pagenum + 1)">
                          <button :disabled="page.pagenum >= pages || pages == 0" class="button">Следующая страница</button>
                    </router-link>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>


<style lang="scss" scoped>
.page-wrapper {
  display: flex;
  .page {
    flex: 1;
    padding: 20px;
    header {
      margin-bottom: 20px;
    }
    &-text {
      margin-bottom: 20px;
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
      pagenum: Number(this.$route.params.page),
      pages: 0,
      page: {}
    };
  },
  methods: {
    fetchPage(isbn, pagenum) {
      this.$http
        .get(`http://localhost:8080/api/book/${isbn}/pages/${pagenum}`)
        .then(resp => {
          this.page = resp.body;
          this.$http
            .get(`http://localhost:8080/api/book/${isbn}/pages`)
            .then(resp => (this.pages = Number(resp.body)));
        });
    }
  },
  beforeRouteUpdate(to, from, next) {
    this.page = {};
    this.fetchPage(to.params.isbn, to.params.page);
    next();
  },
  created() {
    this.fetchPage(this.$route.params.isbn, this.$route.params.page);
  }
};
</script>
