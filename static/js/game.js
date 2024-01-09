class GameObserver {
  constructor(codeToNameMap) {
    this.game = new Game(codeToNameMap);
    this.timer = new Timer();
  }

  start() {
    this.reset();
    this.timer.start();
    this.game.start();
    document.getElementById("guessDisplay").innerHTML =
      this.game.codeToNameMap[this.game.currentDepartement].name;
  }

  stop() {
    this.timer.stop();
    this.game.stop();
  }

  reset() {
    this.game.reset();
    this.timer.reset();
    document.getElementById("guessDisplay").innerHTML = "?";
  }

  next() {
    if (this.game.departementStack.length === 0) {
      this.end();
      return;
    }
    this.game.getNextDepartement();
    document.getElementById("guessDisplay").innerHTML =
      this.game.codeToNameMap[this.game.currentDepartement].name;
  }

  skip() {
    const startButton = document.getElementById("startButton");
    if (startButton.classList.contains("hidden") ){
        this.game.skipDepartement();
        this.next();
    }
  }

  // Only called when the game is won
  end() {
    this.stop();
    toggleClick();
    displayWinMessage();
  }

  guess(event, code) {
    const clickedDep = getDepartementFromCode(this.game.codeToNameMap, code);
    if (this.game.isRunning === false) {
      return;
    }
    if (this.game.isGuessCorrect(code)) {
      event.srcElement.style.fill = "green";
      // console.log("Correct department, you clicked on : " + clickedDep.name);
      this.next();
    } else {
      if (event.srcElement.style.fill !== "green") {
        event.srcElement.style.fill = "red";
      }
      // console.log("Wrong department, you clicked on : " + clickedDep.name);
    }
  }

  playAgain() {
    this.reset();
    removeWinMessage();
    toggleClick();
    this.start();
  }
}

class Game {
  constructor(codeToNameMap) {
    this.departementStack = [];
    this.codeToNameMap = codeToNameMap;
    this.currentDepartement = 0;
    this.isRunning = false;
    this.displayedMap = null;
    this.init();
  }

  isGuessCorrect(code) {
    // Is clicked dep same as current dep ?
    return code === this.codeToNameMap[this.currentDepartement].code;
  }

  reset() {
    document.getElementById("map").innerHTML = "";
  }

  init() {
    this.displayedMap = new jsVectorMap({
      selector: "#map",
      map: "france_departments",
      showTooltip: false,
      draggable: true,
      zoomOnScroll: true,
      zoomButtons: false,
      panOnDrag: true,
      backgroundColor: "#1111a",
      onRegionClick: function (event, code) {
        gameObs.guess(event, code);
      },
    });
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
    if (this.departementStack.length === 0) {
      this.super().end();
      return;
    }
    // Pop a Departement from the stack
    this.currentDepartement = this.departementStack.pop();
    return this.currentDepartement;
  }

  skipDepartement() {
    this.departementStack.unshift(this.currentDepartement);
  }

  start() {
    // Get a random Departement
    this.init();
    this.isRunning = true;
    this.getNextDepartement();
  }

  stop() {
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
    const timerDisplay = `${minutesStr}:${secondsStr}`;
    document.getElementById("timerDisplay").innerHTML = timerDisplay;
  }
}

function startClick(button, gameObs) {
  toggleClick();
  gameObs.start();
}

function stopClick(button, gameObs) {
  toggleClick();
  gameObs.stop();
}

function toggleClick() {
  startButton = document.getElementById("startButton");
  stopButton = document.getElementById("stopButton");
  if (startButton.classList.contains("hidden")) {
    stopButton.classList.add("hidden");
    startButton.classList.remove("hidden");
  } else {
    stopButton.classList.remove("hidden");
    startButton.classList.add("hidden");
  }
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

function displayWinMessage() {
  const winMessage = document.getElementById("winMessage");
  const timerWinMessage = document.getElementById("timerWinMessage");
  winMessage.style.display = "block";
  timerWinMessage.innerHTML =
    document.getElementById("timerDisplay").innerHTML + "s !";
  confetti({
    particleCount: 750,
    spread: 150,
    origin: { y: 0.6 },
  });
}

function removeWinMessage() {
  const winMessage = document.getElementById("winMessage");
  // Hide the message and stop animation
  winMessage.style.display = "none";
}

function toggleButton(button, gameObs) {
  if (button.innerHTML === "Start") {
    startClick(button, gameObs);
  }
  if (button.innerHTML === "Stop") {
    stopClick(button, gameObs);
  }
}
