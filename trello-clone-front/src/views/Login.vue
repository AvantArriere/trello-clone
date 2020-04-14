<template>
<div class="login">
  <div class="title">
    trello
  </div>
  <div class="form-container">
    <v-card class="form-card mx-auto"
            max-width = "450"   > 
      <v-form class="pa-9">
        <h1>Log in to Trello</h1>
        <v-text-field outlined dense
                      placeholder="Enter email"
                      color="rgb(250, 251, 252)"
                      v-model="form.email"
                      :rules="rules.email"
                      required
        ></v-text-field>
        <v-text-field outlined dense
        placeholder="Enter Password"
                      color="#FAFBFC"
                      v-model="form.password"
                      :rules="rules.password"
                      required
        ></v-text-field>
        <v-btn @click="onSubmit" 
                color=success
                width="100%"
                height="30"
        >Log in
        </v-btn>  
        <v-divider class="my-5"></v-divider>
        <ul class="bottom-links">
          <li style="list-style: none;">
            <a class="bottom-link" @click="$router.push('recovery')">Can't log in?</a>
          </li>
          <li>
            <a class="bottom-link" @click="$router.push('signup')">Sign up for an account</a>
          </li>
        </ul>
      </v-form>
    </v-card>
  </div>
</div>
  
</template>
<script>
import rules from '../rules.js'
export default {
  name: 'Signup',
  data: function() {
    return {
      rules: rules(),
      form: {
        email: '',
        fullName: '',
        password: ''
      },
      invalid: true,

    }
  },
  mounted() {
    this.form.email = this.$route.params.email
  },
  methods: {
    onSubmit() {
      console.log('input:' ,this.form)
    }
  }
  
}
</script>
<style scoped>
.signup{
  width: 100%;
  height: 100%;
  background: #F9FAFC;
}
.title {
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 40px;


}
.form-card {
  width: 30%;
  justify-content: center;
}

h1 {
  text-align: center;
  color: #5E6C84;
  font-size: 16px;
  margin-top: 10px;
  margin-bottom: 25px;

}
.bottom-links {
  display: inline-block;
  padding: 0;
  width: 100%;
}

.bottom-link {
  font-size: 14px;
  color:hsl(216, 95%, 41%);
  cursor: pointer;
}
</style>