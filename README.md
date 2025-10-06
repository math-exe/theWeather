# theWeather CLI

![theWeather CLI Banner](https://raw.githubusercontent.com/math-exe/theWeather/master/assets/theweather-banner.png) Um cliente de linha de comando (CLI) simples, rápido e elegante para consultar informações meteorológicas, construído com Go e pensado para ser extensível.

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

---

## Funcionalidades

-   **Clima Atual:** Obtenha a temperatura, descrição do tempo e outras informações em tempo real.
-   **Dados Detalhados:** Consulte temperaturas mínima/máxima, velocidade e direção do vento.
-   **Interface Colorida:** Saída formatada com cores para melhor legibilidade no terminal.
-   **Arquitetura Limpa:** Projetado com uma arquitetura modular e testável para fácil expansão.
-   **Configuração Simples:** Utiliza variáveis de ambiente para gerenciamento seguro da chave de API.

---

## Começando

Siga estas instruções para ter o projeto rodando na sua máquina local.

### Pré-requisitos

Você vai precisar de uma chave de API (API Key) do **OpenWeatherMap**.

1.  Crie uma conta gratuita no site [openweathermap.org](https://openweathermap.org).
2.  Navegue até a seção "API keys" no seu painel e copie sua chave.
3.  Crie um arquivo chamado `.env` na raiz do projeto e adicione sua chave nele:

    ```ini
    # Arquivo: .env
    OPENWEATHER_API_KEY="sua_chave_de_api_aqui"
    ```

### Instalação

1.  Clone o repositório:
    ```bash
    git clone [https://github.com/seu-usuario/theWeather.git](https://github.com/seu-usuario/theWeather.git)
    cd theWeather
    ```

2.  Construa o binário do projeto. O executável `theWeather` será criado na raiz do projeto.
    ```bash
    go build -o theWeather ./cmd/weather
    ```

---

##  Como Usar

Execute o programa a partir da raiz do projeto, passando o nome da cidade como argumento.

```bash
go run .\cmd\weather\ "Cidade"
```

### Atenção
1.   Casos em que a cidade tem nome composto ou acentuação, apenas passe os dentro das aspas (com ou sem acentuação).
```bash
go run .\cmd\weather\ "São Paulo" ou go run .\cmd\weather\ "Sao Paulo" irão funcionar
```
2.  A consulta funciona tanto com o nome original da cidade quanto o nome traduzido.
```bash
go run .\cmd\weather\ "New York" ou go run .\cmd\weather\ "Nova Iorque" irão funcionar
```