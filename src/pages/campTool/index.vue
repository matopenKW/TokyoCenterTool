<template>
    <div id="content">
        <div class="row">
            <div class="col-12">
                <div id="login-form" class="login">
                    <h2 class="login-header">ログイン</h2>
                    <form class="login-container">
                        <p><input type="email" placeholder="Email"></p>
                        <p><input type="password" placeholder="Password"></p>
                        <p><input type="submit" value="Log in" @click="submitLogin"></p>
                        <p><input type="submit" value="Gest User" @click="submitGetUser"></p>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>
<script>

export default {
    layout ({ store }) {
        return 'login'
    },
    data: function() {
        return {
            files: []
        }
    },
    mounted: function () {
        this.$nextTick(function () {
            console.log('load')
        })
    },
    methods: {
        async submitLogin(){
            alert('click Login')
        }


        // async getString() {
        //     this.res = await this.$axios.$get('/resString')
        // },
        // async getProfile() {
        //     let res = await this.$axios.$get('/getProfile',{
        //         params:{}
        //     })
        //     .then(response => {
        //         this.res = response.profile
        //     })
        //     .catch(err => {
        //         console.log(err)
        //     })
        //     .finally(() => console.log('finally'))
        // }
    }
}

</script>

<style>

.login {
  width: 400px;
  margin: 16px auto;
  font-size: 16px;
}

/* Reset top and bottom margins from certain elements */
.login-header,
.login p {
  margin-top: 0;
  margin-bottom: 0;
}

.login-header {
  background: #28d;
  padding: 20px;
  font-size: 1.4em;
  font-weight: normal;
  text-align: center;
  text-transform: uppercase;
  color: #fff;
}

.login-container {
  background: #ebebeb;
  padding: 12px;
}

/* Every row inside .login-container is defined with p tags */
.login p {
  padding: 12px;
}

.login input {
  box-sizing: border-box;
  display: block;
  width: 100%;
  border-width: 1px;
  border-style: solid;
  padding: 16px;
  outline: 0;
  font-family: inherit;
  font-size: 0.95em;
}

.login input[type="email"],
.login input[type="password"] {
  background: #fff;
  border-color: #bbb;
  color: #555;
}

/* Text fields' focus effect */
.login input[type="email"]:focus,
.login input[type="password"]:focus {
  border-color: #888;
}

.login input[type="submit"] {
  background: #28d;
  border-color: transparent;
  color: #fff;
  cursor: pointer;
}

.login input[type="submit"]:hover {
  background: #17c;
}

/* Buttons' focus effect */
.login input[type="submit"]:focus {
  border-color: #05a;
}
</style>