<template>
<div class="signup">
  <div class="title">
    trello

  </div>
  <div class="form-container">
    <v-card class="form-card mx-auto"
            max-width = "450"   > 
      <v-form class="pa-9">
        <p>Sign up for your account</p>
        <v-text-field placeholder="Enter email"
                      color="rgb(250, 251, 252)"
                      :value="form.email"
        ></v-text-field>
        <v-text-field placeholder="Enter Full Name"
                      color="#FAFBFC"
                      :value="form.fullName"
        ></v-text-field>
        <v-text-field placeholder="Create Password"
                      color="#FAFBFC"
                      :value="form.password"
        ></v-text-field>
        <v-btn >Sign Up</v-btn>  
      </v-form>
    </v-card>
  </div>
</div>
  
</template>
<script>
export default {
  name: 'Signup',
  data() {
    return {
      form: {
        email: '',
        fullName: '',
        password: ''
      },

    }
  },
  mounted() {
    this.form.email = this.$route.params.email
  }
  
}
</script>
<style scoped>
.signup{
  width: 100%;
  height: 100%;
}
.form-card {
  width: 50%;
  justify-content: center;
}
</style>