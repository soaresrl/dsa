# Lista 4 - Estruturas de Dados - Hash Estendível

## Árvore de arquivos

<pre>
├── lista_01
├── lista_03
├── lista_04
    └── extendible_hash.go
├── README.md
└── pkg
    ├── bank/
    ├── circ_doubly_list/
    ├── circular_linked_list/
    ├── het_linked_list/
    ├── linked_list/
    ├── ord_doubly_list/
    ├── ordered_linked_list/
    └── static_hash
    └── extendible_hash
        ├── extendible_hash.go
        └── extendible_hash_test.go
</pre>

## Sobre a implementação

Foi implementada a Estrutura de _Hash Estendível_ a função de _hash_ já retorna, para o _Global Depth_ atual da estrutura o índice do bucket que contém o valor passado como chave.

## Programas referentes as Questões

Foram mantidos os arquivos referentes a Lista 01 e 03, de modo a manter o registro incremental das listas. A estrutura de dados de _hash estendível_ com as operações solicitadas foram implementadas 
e estão incluídas na pasta `pkg/extendible_hash/extendible_hash.go`.

O programa para execução contendo a função _main_ com as operações solicitadas encontra-se na pasta `lista_04/extendible_hash.go`.

## Comandos

Para rodar a função _main_ relacionada a Lista 04 deve-se chamar o comando `go run lista_04/extendible_hash.go`.
