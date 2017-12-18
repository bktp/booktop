<template>
  <section class="login">
      <div class="login-form">
          <book-top class="login-form__header"></book-top>
          <form id="login-form" @submit.prevent="login">
              <header>
                  <h4 class="is-size-4">Вход</h4>
              </header>
              <div class="field">
                  <label class="label">Имя пользователя</label>
                  <div class="control">
                      <input v-model="user.name" class="input" type="text" placeholder="username" required>
                  </div>
              </div>
              <div class="field">
                  <label class="label">Пароль</label>
                  <div class="control">
                      <input v-model="user.password" class="input" type="password" placeholder="password" required>
                  </div>
              </div>
              <div class="field is-grouped is-grouped-multiline">
                  <div class="control">
                      <input type="submit" class="button" value="Войти">
                  </div>
                  <div class="control">
                      <input @click="register" type="button" class="button" value="Зарегистрироваться">
                </div>
                  <div class="control">
                      <router-link to="/main" class="button is-text">Войти в гостевом режиме</router-link>
                  </div>
              </div>
          </form>
      </div>
  </section>
</template>

<style lang="scss" scoped>
.login {
  height: 100vh;
  display: flex;
  justify-content: flex-end;
  background-image: url("http://localhost:8080/images/back-min.jpg");
  background-position: center;
  background-size: cover;
}
.login-form {
  height: 100%;
  background-color: white;
  padding: 20px;
  width: 320px;
  box-sizing: border-box;
  &__header {
    margin-bottom: 20px;
  }
}
</style>


<script>
import BookTop from "../../components/BookTop/BookTop";

export default {
  data() {
    return {
      user: {}
    };
  },
  methods: {
    login() {
      this.$http
        .post("http://localhost:8080/api/auth", JSON.stringify(this.user))
        .then(() => {
          this.$router.push("main");
        })
        .catch(err => {
          console.log(err);
          alert(`Ошибка авторизации: ${err.statusText}`);
        });
    },
    register() {
      if (!document.forms["login-form"].checkValidity()) {
        alert("Поля 'Имя пользователя' и 'Пароль' обязательны к вводу");
        return;
      };
      this.$http
        .post("http://localhost:8080/api/users", JSON.stringify(this.user))
        .then(() => {
          this.$http
            .post("http://localhost:8080/api/auth", JSON.stringify(this.user))
            .then(() => {
              this.$router.push("/main");
            });
        })
        .catch(err => alert(`Ошибка регистрации: ${err.statusText}`));
    }
  },
  components: {
    BookTop
  }
};
</script>
