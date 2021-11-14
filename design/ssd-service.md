# Semi-Structured Data Service

## 1. Program Structure

```plantuml
@startuml
ditaa
+-----------------------------------------+
|   servie / interface / ...              | framework
+-----------------------------------------+
|   config / persistence / script / log   | core
+-----------------------------------------+
@enduml
```

## 2. Framework

- service
- interface

## 3. Core

- config
- persistence
- script
- log

## 4. File

```plantuml
@startmindmap
* /
** configs/
***_ default.json
***_ last.json
** logs/
***_ last.log
***_ 2021-07-28T22-30-02.log
***_ 2021-07-28T23-00-01.log
** persistences/
***_ default.json
***_ last.json
***_ 2021-07-28T22-00-00.json
***_ 2021-07-28T23-00-06.json
** scripts/
***_ main.lua
***_ persistence.lua
***_ log.lua
@endmindmap
```
