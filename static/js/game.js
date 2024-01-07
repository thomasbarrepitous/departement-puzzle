class GameMap {
  constructor(stateMap, map) {
    this.map = map;
    this.departementCodes = stateMap;
  }
}

class Game {
  constructor(stateMap) {
    this.departementStack = [];
    this.departementCodes = stateMap;
    this.currentDepartement = 0;
    this.gameMap = new GameMap(stateMap);
    this.timer = new Timer();
    this.isRunning = false;
  }

  reset() {
    this.initializeStack();
    this.initializeMap();
    this.timer.stop();
  }

  initializeMap() {
    // Reset the map
    // this.gameMap.display();
    return this;
  }

  initializeStack() {
    // Randomize the Departement map into a stack
    this.departementStack = this.shuffle(Object.keys(this.departementCodes));
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

  getNextDepartement() {
    // Pop a Departement from the stack
    this.currentDepartement = this.departementStack.pop();
    document.getElementById("departementDisplay").innerHTML =
      "Departement: " + getDepartementName(this, this.currentDepartement);
    return this.currentDepartement;
  }

  start() {
    // Get a random Departement
    this.initializeStack();
    this.timer.start();
    this.isRunning = true;
    this.getNextDepartement();
  }

  stop() {
    this.reset();
    this.isRunning = false;
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

function getDepartementName(game, code) {
  let selectedDepartement;
  if (code.slice(0, 3) != "FR-") {
    code = "FR-" + code;
  }
  for (const key in game.departementCodes) {
    if (game.departementCodes[key].code === code) {
      selectedDepartement = game.departementCodes[key].name;
      break;
    }
  }
  return selectedDepartement;
}

const button = document.getElementById("toggleGameButton");
function toggleButton(button, game) {
  if (button.innerHTML === "Start") {
    startClick(button, game);
  } else {
    stopClick(button, game);
  }
}
