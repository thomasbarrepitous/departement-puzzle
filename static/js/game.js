class GameObserver {
  constructor(codeToNameMap) {
    this.game = new Game(codeToNameMap);
    this.timer = new Timer();
  }

  start() {
    this.timer.start();
    this.game.start();
  }

  stop() {
    this.timer.stop();
    this.game.stop();
  }

  next() {
    this.game.getNextDepartement();
  }
}

class Game {
  constructor(codeToNameMap) {
    this.departementStack = [];
    this.codeToNameMap = codeToNameMap;
    this.currentDepartement = 0;
    this.isRunning = false;
  }

  reset() {
    this.initStack();
  }

  initStack() {
    // Randomize the Departement map into a stack
    this.departementStack = this.shuffle(Object.keys(this.codeToNameMap));
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
      "Departement: " + this.codeToNameMap[this.currentDepartement].name;
    return this.currentDepartement;
  }

  start() {
    // Get a random Departement
    this.initStack();
    this.getNextDepartement();
    this.isRunning = true;
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

function startClick(button, gameObs) {
  button.innerHTML = "Stop";
  button.classList.remove("start");
  button.classList.add("stop");
  gameObs.start();
}

function stopClick(button, gameObs) {
  button.innerHTML = "Start";
  button.classList.remove("stop");
  button.classList.add("start");
  gameObs.stop();
}

function getDepartementFromCode(map, code) {
  let selectedDepartement;
  Object.keys(map).forEach((key) => {
    if (map[key].code === code) {
      selectedDepartement = map[key];
    }
  });
  return selectedDepartement;
}

const button = document.getElementById("toggleGameButton");
function toggleButton(button, gameObs) {
  if (button.innerHTML === "Start") {
    startClick(button, gameObs);
  } else {
    stopClick(button, gameObs);
  }
}
