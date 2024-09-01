# go-live-reload-docker

## Live reload para aplicação Go com Docker (CompileDaemon)

Este repositório oferece um exemplo prático de como implementar o live reload em uma aplicação Go usando Docker. Com o Docker, a configuração do ambiente de desenvolvimento é simplificada e o live reload permite que desenvolvedores observem instantaneamente as alterações feitas no código, sem a necessidade de reiniciar o servidor manualmente.

## Pré-requisitos

Certifique-se de ter as ferramentas abaixo instaladas em sua máquina:

[Go](https://go.dev/doc/)

[Docker](https://www.docker.com/)

[Docker-Compose](https://docs.docker.com/compose/)

## Como Executar

Para executar este projeto, siga estas etapas:

Clone o repo:

```bash
git clone https://github.com/seu-usuario/go-live-reload-docker.git
cd go-live-reload-docker
```

Construa a imagem Docker:

```bash
docker-compose build
```

Inicie o contêiner Docker:

```bash
docker-compose up -d
```

## Onde Acessar

Depois de iniciar o contêiner Docker, você pode validar a aplicação em seu navegador acessando:

http://localhost:8080
