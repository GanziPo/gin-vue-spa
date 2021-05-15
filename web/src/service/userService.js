import request from '../util/request'
//用户注册
const register = ({name,phone,password})=>{
  return request.post('auth/register',{name,phone,password})
}
const login = ({phone,password})=>{
  return request.post('auth/login',{phone,password})
}
//用户信息
const userInfo = ()=>{
  return request.post('auth/info',{})
}
export default{
  register,
  userInfo,
  login
}