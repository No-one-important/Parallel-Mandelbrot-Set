# Parallel-Mandelbrot-Set

A parallel implementation of the Mandlebrot Set construction

The analysis of speedup can be found at the `docs/` directory, but it is in Bulgarian language only <br>

## Prerequisites

A Golang version compatible with `1.18`

## How to run

`go run main.go -g=<granularity> -p=<parallelism>`, where p is the amount of goroutines that will be started by the
program and g is the number of tasks per thread.

## More command line parameters

* `startX` - X coordinate of the starting point of complex plane segment
* `startY` - Y coordinate of the starting point of complex plane segment
* `endX` - X coordinate of the ending point of complex plane segment
* `endY` - Y coordinate of the ending point of complex plane segment

## An example result from the execution of the program (15000x9000 pixels)

![Example mandelbrot set fragment](https://github.com/No-one-important/Parallel-Mandelbrot-Set/blob/main/result.png?raw=true)
