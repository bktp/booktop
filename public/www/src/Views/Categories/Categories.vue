<template>
    <div class="categories">
        <sidebar></sidebar>
        <div class="cats">
            <header>
                <h3 class="is-size-3">Категории</h3>
            </header>
            <section class="cats-list">
                <form @submit.prevent="addCat">
                    <div class="field has-addons">
                        <div class="control">
                            <input v-model="newcat" class="input" type="text" placeholder="Имя категории">
                        </div>
                        <div class="control">
                            <input class="button is-info" type="submit" value="Добавить">
                        </div>
                    </div>
                </form>
                <div class="cat">
                    <div v-for="cat in cats" :key="cat.id" class="field has-addons">
                        <div class="control">
                            <input v-model="cat.name" class="input" type="text" placeholder="Имя категории">
                        </div>
                        <div class="control">
                            <input class="button is-success" @click="renameCat(cat.id, cat.name)" type="button" value="Переименовать">
                        </div>
                        <div class="control">
                            <input class="button is-danger" @click="deleteCat(cat.id)" type="button" value="Удалить">
                        </div>
                    </div>
                </div>
            </section>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.categories {
  display: flex;
  .cats {
    flex: 1;
    padding: 20px;
    header {
      margin-bottom: 20px;
    }
    form {
      margin-bottom: 20px;
    }
    .cat {
      margin-bottom: 5px;
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
      newcat: "",
      cats: []
    };
  },
  methods: {
    updateCats() {
      this.$http
        .get("http://localhost:8080/api/categories")
        .then(resp => (this.cats = resp.body));
    },
    renameCat(id, name) {
      cat = {
        id,
        name
      };
      this.$http
        .put(`http://localhost:8080/api/categories/${id}`, JSON.stringify(cat))
        .then(() => this.updateCats());
    },
    deleteCat(id) {
      this.$http
        .delete(`http://localhost:8080/api/categories/${id}`)
        .then(() => this.updateCats());
    },
    addCat() {
      this.$http
        .post(`http://localhost:8080/api/categories?name=${this.newcat}`)
        .then(() => this.updateCats());
    }
  },
  created() {
    this.updateCats();
  }
};
</script>

