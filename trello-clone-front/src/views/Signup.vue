<template>
<div class="signup">
  <div class="title">
    trello
  </div>
  <div class="form-container">
    <v-card class="form-card mx-auto"
            max-width = "450"   > 
      <v-form class="pa-9">
        <h1>Sign up for your account</h1>
        <v-text-field outlined dense
                      placeholder="Enter email"
                      color="rgb(250, 251, 252)"
                      v-model="form.email"
                      :rules="rules.email"
                      required
        ></v-text-field>
        <v-text-field outlined dense
                      placeholder="Enter Full Name"
                      color="#FAFBFC"
                      v-model="form.fullName"
                      :rules="rules.fullName"
                      required
        ></v-text-field>
        <v-text-field outlined dense
        placeholder="Create Password"
                      color="#FAFBFC"
                      v-model="form.password"
                      :rules="rules.password"
                      required
        ></v-text-field>
        <v-btn @click="onSubmit" 
                :disabled="invalid"
                width="100%"
                height="30"
        >Sign Up
        </v-btn>  
        <v-divider class="my-5"></v-divider>
        <span class="bottom-link"
              @click="$router.push('login')"
        >Already have an account? Log In
        </span>
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
.bottom-link {
  text-align: center;
  font-size: 14px;
  color:hsl(216, 95%, 41%);
  display: block;
  cursor: pointer;

}
</style>