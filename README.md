## Benchmarks Broadcast Implementation in Golang

This repo holds three different implementations of a "Broadcaster" in Golang and benchmark tests for them.

---------------
### Types
#####LinkedBroadcaster
#####AsyncBroadcaster
#####SyncBroadcaster

------------

###Results

These benchmark tests were ran on 64GB / 20 cores Droplet from DigitalOcean.


### Golang 1.4beta1

```

➜ go version
go version go1.4beta1 linux/amd64


➜ GOMAXPROCS=1 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages           100000             26432 ns/op            1968 B/op         37 allocs/op
AsyncBroadcaster1Subscribers5Messages            100000             37263 ns/op            2627 B/op         27 allocs/op
SyncBroadcaster1Subscribers5Messages             100000             54825 ns/op            2146 B/op         26 allocs/op
LinkedBroadcaster2Subscribers10Messages           50000             26214 ns/op            2792 B/op         60 allocs/op
AsyncBroadcaster2Subscribers10Messages            30000             61709 ns/op            8642 B/op         64 allocs/op
SyncBroadcaster2Subscribers10Messages             10000            149002 ns/op            4032 B/op         45 allocs/op
LinkedBroadcaster200Subscribers10Messages          3000            510182 ns/op           31308 B/op        654 allocs/op
AsyncBroadcaster200Subscribers10Messages            300           3951471 ns/op          670585 B/op       4229 allocs/op
SyncBroadcaster200Subscribers10Messages            3000            464353 ns/op           49037 B/op        646 allocs/op
LinkedBroadcaster20Subscribers150Messages          5000            330601 ns/op           25545 B/op        674 allocs/op
AsyncBroadcaster20Subscribers150Messages           1000           1364239 ns/op          190347 B/op       1259 allocs/op
SyncBroadcaster20Subscribers150Messages            2000            987611 ns/op          159993 B/op        949 allocs/op
LinkedBroadcaster200Subscribers1000Messages           1        1913537185 ns/op        268763368 B/op   1539470 allocs/op
AsyncBroadcaster200Subscribers1000Messages           50         223573911 ns/op        42499463 B/op     213437 allocs/op
SyncBroadcaster200Subscribers1000Messages            50          28963061 ns/op         4379486 B/op      26238 allocs/op
ok      _benchmark-broadcast-implementations      515.851s


➜ GOMAXPROCS=2 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages-2         100000             21227 ns/op            1971 B/op         37 allocs/op
AsyncBroadcaster1Subscribers5Messages-2          100000             31860 ns/op            2646 B/op         27 allocs/op
SyncBroadcaster1Subscribers5Messages-2           100000             35297 ns/op            2219 B/op         25 allocs/op
LinkedBroadcaster2Subscribers10Messages-2        100000             17252 ns/op            2905 B/op         60 allocs/op
AsyncBroadcaster2Subscribers10Messages-2          30000             47429 ns/op            8642 B/op         64 allocs/op
SyncBroadcaster2Subscribers10Messages-2           20000             98888 ns/op            5142 B/op         45 allocs/op
LinkedBroadcaster200Subscribers10Messages-2        2000            648436 ns/op           31376 B/op        655 allocs/op
AsyncBroadcaster200Subscribers10Messages-2          500           5367622 ns/op          726708 B/op       4232 allocs/op
SyncBroadcaster200Subscribers10Messages-2          3000            779924 ns/op           39811 B/op        646 allocs/op
LinkedBroadcaster20Subscribers150Messages-2       10000            221864 ns/op           25764 B/op        677 allocs/op
AsyncBroadcaster20Subscribers150Messages-2         1000           4669045 ns/op          780918 B/op       4067 allocs/op
SyncBroadcaster20Subscribers150Messages-2          1000           1567482 ns/op          145080 B/op        634 allocs/op
LinkedBroadcaster200Subscribers1000Messages-2      1000           1556956 ns/op          186510 B/op       4539 allocs/op
AsyncBroadcaster200Subscribers1000Messages-2         20          61432870 ns/op        10657177 B/op      63763 allocs/op
SyncBroadcaster200Subscribers1000Messages-2         200          20505786 ns/op         3364858 B/op      18067 allocs/op
ok      _benchmark-broadcast-implementations      328.258s


➜ GOMAXPROCS=8 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages-8         100000             23390 ns/op            1972 B/op         37 allocs/op
AsyncBroadcaster1Subscribers5Messages-8          100000             27983 ns/op            2647 B/op         27 allocs/op
SyncBroadcaster1Subscribers5Messages-8            50000             45946 ns/op            2238 B/op         26 allocs/op
LinkedBroadcaster2Subscribers10Messages-8        100000             37364 ns/op            2908 B/op         60 allocs/op
AsyncBroadcaster2Subscribers10Messages-8          20000            108492 ns/op            8287 B/op         64 allocs/op
SyncBroadcaster2Subscribers10Messages-8           10000            105230 ns/op            5808 B/op         45 allocs/op
LinkedBroadcaster200Subscribers10Messages-8        2000            831709 ns/op           31628 B/op        659 allocs/op
AsyncBroadcaster200Subscribers10Messages-8          200           7567951 ns/op          780828 B/op       4225 allocs/op
SyncBroadcaster200Subscribers10Messages-8          3000            755846 ns/op           39853 B/op        647 allocs/op
LinkedBroadcaster20Subscribers150Messages-8       10000            155343 ns/op           25961 B/op        680 allocs/op
AsyncBroadcaster20Subscribers150Messages-8          200          12935791 ns/op         1202801 B/op       6238 allocs/op
SyncBroadcaster20Subscribers150Messages-8          2000            938921 ns/op           75158 B/op        523 allocs/op
LinkedBroadcaster200Subscribers1000Messages-8      1000           1476390 ns/op          178805 B/op       4680 allocs/op
AsyncBroadcaster200Subscribers1000Messages-8        100         671279114 ns/op        61982565 B/op     316370 allocs/op
SyncBroadcaster200Subscribers1000Messages-8           3         750416147 ns/op        82151640 B/op     467890 allocs/op
ok      _benchmark-broadcast-implementations      233.211s


➜  GOMAXPROCS=16 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages-16        100000             28514 ns/op            1971 B/op         37 allocs/op
AsyncBroadcaster1Subscribers5Messages-16          50000             37806 ns/op            2652 B/op         27 allocs/op
SyncBroadcaster1Subscribers5Messages-16           50000             56115 ns/op            2265 B/op         26 allocs/op
LinkedBroadcaster2Subscribers10Messages-16        50000             44255 ns/op            2793 B/op         60 allocs/op
AsyncBroadcaster2Subscribers10Messages-16         10000            109002 ns/op            8299 B/op         64 allocs/op
SyncBroadcaster2Subscribers10Messages-16          20000            126231 ns/op            4599 B/op         44 allocs/op
LinkedBroadcaster200Subscribers10Messages-16       2000           1226104 ns/op           31494 B/op        656 allocs/op
AsyncBroadcaster200Subscribers10Messages-16         300           7019840 ns/op          777832 B/op       4233 allocs/op
SyncBroadcaster200Subscribers10Messages-16         1000           1048664 ns/op           40204 B/op        652 allocs/op
LinkedBroadcaster20Subscribers150Messages-16      10000            277616 ns/op           25840 B/op        678 allocs/op
AsyncBroadcaster20Subscribers150Messages-16         200           7196917 ns/op         1336903 B/op       6370 allocs/op
SyncBroadcaster20Subscribers150Messages-16         3000            797929 ns/op           57811 B/op        523 allocs/op
LinkedBroadcaster200Subscribers1000Messages-16     1000           2286091 ns/op          178775 B/op       4687 allocs/op
AsyncBroadcaster200Subscribers1000Messages-16        10         130006831 ns/op        23795390 B/op     116599 allocs/op
SyncBroadcaster200Subscribers1000Messages-16         50          55689028 ns/op         5084004 B/op      30239 allocs/op
ok      _benchmark-broadcast-implementations      96.386s


➜ GOMAXPROCS=19 go test --bench="Bench" --benchmem --benchtime 2s

LinkedBroadcaster1Subscribers5Messages-19        200000             26725 ns/op            1973 B/op         37 allocs/op
AsyncBroadcaster1Subscribers5Messages-19         200000             35927 ns/op            2656 B/op         27 allocs/op
SyncBroadcaster1Subscribers5Messages-19          200000             76913 ns/op            2216 B/op         26 allocs/op
LinkedBroadcaster2Subscribers10Messages-19       200000             36187 ns/op            2903 B/op         60 allocs/op
AsyncBroadcaster2Subscribers10Messages-19         20000            171261 ns/op            8965 B/op         64 allocs/op
SyncBroadcaster2Subscribers10Messages-19          20000            292127 ns/op            4032 B/op         45 allocs/op
LinkedBroadcaster200Subscribers10Messages-19       5000            796889 ns/op           31387 B/op        655 allocs/op
AsyncBroadcaster200Subscribers10Messages-19         200          19994359 ns/op          844535 B/op       4228 allocs/op
SyncBroadcaster200Subscribers10Messages-19         3000           1903834 ns/op           39946 B/op        648 allocs/op
LinkedBroadcaster20Subscribers150Messages-19      10000            222980 ns/op           25739 B/op        677 allocs/op
AsyncBroadcaster20Subscribers150Messages-19         100          20089314 ns/op         1014029 B/op       5954 allocs/op
SyncBroadcaster20Subscribers150Messages-19         1000           3715871 ns/op           57885 B/op        523 allocs/op
LinkedBroadcaster200Subscribers1000Messages-19     1000           2756329 ns/op          177972 B/op       4670 allocs/op
AsyncBroadcaster200Subscribers1000Messages-19       100        1459779370 ns/op        61821377 B/op     315443 allocs/op
SyncBroadcaster200Subscribers1000Messages-19          3        2152116114 ns/op        70322189 B/op     401113 allocs/op
ok      _benchmark-broadcast-implementations      334.924s



➜  GOMAXPROCS=20 go test --bench="Bench" --benchmem --timeout 1h
LinkedBroadcaster1Subscribers5Messages-20        100000             23863 ns/op            1971 B/op         37 allocs/op
AsyncBroadcaster1Subscribers5Messages-20          30000             35310 ns/op            2674 B/op         27 allocs/op
SyncBroadcaster1Subscribers5Messages-20          100000             56431 ns/op            2270 B/op         26 allocs/op
LinkedBroadcaster2Subscribers10Messages-20        50000             32115 ns/op            2794 B/op         60 allocs/op
AsyncBroadcaster2Subscribers10Messages-20         20000            105322 ns/op            8145 B/op         64 allocs/op
SyncBroadcaster2Subscribers10Messages-20          10000            113347 ns/op            5453 B/op         45 allocs/op
LinkedBroadcaster200Subscribers10Messages-20       2000            675250 ns/op           31648 B/op        659 allocs/op
AsyncBroadcaster200Subscribers10Messages-20         300           7949583 ns/op          732240 B/op       4244 allocs/op
SyncBroadcaster200Subscribers10Messages-20         1000           1387634 ns/op           40132 B/op        651 allocs/op
LinkedBroadcaster20Subscribers150Messages-20      10000            235048 ns/op           25846 B/op        678 allocs/op
AsyncBroadcaster20Subscribers150Messages-20         100          10638960 ns/op         1217632 B/op       5849 allocs/op
SyncBroadcaster20Subscribers150Messages-20         2000           1031371 ns/op           71705 B/op        523 allocs/op
LinkedBroadcaster200Subscribers1000Messages-20      500           2673714 ns/op          180555 B/op       4700 allocs/op
AsyncBroadcaster200Subscribers1000Messages-20       100         462235444 ns/op        58985757 B/op     297312 allocs/op
SyncBroadcaster200Subscribers1000Messages-20        100         241017198 ns/op        20575512 B/op     106552 allocs/op
ok      _benchmark-broadcast-implementations      165.981s

```

------

### Golang 1.3.3


```golang
➜ go version
go version go1.3.3 linux/amd64


➜ GOMAXPROCS=1 go test --bench="Bench" --benchmem --timeout 1h
LinkedBroadcaster1Subscribers5Messages           100000             20683 ns/op            3098 B/op         31 allocs/op
AsyncBroadcaster1Subscribers5Messages            100000             32264 ns/op            2222 B/op         17 allocs/op
SyncBroadcaster1Subscribers5Messages             100000            103569 ns/op            1940 B/op         16 allocs/op
LinkedBroadcaster2Subscribers10Messages          100000             19958 ns/op            5043 B/op         50 allocs/op
AsyncBroadcaster2Subscribers10Messages            20000             67464 ns/op            6346 B/op         35 allocs/op
SyncBroadcaster2Subscribers10Messages             10000            263390 ns/op            3393 B/op         25 allocs/op
LinkedBroadcaster200Subscribers10Messages          5000            347224 ns/op           71737 B/op        842 allocs/op
AsyncBroadcaster200Subscribers10Messages            500           5767691 ns/op          628566 B/op       2482 allocs/op
SyncBroadcaster200Subscribers10Messages            5000            491178 ns/op           39436 B/op        628 allocs/op
LinkedBroadcaster20Subscribers150Messages         10000            225689 ns/op           56277 B/op        544 allocs/op
AsyncBroadcaster20Subscribers150Messages           2000           1640351 ns/op          158733 B/op        623 allocs/op
SyncBroadcaster20Subscribers150Messages            1000          11378568 ns/op         1002918 B/op       3385 allocs/op
LinkedBroadcaster200Subscribers1000Messages           1        6955977458 ns/op        639434528 B/op   1796477 allocs/op
AsyncBroadcaster200Subscribers1000Messages            1        1118243109 ns/op        100974184 B/op    359512 allocs/op
SyncBroadcaster200Subscribers1000Messages           100          10040688 ns/op          907585 B/op       3726 allocs/op
ok      _benchmark-broadcast-implementations      312.432s


➜ GOMAXPROCS=2 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages-2         100000             19784 ns/op            3098 B/op         31 allocs/op
AsyncBroadcaster1Subscribers5Messages-2          100000             33315 ns/op            2222 B/op         17 allocs/op
SyncBroadcaster1Subscribers5Messages-2           100000             56527 ns/op            1940 B/op         16 allocs/op
LinkedBroadcaster2Subscribers10Messages-2        200000             28264 ns/op            5042 B/op         50 allocs/op
AsyncBroadcaster2Subscribers10Messages-2          50000             56609 ns/op            7018 B/op         35 allocs/op
SyncBroadcaster2Subscribers10Messages-2           10000            148528 ns/op            3393 B/op         25 allocs/op
LinkedBroadcaster200Subscribers10Messages-2        5000            511355 ns/op           71741 B/op        842 allocs/op
AsyncBroadcaster200Subscribers10Messages-2          500           4531706 ns/op          561387 B/op       2482 allocs/op
SyncBroadcaster200Subscribers10Messages-2          5000            714000 ns/op           39211 B/op        627 allocs/op
LinkedBroadcaster20Subscribers150Messages-2       10000            134984 ns/op           56107 B/op        543 allocs/op
AsyncBroadcaster20Subscribers150Messages-2          500           6092974 ns/op          755541 B/op       2267 allocs/op
SyncBroadcaster20Subscribers150Messages-2          1000           1902050 ns/op          210760 B/op        808 allocs/op
LinkedBroadcaster200Subscribers1000Messages-2      2000           1336486 ns/op          389019 B/op       3744 allocs/op
AsyncBroadcaster200Subscribers1000Messages-2         20         127110924 ns/op         9591200 B/op      34611 allocs/op
SyncBroadcaster200Subscribers1000Messages-2         200          12845640 ns/op         2046651 B/op       7761 allocs/op
ok      _benchmark-broadcast-implementations      237.480s

```

**The droplet ran out of memory(64gb) with `GOMAXPROCS=8` and golang 1.3.3.**

That is why I cut in half the number of Subscribers in the last 3 tests.
```
➜ GOMAXPROCS=8 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages-8         100000             28911 ns/op            3099 B/op         31 allocs/op
AsyncBroadcaster1Subscribers5Messages-8           50000             53272 ns/op            2178 B/op         17 allocs/op
SyncBroadcaster1Subscribers5Messages-8             50000             81968 ns/op            1936 B/op         16 allocs/op
LinkedBroadcaster2Subscribers10Messages-8         50000             34197 ns/op            5041 B/op         50 allocs/op
AsyncBroadcaster2Subscribers10Messages-8          10000            140888 ns/op            6346 B/op         35 allocs/op
SyncBroadcaster2Subscribers10Messages-8           10000            157383 ns/op            3393 B/op         25 allocs/op
LinkedBroadcaster200Subscribers10Messages-8        5000            783756 ns/op           72049 B/op        843 allocs/op
AsyncBroadcaster200Subscribers10Messages-8          100          12593925 ns/op          728716 B/op       2481 allocs/op
SyncBroadcaster200Subscribers10Messages-8          2000           1030189 ns/op           39181 B/op        627 allocs/op
LinkedBroadcaster20Subscribers150Messages-8       10000            146779 ns/op           56289 B/op        544 allocs/op
AsyncBroadcaster20Subscribers150Messages-8          100          27286847 ns/op          844461 B/op       3062 allocs/op
SyncBroadcaster20Subscribers150Messages-8          1000           1750697 ns/op           47547 B/op        228 allocs/op
LinkedBroadcaster100Subscribers1000Messages-8      2000           1157189 ns/op          356442 B/op       3423 allocs/op
AsyncBroadcaster100Subscribers1000Messages-8        100        1064055638 ns/op        30160035 B/op      99091 allocs/op
SyncBroadcaster100Subscribers1000Messages-8           1        1059058742 ns/op        24110376 B/op      85946 allocs/op
ok      _benchmark-broadcast-implementations      182.872s


➜ GOMAXPROCS=16 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages-16        100000             27942 ns/op            3098 B/op         31 allocs/op
AsyncBroadcaster1Subscribers5Messages-16          50000             50602 ns/op            2179 B/op         17 allocs/op
SyncBroadcaster1Subscribers5Messages-16           50000             97254 ns/op            1938 B/op         16 allocs/op
LinkedBroadcaster2Subscribers10Messages-16       100000             24711 ns/op            5046 B/op         50 allocs/op
AsyncBroadcaster2Subscribers10Messages-16         10000            119958 ns/op            6345 B/op         35 allocs/op
SyncBroadcaster2Subscribers10Messages-16          10000            163193 ns/op            3402 B/op         25 allocs/op
LinkedBroadcaster200Subscribers10Messages-16       5000            505067 ns/op           75110 B/op        842 allocs/op
AsyncBroadcaster200Subscribers10Messages-16         100          11417617 ns/op          561872 B/op       2485 allocs/op
SyncBroadcaster200Subscribers10Messages-16         2000            779904 ns/op           39177 B/op        627 allocs/op
LinkedBroadcaster20Subscribers150Messages-16      10000            238403 ns/op           56166 B/op        543 allocs/op
AsyncBroadcaster20Subscribers150Messages-16         100          27282323 ns/op          813118 B/op       2950 allocs/op
SyncBroadcaster20Subscribers150Messages-16         1000           2022583 ns/op           47623 B/op        229 allocs/op
LinkedBroadcaster100Subscribers1000Messages-16     2000            790998 ns/op          356313 B/op       3419 allocs/op
AsyncBroadcaster100Subscribers1000Messages-16       100         793839620 ns/op        28858009 B/op      94467 allocs/op
SyncBroadcaster100Subscribers1000Messages-16         10         220275744 ns/op         8492276 B/op      30433 allocs/op
ok      _benchmark-broadcast-implementations      158.801s


➜ GOMAXPROCS=20 go test --bench="Bench" --benchmem --timeout 1h

LinkedBroadcaster1Subscribers5Messages-20        100000             31435 ns/op            3098 B/op         31 allocs/op
AsyncBroadcaster1Subscribers5Messages-20          50000             45891 ns/op            2179 B/op         17 allocs/op
SyncBroadcaster1Subscribers5Messages-20           50000             81975 ns/op            1939 B/op         16 allocs/op
LinkedBroadcaster2Subscribers10Messages-20        50000             39811 ns/op            5040 B/op         50 allocs/op
AsyncBroadcaster2Subscribers10Messages-20         10000            109334 ns/op            6347 B/op         35 allocs/op
SyncBroadcaster2Subscribers10Messages-20          10000            162710 ns/op            3422 B/op         25 allocs/op
LinkedBroadcaster200Subscribers10Messages-20       5000            484828 ns/op           72069 B/op        843 allocs/op
AsyncBroadcaster200Subscribers10Messages-20         200          18129386 ns/op          560581 B/op       2480 allocs/op
SyncBroadcaster200Subscribers10Messages-20         2000           1841020 ns/op           39177 B/op        627 allocs/op
LinkedBroadcaster20Subscribers150Messages-20      10000            223173 ns/op           56292 B/op        544 allocs/op
AsyncBroadcaster20Subscribers150Messages-20         100          24076729 ns/op          832167 B/op       3018 allocs/op
SyncBroadcaster20Subscribers150Messages-20         1000           1180329 ns/op           47628 B/op        229 allocs/op
LinkedBroadcaster100Subscribers1000Messages-20     2000            727146 ns/op          356407 B/op       3420 allocs/op
AsyncBroadcaster100Subscribers1000Messages-20       100         959738891 ns/op        27220439 B/op      88656 allocs/op
SyncBroadcaster100Subscribers1000Messages-20         10         323884987 ns/op         8755200 B/op      31369 allocs/op
ok      _benchmark-broadcast-implementations      182.859s

```
