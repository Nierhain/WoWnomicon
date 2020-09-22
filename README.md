# WoWnomicon
A World of Warcraft dashboard app

### Features
World of Warcraft dashboard
backup your interface
overview of all characters (including reps, lockouts, armory, etc.)

### Technology Stack

- Go
- Nodejs
- Vuejs
- Vuex
- Vue Router

## Building from source

### Dependencies
- Go
- Nodejs
- Wails
- gcc 
- gtk, webkit on Linux

Just follow the steps for installing [Wails](https://wails.app/gettingstarted/). Everything else is handled by yarn/npm

#### Build the binary
navigate to the project folder and compile using
```wails build```

#### Run the dev environment
navigate into the project folder and issue\
```wails serve```

change to the frontend directory and run\
``` 
/* yarn */
yarn serve

/* npm */
npm run serve //for npm
```