<template>
  <div>
    <h2>登录</h2>
    <form @submit.prevent="login">
      <label for="username">用户名:</label>
      <input type="text" v-model="username" id="username" />

      <label for="password">密码:</label>
      <input type="password" v-model="password" id="password" />

      <button type="submit">登录</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'UserLogin',
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    async login() {
      try {
        const loginData = {
          username: this.username,
          password: this.password,
          token: this.token
        };
        
        // 发送 POST 请求到后端进行登录验证
        const response = await axios.post('https://192.168.0.107:8888/user/login', loginData, {
          headers: {
            'Content-Type': 'application/json'
          }
        });
        
        // 根据返回结果处理
        if (response.status === 200) {
          // alert('登录成功！');
          console.log("resp = ",response.data)
          console.log("login token = ",response.data.token)
          this.token = response.data.token; // 保存 token
          console.log("this token = ",this.token)
          localStorage.setItem('token', this.token); // 将 token 存储到 localStorage
          this.$emit('loginSuccess', this.username,this.token);  // 触发父组件的登录成功事件
        } else {
          alert('登录失败，用户名或密码错误。');
        }
      } catch (error) {
        console.error('登录请求失败:', error);
        alert('登录请求失败，请检查服务器。');
      }
    }
  }
};
</script>
