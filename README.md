# Go Horse Racing 

Read original idea here: [Learning concurrency in Golang](https://chemaclass.com/blog/learning-concurrency-in-golang/)

### A horse racing emulator

I remember building a similar game in Java (back in 2012) when I was learning multithreading, and I thought it would be a great opportunity to do it again with the modern Go language.

I built a terminal game emulator that mimics one horse racing. Each horse is a goroutine that runs in a shared bidimensional matrix. Once a horse reaches the end, it is notified to a shared channel between all other horses -running in different processes- and they all stop, showing in the terminal the winner of the race.

## Run the game

```bash
go run main.go
```

## Testing

This command will recursively search for tests in the current directory and its subdirectories and run them.
```bash
go test ./...
```

Alternatively, you can specify the relative path to the directory containing your tests:
```bash
go test ./tests
```

## TODO (Ideas)

1. (easy) use os.ENV to define maxSleepDelay, renderDelay, lines, lineLength
2. (hard) Do not clear the entire board everytime. Instead leave the Board fixed, and use the cursor to access to the position of the Horse, remove its name and place it in the name position (we should lock using Mutex the cursor, so only one horse can use it at a time)
3. ?

> Other ideas welcome. Feel free to submit your PRs!