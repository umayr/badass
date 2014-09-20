
Badass
=======

A tiny script that can make you look *badass* on github. 

It generates false commits having older dates, so when you push your commits to github; your [contribution graph](https://help.github.com/articles/viewing-contributions-on-your-profile-page#contributions-calendar) gets smth similar to other badass devs. If your graph doesn't get affected after execution, then you should read this [article](https://help.github.com/articles/why-are-my-contributions-not-showing-up-on-my-profile). Cheers ^^


Instructions
----

```sh
# Move to any repo you want to generate commits for
$ cd path/to/repo
$ go get github.com/umayr/badass
$ go build github.com/umayr/badass
$ chmod +x badass.exe
$ badass --date 02/21/2013 --seed 50
...
$ git push origin master

```

Make sure you have set [GOPATH](https://code.google.com/p/go-wiki/wiki/GOPATH). Later you can remove the bins or whatever.


Version
----

0.2

Flags
-----

Badass takes various flags as input

* **--saturation, -s** : Higher the value of seeds, higher the saturations of commits. *Default: 72*
* **--columns, -c** - Denotes one column of contribution graph. *Default: 81*
* **--date, -d** - Start date for commits *Default: 1 Year*
* **--max-commits, -m** - Maximum commits per day *Default: 74*


License
----

MIT

Contact
-----

Twitter: [@Bandooqwala](https://twitter.com/Bandooqwala)


Inspired by [this](https://github.com/bd808/profile-life) repo. PRs are welcomed.
