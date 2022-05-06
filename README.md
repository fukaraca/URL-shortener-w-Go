# URL Squeezer

A basic URL shortener written with Go. MongoDB has been used for database requirements. Youtube presentation can be foun [here](https://youtu.be/8fzwr5jl7uw)

```
git clone https://github.com/fukaraca/URL-shortener-w-Go.git
```


Expected POST request is consisted of three form key-value pairs. These are:

| Method | key      | sample value                                                |
|--------|----------|-------------------------------------------------------------|
| POST   | url      | www.google.com                                              |
| POST   | custom   | custom-url-path (optional)                                  |
| POST   | validfor | 23h30m5s             (optional- 0 is default means forever) |

For `validfor` key, only hour, minutes and seconds are accepted. In order to prevent possible errors, days or weeks must be converted to hours and so on.


