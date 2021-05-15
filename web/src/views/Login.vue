<template>
  <div class="home">
    <b-row>
      <b-col
        md="8"
        offset-md="2"
        glg="6"
        offset-lg="3"
      >

        <b-card
          title="登录"
          class="mt-5"
        >
          <b-form>
            <b-form-group label="手机号:">
              <b-form-input
                v-model="$v.user.phone.$model"
                type="text"
                placeholder="输入手机号"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('phone')">
                手机号11
              </b-form-invalid-feedback>

            </b-form-group>
            <b-form-group label="密码:">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入密码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码长度在2-10
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                variant="outline-primary"
                block
                @click="Login"
              >
                登录
              </b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue';
import { required, minLength, maxLength } from 'vuelidate/lib/validators'
import validatePhone from '../helper/validator.js'
import storageService from '../service/storageService.js'

// import { mapMutations } from 'vuex'
import { mapActions } from 'vuex'
export default {
  name: 'login',
  components: {
    HelloWorld,
  },
  validations: {
    user: {
      phone: {
        required,
        // phone: validatePhone,
        minLength: minLength(11),
        maxLength: maxLength(11)
      },
      password: {
        required,
        minLength: minLength(2),
        maxLength: maxLength(20)
      },
    }
  },
  data () {
    return {
      user: {
        password: null,
        phone: null

      },
      // validPhone: null
    }
  },
  mounted () {
    console.log('process.env:', process.env)
  },
  methods: {
    // ...mapMutations('userModule', ['SET_TOKEN', 'SET_USERINFO']),
    ...mapActions('userModule', { userLogin: 'Login' }),
    validateState (name) {
      const { $dirty, $error } = this.$v.user[name]
      return $dirty ? !$error : null
    },
    Login () {
      // 验证数据
      this.$v.user.$touch()//设置dirty 为true
      if (this.$v.user.$anyError) {
        return
      }
      //发送请求
      this.userLogin(this.user).then(res => {
        this.$router.replace({ name: 'Home' })
      }).catch(err => {
        this.$bvToast.toast(err.data.msg, {
          title: 'Go-v',
          variant: 'danger',
          autoHideDelay: 5000
        })
        console.log('catch:', err.data.msg)
      })


    }
  }
};
</script>
