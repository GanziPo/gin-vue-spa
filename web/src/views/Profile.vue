<template>
  <h3>你好 <span v-if="userInfo">{{userInfo.name}}</span></h3>
</template>

<script>
// @ is an alias to /src
// import { mapMutations } from 'vuex'
// import { mapActions } from 'vuex'
import { mapState, mapActions } from 'vuex'
export default {
  name: 'Profile',
  components: {
    // HelloWorld,
  },
  computed: mapState({
    userInfo: (state) => state.userModule.info
  }),
  data () {
    return {
    }
  },
  mounted () {
    console.log('process.env:', process.env)
  },
  methods: {
  }
};
</script>
