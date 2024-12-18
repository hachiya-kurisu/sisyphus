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
sisyphus.Cook(os.Stdin, os.Stdout, &sisyphus.Html{})

// markdown
sisyphus.Cook(os.Stdin, os.Stdout, &sisyphus.Markdown{})

// convert from a string
sisyphus.Convert("=> somewhere a link", &sisyphus.Html{})

// hooks
flavor := &sisyphus.Html{}
flavor.OnLink(".jpg", hook)
```

