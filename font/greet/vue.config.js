const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})
// vue.config.js
module.exports = {
//   devServer: {
//       proxy: {
//           '/api': {
//               target: 'http://192.168.0.107:8888',
//               changeOrigin: true,
//               pathRewrite: { '^/user': '' },
//           },
//       },
//   },
    // devServer: {
    //     proxy: 'http://192.168.0.107:8888',
    // }
};