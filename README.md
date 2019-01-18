# SWAPI 项目文档

## Usage:

- 前端：使用JavaScript语言编写；
- 后端：使用Go语言实现，运行在localhost:8080，使用命令`go get github.com/SCGroupX/SWAPI/server`安装；

## Description:

本项目模仿`https://swapi.co/`实现SWAPI相关的功能，添加了资源的授权访问部分的内容。

- 数据来源：

  资源数据均抓取于源网站 `https://swapi.co/`；

- 资源类型：

  `films`、`people`、`planets`、`species`、`starships`；

- 授权访问采用 jwt 验证；

- 服务API：

  - `films`资源 
    - `/api/films/?page={id} GET `
    - `/api/films/pages GET `
    - `/api/films/{id} GET `
  - `people`资源 
    - `/api/people/?page={id} GET `
    - `/api/people/pages GET `
    - `/api/people/{id} GET `
  - `planets`资源 
    - `/api/planets/?page={id} GET `
    - `/api/planets/pages GET `
    - `/api/planets/{id} GET `
  - `species`资源 
    - `/api/species/?page={id} GET `
    - `/api/species/pages GET `
    - `/api/species/{id} GET `
  - `starships`资源 
    - `/api/starships/?page={id} GET `
    - `/api/starships/pages GET `
    - `api/starships/{id} GET`

