<template>
  <div id="app">
    <div class="container">
      <img src="./assets/image/head.png" alt="image">
      <OrderSummary msg="Order Summary"/>
    </div>
  </div>
</template>

<script>
import OrderSummary from './components/OrderSummary.vue'

export default {
  name: 'App',
  components: {
    OrderSummary
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: rgb(219, 209, 245);
  width: 100%;
  height: auto;
  position: absolute;
  top: 0;
}
.container {
  width: 34%;
  margin: 70px auto;
  background-color: #fff;
  border-radius: 20px;
  
}
img {
  width: 100%;
  border-top-left-radius: 20px;
  border-top-right-radius: 20px;
}
</style>
