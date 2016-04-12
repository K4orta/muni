# MUNI
-----
[![Build Status](https://travis-ci.org/K4orta/muni.svg?branch=master)](https://travis-ci.org/K4orta/muni)

A library for requesting data from the SF Nextbus API

## Installation
`go get github.com/k4orta/muni`

## Usage
```go
  import (
    "encoding/json"
    "fmt"

      "github.com/k4orta/muni"
  )

  func AllVehicles() {
    vd, _ := muni.GetMultiVehicleData([]string{"N", "L", "J", "KT", "M"})
    out, _ := json.Marshal(vd)
    fmt.Fprint(w, string(out))
  }
```
