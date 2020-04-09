# pomodoro
Simple timer built for personal use

- On every 45 Min, it reminds to drink water and give rest to eyes

NOTE: It assumes that you have VLC installed and you are on Linux.

USAGE: 
```
go run pomodoro.go <number of minutes>
```
OR
```
go build pomodoro.go && ./pomodoro <number of minutes>
```

SAMPLE:
```
➜  pomodoro git:(master) ./pomodoro 60         
Did you exercise for one hour?
true
Yay, you are going great!
It's great to be on time, you are going to have a great day!
These are your goals:
- Become a really really good programmer
- Be fit
- Be loved, wanted, charismatic, convincing, powerful and influencial

Running pomodoro for 60 minutes
59
58
^C
You worked for 2 minutes, but aim was 60 minutes
➜  pomodoro git:(master) 

```
