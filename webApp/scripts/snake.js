const startGame = document.getElementByClass('game-start');

console.log("We loaded")

document.getElementById('start-game').addEventListener('click', event => {

  if (window['WebSocket']) {
    const conn = new WebSocket('ws://' + document.location.host + '/ws');

    var sendkeytoserver = function (e) {
      handler.data.push(e);
      console.log(handler.data);
    }
    sendkeytoserver.data = [];

    window.addeventlistener("keydown", sendkeytoserver);

    //submitWinnerButton.onclick = event => {
    //    conn.send()
    //    gameEndContainer.hidden = false
    //    gameContainer.hidden = true
   // }

    conn.onclose = evt => {
    //    document.find.innerText = 'Connection closed'
    //}

    //conn.onmessage = evt => {
    //    blindContainer.innerText = evt.data
    //}

    conn.onopen = function () {
      conn.send("HI")
      console.log("Hi again")
    }
  }
})
