# manipula-dados

O serviço consistirá em uma aplicação Go que lê um arquivo CSV/TXT de entrada, processa os dados e os persiste em um banco de dados PostgreSQL. A aplicação será empacotada em um contêiner Docker e será executada com o Docker Compose. A entrada será validada usando uma biblioteca externa para validar os CPFs e CNPJs. A aplicação será dividida em pacotes separados para melhor organização e manutenção do código.

Requisitos
Para construir e executar o serviço, você precisará das seguintes ferramentas:

Docker
Docker Compose

Instruções
Siga estas etapas para executar o serviço:

Clone este repositório: git clone https://github.com/seu-usuario/nome-do-repo.git.
Navegue para o diretório do projeto: cd nome-do-repo.
Crie um arquivo .env com as seguintes variáveis de ambiente: