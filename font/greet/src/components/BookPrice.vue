<template>
    <div v-if="isLoggedIn">
      <h2>查询图书价格</h2>
      <form @submit.prevent="searchPrice">
        <label for="bookName">图书名称:</label>
        <input type="text" v-model="bookName" id="bookName" />
  
        <button type="submit">查询</button>
      </form>
      <div v-if="price !== null">
        <p>图书价格: {{ price }} 元</p>
      </div>
    </div>
    <div v-else>
      <p>请先登录。</p>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    name: 'SearchBookPrice',
    props: ['isLoggedIn', 'token'],
    data() {
      return {
        bookName: '',
        price: null
      };
    },
    methods: {
      async searchPrice() {
        const token = localStorage.getItem('token'); // 从 localStorage 获取 token
    try {
        const response = await axios.get(
            `https://192.168.0.107:8888/check`, // 请求 URL
            {
                 params: { book: this.bookName }, // 将 bookName 作为查询参数
                headers: {
                    'Authorization': `Bearer ${token}`,  // 携带 token
                    'Content-Type': 'application/json'
                }
            }
        );
        if (response.status === 200) {
            this.price = response.data.price; // 处理返回的价格
        } else {
            this.price = '未知'; // 如果状态不是 200，设置为未知
        }
    } catch (error) {
      console.log("book = ",this.bookName)
        console.error('查询图书价格失败:', error); // 错误日志

        // alert('查询图书价格失败，请检查服务器。'); // 弹窗提示
    }
}

    }
  };
  </script>
  