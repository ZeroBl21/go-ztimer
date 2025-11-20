# go-ztimer

> [!WARNING]  
> This project is a personal experiment and not intended for general use.

A TUI Pomodoro app that tracks your time spent on pomodoros and breaks. It supports in-memory storage or SQLite, depending on how you build it. The app features graphical representations using `termdash` and a CLI powered by `Cobra` and `Viper`.

## Features
- Pomodoro and break tracking
- Graphical visualization with `termdash`
- CLI management with `Cobra` and `Viper`
- Storage options: in-memory or SQLite

## Installation
```sh
# Clone the repository
git clone https://github.com/zerobl21/go-ztimer.git
cd go-ztimer

# Build the project
go build -o go-ztimer

# System wide install
go install
```

## Usage
```sh
# Start a Pomodoro session with default settings
./go-ztimer
```

```sh
# Customize Pomodoro durations and database file
./go-ztimer pomo --pomo 30m --short 10m --long 20m --db mydatabase.db
```

## License
This project is for personal use only and is not intended for general usage.

