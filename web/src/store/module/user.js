// import { reject, resolve } from "core-js/fn/promise";
import storageService from "../../service/storageService";
import UService from '../../service/userService'
const userModule = {
  namespaced:true,
  state:{
    token:storageService.get(storageService.USER_TOKEN),
    info:JSON.parse(storageService.get(storageService.USER_INFO)),
  },
  mutations:{
    LOGOUT(state){
      storageService.rem(storageService.USER_TOKEN)
      storageService.rem(storageService.USER_INFO)
      state.token = null
      state.info = null
    },
    SET_TOKEN(state,token){
      //更新本地缓存
      storageService.set(storageService.USER_TOKEN,token)
      state.token = token
    },
    SET_USERINFO(state,info){
      //更新本地缓存
      storageService.set(storageService.USER_INFO,JSON.stringify(info))
      state.info = info
    }
  },
  actions:{
    logOut ({commit},{}) {
      return new Promise((resolve,reject)=>{
        commit('LOGOUT')
        resolve()
      })
    },
    Regi (context,{name,phone,password}) {
      return new Promise((resolve,reject)=>{
        //发送请求
      UService.register({name,phone,password}).then(res => {
        if(res.data.code != 200){
          reject(res)
          return
        }
        //保存token
        context.commit('SET_TOKEN',res.data.data.token)
        //获取用户信息保存
        return UService.userInfo();
      }).then(res => {
        //跳转主页
        context.commit('SET_USERINFO',res.data.data.user)
        resolve(res)
      }).catch(err => {
        reject(err)
      })
      })
    },
    Login (context,{name,phone,password}) {
      return new Promise((resolve,reject)=>{
        //发送请求
      UService.login({name,phone,password}).then(res => {
        if(res.data.code != 200){
          reject(res)
          return
        }
        //保存token
        context.commit('SET_TOKEN',res.data.data.token)
        //获取用户信息保存
        return UService.userInfo();
      }).then(res => {
        //跳转主页
        context.commit('SET_USERINFO',res.data.data.user)
        resolve(res)
      }).catch(err => {
        reject(err)
      })
      })
    }
  }
}
export default userModule