<template>
  <div>
    <h2>注册</h2>
    <form @submit.prevent="register">
      <label for="username">用户名:</label>
      <input type="text" v-model="username" id="username" />

      <label for="password">密码:</label>
      <input type="password" v-model="password" id="password" />

      <label for="age">年龄:</label>
      <input type="number" v-model="age" id="age" />

      <label for="gender">性别:</label>
      <select v-model="gender" id="gender">
        <option value="男">男</option>
        <option value="女">女</option>
      </select>

      <button type="submit">注册</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'UserRegister',
  data() {
    return {
      username: '',
      password: '',
      age: '',
      gender: '男'
    };
  },
  methods: {
    async register() {
      try {
        const userData = {
          username: this.username,
          password: this.password,
          age: this.age,
          gender: this.gender
        };
        
        const response = await axios.post('https://192.168.0.107:8888/user/register', userData, {
          headers: {
            'Content-Type': 'application/json'
          }
        });
        
        // 根据返回结果处理
        if (response.status === 200) {
          alert('注册成功！');
        } else {
          // alert('注册失败，请重试。');
        }
      } catch (error) {
        console.error('注册请求失败:', error);
        // alert('注册请求失败，请检查服务器。');
      }
    }
  }
};
</script>
