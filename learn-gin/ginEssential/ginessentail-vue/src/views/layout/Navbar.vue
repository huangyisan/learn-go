<template>
<div>
  <b-navbar toggleable="lg" type="dark" variant="info">
    <b-container>
    <b-navbar-brand href="#">OC</b-navbar-brand>

    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

    <b-collapse id="nav-collapse" is-nav>

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">

       <b-navbar-nav>
        <b-nav-item @click="$router.replace({name: 'login'})">登陆</b-nav-item>
        <b-nav-item @click="$router.replace({name: 'register'})">注册</b-nav-item>
      </b-navbar-nav>
      </b-navbar-nav>
    </b-collapse>
    </b-container>
  </b-navbar>
</div>
</template>