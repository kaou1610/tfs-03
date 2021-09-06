<template>
  <div class="chat">
        <div class="screen">
        <div id="header">
            <div class="nav avt"><p>A</p></div>
            <div class=" nav name">Anne</div>
        </div>
        <div id="container">
            <ul class="chatbox">
                <li v-for="item in msg" :key="item">{{item}}</li>
            </ul>
        </div>
        <div class="send">        
            <input type="text" class="input-mess" @keyup.enter="send()" v-model="title" placeholder="Type your message">
            <button type="submit" class="btn-mess" @click="send()">Send</button>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Chat',
  data() {
      return {
          title: '',
          msg: [

          ]
      }
  },
  methods: {
        send() {
            if (this.title!='') {
              this.msg.push(this.title)
              this.title=''
            }
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.chat {
    width: 100%;
    margin: 0;
    padding: 0;
}
.screen {
    width: 80%;
    height: 100%;
    margin: auto;
    position: absolute;
    left: 10%;
}

#header {
    height: 60px;
    width: 80%;
    background-color: black;
    position: fixed;
    top: 0;
    z-index: 1;
}

.avt {
    width: 40px;
    height: 40px;
    background-color: #ccc;
    position: absolute;
    top: 10px;
    left: 3%;
    border: 1px solid #000;
    border-radius: 50%;

}
.avt p {
    color: #fff;
    margin: 9px 14px;
}

.name {
    position: absolute;
    color: #fff;
    margin-left: 90px;
    top: 20px;

}

#container {
    background-color: #ccc;
    height: 100%;
    position: relative;
}

.chatbox {
    height: 87%;
    margin-bottom: 131px;
    overflow-x: hidden;
    padding: 15px;
    margin: 0;
}

li {
    background-color: rgb(231, 178, 245);
    list-style-type: none;
    height: 30px;
    width: 45%;
    padding: 20px;
    margin: 10px 0 5px 0px;
    position: relative;
    border-radius: 10px;
    color: black;
}

li:nth-child(2n) {
    background-color: rgb(129, 131, 236);
    list-style-type: none;
    height: 30px;
    width: 45%;
    padding: 20px;
    margin: 10px 0 5px 0px;
    margin-left: 52%;
    position: relative;
    border-radius: 10px;
    right: 10px;
}

.send {
    height: 10%;
    width: 80%;
    background-color: black;
    position: fixed;
    bottom: 0;
    z-index: 1;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
}

.input-mess {
    width: 65%;
    height: 60%;
    padding: 10px;
    /*margin: 12px 0 12px 90px; */
    /* margin-left: 16%;  */
    border: none;
    border-radius: 8px;
}
 .btsend {
    width: 6%;
    margin-left: 8px;
    border: 1px solid #000;
    border-radius: 7px;
 }

</style>
