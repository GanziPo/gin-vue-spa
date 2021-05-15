// 本地缓存服务

const PREFIX = "GOV_"
const USER_PREFIX = `${PREFIX}user_`
const USER_TOKEN = `${USER_PREFIX}token`
const USER_INFO = `${USER_PREFIX}info`

//储存
const set = (key,data)=>{
  localStorage.setItem(key,data)
}

const get = (key,data)=>{
  return localStorage.getItem(key)
}
const rem = (key,data)=>{
  return localStorage.removeItem(key)
}

export default {
  set,
  get,
  rem,
  USER_TOKEN,
  USER_INFO

}