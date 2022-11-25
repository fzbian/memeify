
# Meme generator

Meme generator is a code created with the simple purpose of learning how image management works in this language, this project is quite simple.


## API Reference

#### Custom Meme

In Custom Meme we will indicate the text at the top, the text at the bottom and the url of the image, the api will return the meme done.

#### How to use

```http
  GET /?top=top/?bottom=bottom/?url=url
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `top` | `string` | The text above the image |
| `bottom` | `string` | The text below the image |
| `url` | `string` | The URL of the image to use |

#### Example

Input
```http
  GET /custom/?top=Go&bottom=Is%20the%20best%20language&url=https://nwdistrict.ifas.ufl.edu/nat/files/2021/01/Groundhog.jpg
```
Output

![](http://imgfz.com/i/MUN7gSd.jpg)

#### Local Memes

In this function you will send what meme template you want and the parameters you are given.
#### How to use

```http
  GET /meme/params
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `meme` | `string` | The meme template you want to use |
| `params` | `string` | Depending on the meme you will be asked for several parameters |

#### Templates

| ID | Preview | Params |
| :-------- | :------- | :------------------------- |
| `change_my_mind` | [change_my_mind](https://github.com/fzbian/meme-generator/blob/main/memes/change_my_mind.png?raw=true) | `text` |
| `disappointed_black_guy` | [disappointed_black_guy](https://github.com/fzbian/meme-generator/blob/main/memes/disappointed_black_guy.png?raw=true) | `text1`, `text2` |
| `distracted_boyfriend` | [distracted_boyfriend](https://github.com/fzbian/meme-generator/blob/main/memes/distracted_boyfriend.png?raw=true) | `text1`, `text2`, `text3` |
| `drake` | [drake](https://github.com/fzbian/meme-generator/blob/main/memes/drake.png?raw=true) | `text1`, `text2` |
| `expanding_brain` | [expanding_brain](https://github.com/fzbian/meme-generator/blob/main/memes/expanding_brain.png?raw=true) | `text1`, `text2`, `text3`, `text4` |
| `jason_momoa_henry_cavil` | [jason_momoa_henry_cavil](https://github.com/fzbian/meme-generator/blob/main/memes/jason_momoa_henry_cavil.png?raw=true) | `text1`, `text2` |
| `left_right` | [left_right](https://github.com/fzbian/meme-generator/blob/main/memes/left_right.png?raw=true) | `text1`, `text2` |
| `running_away_balloon` | [running_away_balloon](https://github.com/fzbian/meme-generator/blob/main/memes/running_away_balloon.png?raw=true) | `text1`, `text2`, `text3` |
| `spiderman` | [spiderman](https://github.com/fzbian/meme-generator/blob/main/memes/spiderman.png?raw=true) | `text1`, `text2` |
| `this_is` | [this_is](https://github.com/fzbian/meme-generator/blob/main/memes/this_is.png?raw=true) | `url` |
| `three_headed_dragon` | [three_headed_dragon](https://github.com/fzbian/meme-generator/blob/main/memes/three_headed_dragon.png?raw=true) | `text1`, `text2`, `text3` |
| `undertaker` | [undertaker](https://github.com/fzbian/meme-generator/blob/main/memes/undertaker.png?raw=true) | `text1`, `text2` |
| `grim_reaper_knocking_door` | [grim_reaper_knocking_door](https://github.com/fzbian/meme-generator/blob/main/memes/grim_reaper_knocking_door.png?raw=true) | `url1`, `url2`, `url3`, `url4` |

#### Example

Input
```http
  GET /spiderman/?text1=Python&text2=Slow
```
Output

![](http://imgfz.com/i/OBaGNqh.jpeg)





## Deployment

[![Run on Repl.it](https://repl.it/badge/github/fzbian/meme-generator)](https://repl.it/github/fzbian/meme-generator)

To deploy this project run

```bash
  go run main.go
```


## Acknowledgements

 - [Fogleman by gg](https://github.com/fogleman/gg)
 - [AndresXLP for helping me in my learning](https://github.com/AndresXLP)
