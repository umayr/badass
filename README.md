
Badass
=======

A tiny script that can make you look *badass* on github. 

It generates false commits having older dates, so when you push your commits to github; your [contribution graph](https://help.github.com/articles/viewing-contributions-on-your-profile-page#contributions-calendar) gets smth similar to other badass devs. If your graph doesn't get affected after execution, then you should read this [article](https://help.github.com/articles/why-are-my-contributions-not-showing-up-on-my-profile). Cheers ^^

Screenshot
----
![Screenshot](http://i.imgur.com/6evEG7N.png)

Instructions
----

```sh
# Move to any repo you want to generate commits for
$ cd path/to/repo
$ go get github.com/umayr/badass
$ go build github.com/umayr/badass
$ chmod +x badass.exe
$ badass --date 02/21/2013 --saturation 50 --max-commits 60
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

The MIT License (MIT)

Copyright (c) 2014 Umayr Shahid ([umayrr@hotmail.co.uk](mailto:umayrr@hotmail.co.uk))

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Contact
-----

Twitter: [@Bandooqwala](https://twitter.com/Bandooqwala)


Inspired by [this](https://github.com/bd808/profile-life) repo. PRs are welcomed.
