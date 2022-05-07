
# how to get swagger example

https://www.youtube.com/watch?v=To9s4dOAz50

# experience

bad

# what's missing

1. i use `autorest --go --input-file=pets.json --namespace=github.com/ilovejs/auto_demo` for golang code gen

2. it output to `generated` folder, which is ugly. Need to put model and client in seperate folder or flat strucuter !!

3. code won't compile after generation..., bad format as well.

4. fix by adding import in IDE

```go
"github.com/ilovejs/autodemoapi/client"
. "github.com/ilovejs/autodemoapi/model"
```

5. there is no go mod, so u have to init

# biggest concern

> where is runnable bash script, docs, online demo (meh, as usual)
> where is ur 1 min demo ?

why youtube tech vid don't have straight sourcecode... (who likes to try powershell client, as they demoed)
why do u hv to watch from begin to the end.. 1hr worthy of vid.

> i just want to know how good/capable. and hv just a little interest in knowing what is autorest.

# pros

1. autorest go implementation looks interesting, you can extend to
    - generate better go
    - generate mock
    proto is better then open api or swagger ? idk

2. let some tricks, especially on generated types.
    - struct cost more time to write from stratch. Swagger win.

# my takes

i started on looking at autorest.go folder structure and mocks. Was expected to learn some good practice.

probably go impl for auto rest is far from mature.. but azure team have kept azure cloud api in standard swagger format

what's possible is auto-gen a go sdk / client for azure...

considering they hv internal hub and docs for such a simple tools.. i believe it's half ass project.

It does have purpose, but msft has no interest to share more openly for greater benefit/impact.
Sounds like goal archieved, zombie project.
You can argue open source has no secret, but hv they do a good job to keep it ?

Probably not :(
