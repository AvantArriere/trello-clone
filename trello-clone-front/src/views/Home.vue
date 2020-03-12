<template>
<div class="main">
  <nav class="nav-bar">
    <img class="logo" src="../assets/trello-logo.png"/>   
    <v-btn class="login">login </v-btn>
    <v-btn class="signup float-right">signup </v-btn>
  </nav>
  <div class="body">
  </div>
</div>
  
</template>

<script>
export default {
  name: 'Home',
  components: {
  }
}
</script>
<style scoped>
.main {
  height: 100%;
  background: linear-gradient(135deg, #0079bf, #5067c5);

}
.nav-bar {
  height: 100px;
  width: 100%;
  display: flex;
  position: relative;
}
.logo {
  margin: 20px;
  width: 200px;
  height: 50%;
}
.login .signup {
  position: absolute;
  /* display: block; */
  float: right;
}
.body {
  width: 100%;
  /* height: 100%; */
}
</style>
