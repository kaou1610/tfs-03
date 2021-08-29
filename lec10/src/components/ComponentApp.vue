<template>
  <div class="component-app">
    <h1>{{ msg }}</h1>
    <p class="intro">{{intro}}</p>
    <span class="content">
      <div class="box-icon">
        <font-awesome-icon class="icon" icon="music" />
      </div>
      <div class="price">
        <span v-show="isShow" >
          <h4>Anual Plan</h4>
          <p>$55.99/year</p>
        </span>
        <span v-show="!isShow" >
          <h4>Pro Plan</h4>
          <p>$99.99/year</p>
        </span>
      </div>
      <div class="change" @click="isShow =! isShow">Change</div>
    </span>
    <button @click="showModal=true" class="btn">Proceed to Payments</button>
    <Modal v-if="showModal" @close="showModal=false"></Modal>
    <p class="cancle"><strong>Cancle Order</strong></p>
  </div>
</template>

<script>
import Modal from './Modal.vue'
export default {
  name: 'ComponentApp',
  props: {
    msg: String
  },
  data() {
    return {
      intro: 'You can now listen to millions of songs audiobooks and podcasts on any device anywhere you like!',
      isShow: true,
      showModal: false
    }
  },
  components: {
    Modal,
  },

}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.intro {
  width: 70%;
  margin: auto;
  opacity: 0.7;
  padding: 10px 60px;
}
.content {
  background-color: rgba(244, 239, 248, 0.918);
  width: 70%;
  margin: 20px auto;
  border-radius: 10px;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-evenly;
}
.box-icon {
  width: 55px;
  height: 55px;
  background-color: rgba(227, 216, 236, 0.918);
  border-radius: 50%;
  position: relative;
}
.icon {
  position: absolute;
  top: 18px;
  left: 18px;
}
.change {
  color: blue;
  text-decoration: underline;
  cursor: pointer;
}
.btn {
  width: 70%;
  height: 50px;
  margin: 20px auto;
  background-color: rgb(64, 64, 177);
  color: #fff;
  border-radius: 10px;
  border: none;
  box-shadow: 0px 10px 10px rgb(195, 195, 240);
  cursor: pointer;
}

.cancle {
  padding-bottom: 40px;
}
</style>
