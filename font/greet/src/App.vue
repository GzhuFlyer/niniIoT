<template>
  <div id="app">
    <Register />
    <Login @loginSuccess="handleLoginSuccess" />
    <BookPrice :isLoggedIn="isLoggedIn" />
  </div>
</template>

<script>
import Register from './components/UserRegister.vue';
import Login from './components/UserLogin.vue';
import BookPrice from './components/BookPrice.vue';

export default {
  components: {
    Register,
    Login,
    BookPrice
  },
  data() {
    return {
      isLoggedIn: false,
      token:""
    };
  },
  methods: {
    handleLoginSuccess(username,token) {
      this.isLoggedIn = true;
      this.token = token
      console.log(username, '登录成功');
    }
  }
};
</script>
