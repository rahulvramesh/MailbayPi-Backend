MailbayPi-Backend
=================

![Go](https://github.com/rahulvramesh/MailbayPi-Backend/workflows/Go/badge.svg?branch=master)

Tiny SMTP server written in GoLang for `raspberrypi` , Based on [go-guerrilla](github.com/flashmob/go-guerrilla).

#### Installation
Required OS : `Linux raspberrypi 4.19.75-v7l+ armv7l`

#### Download Binary File:

    wget https://github.com/rahulvramesh/MailbayPi-Backend/releases/download/1.0/mailbaypi


#### Setup Config File:

    wget https://raw.githubusercontent.com/rahulvramesh/MailbayPi-Backend/master/guerrillad.conf.json.sample 

    mv configs/guerrillad.conf.json.sample guerrillad.conf.json

#### Run The File:
    ./mailbaypi

`make sure your incoming port is open` 


## Contributors ✨

<table>
  <tr>
    <td align="center"><a href="https://github.com/rahulvramesh"><img src="https://avatars0.githubusercontent.com/u/4546714?s=460&u=ac1b3c6f4343c0f783b9c6c59fdefb87616a7856&v=44" width="100px;" alt="4dd3r"/><br /><sub><b>rahulvramesh</b></sub></a><br /><a href="https://github.com/rahulvramesh/MailbayPi-Backend/commits?author=rahulvramesh" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/sivsivsree"><img src="https://avatars0.githubusercontent.com/u/2310755?s=460&u=b92546165caed042f37f616f975628d1ce84a95e&v=4" width="100px;" alt="4dd3r"/><br /><sub><b>sivsivsree</b></sub></a><br /><a href="https://github.com/rahulvramesh/MailbayPi-Backend/commits?author=sivsivsree" title="Code">💻</a></td>
    </tr></table>
