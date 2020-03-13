<template>
<div class="main">
  <nav class="nav-bar">
    <img class="logo" src="../assets/trello-logo.png"/>   
    <v-btn class="login">login </v-btn>
    <v-btn class="signup float-right" 
           @click="$router.push('signup')"
    >signup </v-btn>
  </nav>
  <div class="body">
    <div class="free-login mx-auto">
      <v-row no-gutters class="d-flex">
        <v-col cols="5" class=" main-text">
            <p style="font-size: 50px; font-weight: bold; color: white; line-height: 1.2em;">Trello lets you work more collaboratively and get more done.</p>
            <p style="font-size: 20px; color: white;  line-height: 1.5em;">Trello’s boards, lists, and cards enable you to organize and prioritize your projects in a fun, flexible, and rewarding way.</p>
        </v-col>
        <v-col cols="7" class="float-right">
          <img src="../assets/trello-main.svg" style="width:80%"/>
        </v-col>
        
      </v-row>
      <v-form class="login-form" 
              ref="form"
              light
              >
        <v-text-field v-model="email"
                      class="mr-5 pa-0"
                      background-color='white'
                      placeholder="Email"
                      height="50"
        ></v-text-field>
        <v-btn color="success" 
                x-large 
                @click="onSubmit"
                >sign up - it's free</v-btn>
      </v-form>
    </div>
  </div>
</div>
  
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      email: '',
    }
  },
  methods: {
    onSubmit(){
      if (this.email) {
        this.$router.push({name: 'signup', params: {email: this.email}})
      } else {
        this.$router.push('signup')
      }
      
    }
  }
}
</script>
<style scoped>
.main {
  height: 100%;
  background: linear-gradient(135deg, #0079bf, #5067c5);

}
.nav-bar {
  height: 10%;
  width: 100%;
  display: flex;
  position: relative;
}
.logo {
  margin: 20px;
  width: 130px;
  height: 50%;
}
.login .signup {
  position: absolute;
  /* display: block; */
  float: right;
}
.body {
  width: 100%;
  height: 90%;
  position: relative;
}
.free-login {
  position: relative;
  top: 30px;
  width: 80%;
  height: 80%;
}
.login-form{
  display: flex;
  width: 60%;
  
}
</style>
