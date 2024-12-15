# sisyphus

convert gemtext to html/markdown

## cli usage

create a markdown readme from gemtext:
```
$ sisyphus -f markdown <README.gmi >README.md
```

convert gemtext to html:
```
$ sisyphus <index.gmi >index.html
```

## from go

```
sisyphus.Gem(os.Stdin, os.Stdout, &sisyphus.Html{})
```

