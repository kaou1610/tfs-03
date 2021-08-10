const chatArea = document.querySelector('.chatbox')

function scroll() {
  chatArea.scroll(0,chatArea.scrollHeight);
}

function btnSend() {
    var x = document.getElementById('msg').value;
    if (x != '') {
        var y = document.createElement('li')
        var t= document.createTextNode(x)
        y.appendChild(t)
        document.querySelector('.chatbox').appendChild(y);
        document.getElementById('msg').value=''
        scroll()
    }
    

}