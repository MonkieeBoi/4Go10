# 4Go10

Simple cross-platform game to train number memory written in Golang.

## Building

Clone and cd into the repo code, depending on the desired platform the steps to build differ.

```
git clone https://github.com/MonkieeBoi/4Go10
cd 4Go10
```

Fyne is required to build the application and should be installed using

```
go install fyne.io/tools/cmd/fyne@latest
```


### Desktop

```
fyne package -os darwin -icon Icon.png
fyne package -os linux -icon Icon.png
fyne package -os windows -icon icon.png
```

### Mobile

```
fyne package -os android -app-id com.monkieeboi.fourgoten -icon Icon.png
fyne package -os ios -app-id com.monkieeboi.fourgoten -icon Icon.png
```
