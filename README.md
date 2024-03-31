# Go Horse Racing 

### A horse racing emulator

I wanted to learn a new programming language, so after trying some, I ended up with Golang as one of my favorites, for its simplicity and capabilities. It has features that I haven’t used in years, like multithreading“.

Go is not a multithreading language but concurrency. Although these terms may seem similar to a regular developer when working with them (as the APIs you will use are very similar), they are actually different concepts.

I remember building a similar game in Java (back in 2012) when I was learning multithreading, and I thought it would be a great opportunity to do it again with the modern Go language.

I built a terminal game emulator that mimics one horse racing. Each horse is a goroutine that runs in a shared bidimensional matrix. Once a horse reaches the end, it is notified to a shared channel between all other horses -running in different processes- and they all stop, showing in the terminal the winner of the race.

## How to play?

```bash
go run .
```
