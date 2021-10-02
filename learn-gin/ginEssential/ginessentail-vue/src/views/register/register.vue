<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册">
          <b-form>
            <b-form-group>
              <b-form-input
                id="input-1"
                v-model="user.name"
                type="text"
                placeholder="输入用户名"
              ></b-form-input>
            </b-form-group>
            <b-form-group>
              <b-form-input
                id="input-1"
                v-model.trim="$v.user.telephone.$model"
                type="number"
                placeholder="输入手机号码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateStatus('telephone')">
                手机号码必须为11位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-form-input
                id="input-1"
                v-model.trim="$v.user.password.$model"
                type="password"
                placeholder="输入密码"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateStatus('password')">
                密码必须大于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button block variant="outline-primary" @click="register"
                >注册</b-button
              >
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength, maxLength } from "vuelidate/lib/validators";
export default {
  data() {
    return {
      user: {
        name: "",
        telephone: "",
        password: "",
      },
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        minLength: minLength(11),
        maxLength: maxLength(11),
      },
      password: {
        required,
        minLength: minLength(4),
      },
    },
  },
  methods: {
    validateStatus(name) {
      // 结构,只获取$dirty, $error的内容
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      console.log("register");
    },
  },
};
</script>

<style lang="scss" scoped>
</style>