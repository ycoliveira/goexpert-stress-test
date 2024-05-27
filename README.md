Um simples testador de carga web escrito em Go, que permite realizar testes de carga em um serviço web específico.

Como Usar
Pré-requisitos
Certifique-se de que você tem o Docker instalado na sua máquina.

Execução
Execute o seguinte comando no seu terminal:

Criando a imagem Docker:
docker build -t stresstest .
Executando o contêiner Docker:
docker run stresstest --url=http://example.com --requests=1000 --concurrency=10
Substitua "http://example.com" pela URL do serviço que você deseja testar. --requests é o número total de requisições a serem feitas e --concurrency é o número de chamadas simultâneas.

Resultado
Após a execução, o programa irá gerar um relatório que inclui:

Relatório de Teste:
Tempo total gasto na execução:
Quantidade total de requests realizados:
Quantidade de requests com status HTTP 200:
Distribuição de outros códigos de status HTTP:
