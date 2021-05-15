<template>
  <div>
    <b-navbar
      toggleable="lg"
      type="dark"
      variant="info"
    >
      <b-container>
        <b-navbar-brand
          href="#"
          @click="$router.push({name:'home'})"
        >Go-V</b-navbar-brand>

        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse
          id="nav-collapse"
          is-nav
        >

          <!-- Right aligned nav items -->
          <b-navbar-nav class="ml-auto">

            <!-- <b-nav-item-dropdown
              text="Lang"
              right
            >
              <b-dropdown-item href="#">EN</b-dropdown-item>
              <b-dropdown-item href="#">ES</b-dropdown-item>
              <b-dropdown-item href="#">RU</b-dropdown-item>
              <b-dropdown-item href="#">FA</b-dropdown-item>
            </b-nav-item-dropdown> -->

            <b-nav-item-dropdown
              right
              v-if="userInfo"
            >
              <template #button-content>
                <em>{{userInfo.name}}</em>
              </template>
              <b-dropdown-item
                href="#"
                @click="$router.replace('profile')"
              >个人主页</b-dropdown-item>
              <b-dropdown-item
                href="#"
                @click="Logout1"
              >登出</b-dropdown-item>
            </b-nav-item-dropdown>
            <div v-if="!userInfo">
              <b-nav-item
                v-if="$route.name != 'login'"
                @click="$router.replace('login')"
              >登录</b-nav-item>
              <b-nav-item
                v-if="$route.name != 'regi'"
                @click="$router.replace('regi')"
              >注册</b-nav-item>
            </div>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>
<script>
import storageService from '../../service/storageService'
import userService from '../../service/userService'
import { mapState, mapActions } from 'vuex'

export default {
  computed: mapState({
    userInfo: (state) => state.userModule.info
  }),
  methods: {
    ...mapActions('userModule', { userLogout: 'logOut' }),
    Logout1 () {
      // storageService.rem(storageService.USER_TOKEN)
      // storageService.rem(storageService.USER_INFO)
      this.userLogout({}).then(res => {
        this.$router.replace({ name: 'login' })
      })

    }
  }
}
</script>