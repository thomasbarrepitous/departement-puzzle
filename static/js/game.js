class GameMap {
  constructor(stateMap, map) {
    this.map = map;
    this.departementCodes = stateMap;
    this.display();
  }

  display() {
    // const map = new jsVectorMap({
    //   selector: "#map",
    //   map: "france_departments",
    //   showTooltip: false,
    //   draggable: false,
    //   zoomOnScroll: false,
    //   zoomButtons: false,
    //   panOnDrag: false,
    //   backgroundColor: "#fff",
    //   onRegionClick: function (event, code) {
    //     for (const key in this.departementCodes) {
    //       if (this.departementCodes[key].code === code) {
    //         console.log(this.departementCodes[key].name);
    //         event.srcElement.style.fill = "green";
    //       }
    //     }
    //   },
    // });
    // this.map = map;
  }
}
class Game {
  constructor(stateMap) {
    this.departmentStack = [];
    this.departementCodes = stateMap;
    this.initializeStack();
    this.gameMap = new GameMap(stateMap);
    this.timer = new Timer();
  }

  reset() {
    this.initializeStack();
    this.initializeMap();
    this.timer.stop();
  }

  initializeMap() {
    // Reset the map
    this.gameMap.display();
    return this;
  }

  initializeStack() {
    // Randomize the department map into a stack
    this.departmentStack = this.shuffle(Object.keys(this.departementCodes));
    return this;
  }

  // Fisher-Yates shuffle
  shuffle(array) {
    let currentIndex = array.length,
      randomIndex;

    // While there remain elements to shuffle...
    while (currentIndex != 0) {
      // Pick a remaining element...
      randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex--;

      // And swap it with the current element.
      [array[currentIndex], array[randomIndex]] = [
        array[randomIndex],
        array[currentIndex],
      ];
    }

    return array;
  }

  getRandomDepartment() {
    if (this.departmentStack.length === 0) {
      // If the stack is empty, refill it with the original map
      console.log("Refilling stack");
      this.initializeStack();
    }

    // Pop a department from the stack
    const randomDepartment = this.departmentStack.pop();
    return randomDepartment;
  }

  start() {
    // Get a random department
    const randomDepartment = this.getRandomDepartment();
    console.log(randomDepartment);
    this.timer.start();
  }

  stop() {
    this.reset();
  }
}

class Timer {
  constructor() {
    this.timer = null;
    this.startTime = null;
    this.stopTime = null;
    this.microseconds = 0;
  }

  reset() {
    this.microseconds = 0;
    this.startTime = null;
    this.stopTime = null;
    this.updateTimerDisplay(0);
  }

  start() {
    this.startTime = Date.now();
    this.timer = setInterval(() => {
      this.microseconds = Date.now() - this.startTime;
      this.updateTimerDisplay(this.microseconds);
    }, 1);
  }

  stop() {
    clearInterval(this.timer);
    this.reset();
    console.log("Timer stopped");
  }

  updateTimerDisplay(time) {
    const seconds = Math.floor((time % 1000000) / 1000);
    const minutes = Math.floor((time % 60000000) / 60000);
    let secondsStr = (seconds % 60).toString();
    let minutesStr = minutes.toString();
    if (seconds % 60 < 10) {
      secondsStr = "0" + secondsStr;
    }
    if (minutes < 10) {
      minutesStr = "0" + minutesStr;
    }
    const timerDisplay = `Timer: ${minutesStr}:${secondsStr}`;
    document.getElementById("timerDisplay").innerHTML = timerDisplay;
  }
}

function startClick(button, game) {
  button.innerHTML = "Stop";
  button.classList.remove("start");
  button.classList.add("stop");
  game.start();
}

function stopClick(button, game) {
  button.innerHTML = "Start";
  button.classList.remove("stop");
  button.classList.add("start");
  game.stop();
}

const button = document.getElementById("toggleGameButton");
function toggleButton(button, game) {
  if (button.innerHTML === "Start") {
    startClick(button, game);
  } else {
    stopClick(button, game);
  }
}
