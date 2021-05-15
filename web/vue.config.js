const resolve = dir => require('path').join(__dirname, dir)
module.exports = {
  lintOnSave: false,
  devServer: {
    // port :8088,// 更改默认启动端口
    proxy: {
      '/xapi': {     //这里最好有一个 /
          target: 'http://127.0.0.1:1016',  // 后台接口域名
          ws: true,        //如果要代理 websockets，配置这个参数
          secure: false,  // 如果是https接口，需要配置这个参数
          changeOrigin: true,  //是否跨域
          pathRewrite:{
              '^/xapi':''
          }
      }
  }
  },
  // chainWebpack: config => {
  //   config.resolve.alias
  //     .set("@", resolve("src"));
  // }
}