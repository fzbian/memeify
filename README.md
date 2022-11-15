
# Meme generator

Meme generator is a code created with the simple purpose of learning how image management works in this language, this project is quite simple.

## API Reference

#### Get the image

```http
  GET /meme/text1/text2
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `meme` | `string` | the name of the meme to be used |
| `text1` | `string` | the text at the top of the page to be used |
| `text2` | `string` | the text at the bottom of the page to be used |


## Types of memes
 - bad-luck-brian
 - batman-slaps-robin
 - crying-peter-parker
 - filosoraptor
 - gopher
 - kermit-the-frog
 - mr-bean
 - trump
 - willy-wonka
## Demo

The code has been uploaded to repl.it so that anyone can see how the code works.

Input

https://meme-generator.fzbian.repl.co/gopher/golang/for%20life

Output

![](https://github.com/fzbian/meme-generator/blob/main/readmeFiles/golang-forlife.jpg?raw=true)
## Deployment

[![Run on Repl.it](https://repl.it/badge/github/fzbian/meme-generator)](https://repl.it/github/fzbian/meme-generator)

To deploy this project run

```bash
  go run main.go
```


## Acknowledgements

 - [Fogleman by gg](https://github.com/fogleman/gg)
 - [AndresXLP for helping me in my learning](https://github.com/AndresXLP)