# sisyphus

convert gemtext to html/markdown

(api, interfaces, cli tool - everything is likely to change drastically 😂

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
// html
sisyphus.Gem(os.Stdin, os.Stdout, &sisyphus.Html{})

// markdown
sisyphus.Gem(os.Stdin, os.Stdout, &sisyphus.Markdown{})
```

