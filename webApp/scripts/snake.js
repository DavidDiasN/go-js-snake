function removeFirst(arr) {
    return arr.map((subArray, index) => index > 0 ? subArray : null).filter(Boolean);
}

function removeLast(arr) {
    let l = arr.length - 1
    return arr.filter((_, index) => index != l);
}

function prepend(arr, element) {
    return [[...element], ...arr];
}

function append(arr, element) {
    const newArr = [...arr];
    
    if (Array.isArray(element) && element.length > 2) {
        element.forEach(item => newArr.push(item));
    } else {
        newArr.push(element);
    }
    
    return newArr;
}

function sliceArray(arr, start, end) {
  return arr.filter((_, index) => index >= start && index <= end)
}


const startGame = document.getElementById("game-start");

startGame.addEventListener('click', _ => {

  if (window['WebSocket']) {
    const conn = new WebSocket('ws://' + document.location.host + '/ws');
    let score = 0;
    startGame.innerText = "Restart Game"
    startGame.setAttribute('disabled', 'Restart Game');

    let foodLocation = [12, 12];
    let snakeState = [];
    snakeState.push([12,12])
    let rowcol = `row${snakeState[0][0]}-col${snakeState[0][1]}` 
    document.getElementById(rowcol).classList.add('snake-block')

    conn.onopen = function () {
      // Send initial data on WebSocket open
      conn.send("Initial message");
    }

    document.addEventListener("keydown", function(event) {
      conn.send(JSON.stringify(event.key));
      });


    conn.onclose = ()  => {
      document.getElementById("score-display").innerText = "Score: " + 0;
      // Document events or actions on WebSocket close
    };

    conn.onmessage = evt => {
      let data = JSON.parse(evt.data);
      
      
      if (data == null || data == NaN) {
        console.log("Decode error")
        return
      }

      if (data === "You Died") {
        console.log("\rYou Died");
        
        rowcol = `row${foodLocation[0]}-col${foodLocation[1]}` 
        document.getElementById(rowcol).classList.remove('food-block')

        for (i = 0; i < snakeState.length; i++) {
          rowcol = `row${snakeState[i][0]}-col${snakeState[i][1]}` 
          document.getElementById(rowcol).classList.remove("snake-block");
        }

        startGame.removeAttribute('disabled');
        conn.close();
        return;
      }

      if (data.length > 2) {
        
        rowcol = `row${foodLocation[0]}-col${foodLocation[1]}` 
        document.getElementById(rowcol).classList.remove('food-block')
        rowcol = `row${data[0][0]}-col${data[0][1]}` 
        document.getElementById(rowcol).classList.add('food-block');

        foodLocation = data[0]

        let l = snakeState.length - 1
        rowcol = `row${snakeState[l][0]}-col${snakeState[l][1]}` 
        document.getElementById(rowcol).classList.remove("snake-block");

        snakeState = removeLast(snakeState);

        snakeState = prepend(snakeState, data[1]);

        snakeState = append(snakeState, sliceArray(data, 2, data.length))

        
        rowcol = `row${snakeState[0][0]}-col${snakeState[0][1]}` 
        document.getElementById(rowcol).classList.add("snake-block");

        for (i = 2; i < data.length; i++) {
          rowcol = `row${data[i][0]}-col${data[i][1]}` 
          document.getElementById(rowcol).classList.add("snake-block");
        }

        score = snakeState.length
        document.getElementById("score-display").innerText = "Score: " + score;
      } else  {


        let l = snakeState.length - 1;

        rowcol = `row${snakeState[l][0]}-col${snakeState[l][1]}` 
        document.getElementById(rowcol).classList.remove("snake-block");
        snakeState = sliceArray(snakeState, 0, l-1)

        snakeState = prepend(snakeState, data[0])

        rowcol = `row${snakeState[0][0]}-col${snakeState[0][1]}` 
        document.getElementById(rowcol).classList.add("snake-block");
        score = snakeState.length
        document.getElementById("score-display").innerText = "Score: " + score;
      }
    };

  }
});

