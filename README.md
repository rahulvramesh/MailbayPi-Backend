MailbayPi-Backend
=================

Tiny SMTP server written in GoLang for `raspberrypi` , Based on [go-guerrilla](github.com/flashmob/go-guerrilla).

#### Installation
Required OS : `Linux raspberrypi 4.19.75-v7l+ #1270 SMP Tue Sep 24 18:51:41 BST 2019 armv7l`

#### Download Binary File:

    wget https://github.com/rahulvramesh/MailbayPi-Backend/releases/download/1.0/mailbaypi


#### Setup Config File:

    wget https://raw.githubusercontent.com/rahulvramesh/MailbayPi-Backend/master/guerrillad.conf.json.sample

    mv guerrillad.conf.json.sample guerrillad.conf.json

#### Run The File:
    ./mailbaypi

`Makesure your incoming port is open`
