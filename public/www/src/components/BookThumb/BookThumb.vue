<template>
    <router-link class="book-link" :to="'/book/' + book.isbn">
        <div class="book">
            <div class="img-wrapper">
                <img :src="book.cover" alt="">
            </div>
            <p class="is-size-7">{{book.authors.join(", ")}}</p>
            <p class="is-size-5">{{book.name}}</p>
        </div>
    </router-link>
</template>

<style lang="scss" scoped>
.book {
  width: 200px;
  .img-wrapper {
    height: 300px;
    img {
      height: 100%;
      width: auto;
    }
  }
}
.book-link {
  text-decoration: none;
  color: inherit;
}
</style>


<script>
export default {
  props: ["book"]
};
</script>

