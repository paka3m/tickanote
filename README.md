# Tickanote
## Description
Tickanote(pron. tíkənəʊt/ ちかのて) is a pure-Go DynamicDNS IP notification tool.

You can automatically notify your Dynamic IP address to MyDNS.jp.

## Installation and Usages
It's easy-peasy!

    go get github.com/paka3m/tickanote
    tickanote start --auth mydns080211:sRuouqA -interval 1h

### Child IDs
if you have child IDs(e.g. mydns3151515:uuuummm, mydns114514:qqqqqq ), command like below.

    tickanote start --auth mydns080211:sRuouqA mydns3151515:uuuummm mydns114514:qqqqqq -interval 1h

## Author
paka3m 

## Licence
MIT