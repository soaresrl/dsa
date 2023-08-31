# Lista 1 - Estruturas de Dados

## Árvore de arquivos

<pre>
├── 01_singly_list.go
├── 02_ordered.go
├── 03_ordered_doubly_list.go
├── 04_circular_linked_list.go
├── 05_circ_doubly_list.go
├── 06_bank.go
├── README.md
└── pkg
    ├── bank
        ├── bank.go
        └── bank_account_test.go
    ├── circ_doubly_list
        ├── circ_doubly_list.go
        └── circ_doubly_list_test.go
    ├── circular_linked_list
        ├── circular_linked_list.go
        └── circular_linked_list_test.go
    ├── het_linked_list
        ├── het_linked_list.go
        └── het_linked_list_test.go
    ├── linked_list
        ├── linked_list.go
        └── linked_list_test.go
    ├── ord_doubly_list
        ├── ord_doubly_list.go
        └── ord_doubly_list_test.go
    └── ordered_linked_list
        ├── ordered_linked_list.go
        └── ordered_linked_list_test.go
</pre>

## Programas referentes as Questões

Os programas numerados de 01 a 06 seguidos de uma nomenclatura referente a estrutura de dados utilizada na questão, 
são utilizados como respostas (exemplos de utilização) das estruturas de dados implementadas.

As estruturas de dados implementadas encontram-se na pasta `pkg` e cada subpasta contém as estruturas implementadas e um arquivo de testes com testes unitários.

## Comandos

Para rodar os testes de uma estrutura de dados específica, é necessário navegar para a pasta da estrutura e rodar o comando `go test -v`.

Para rodar as questões pode-se rodar o comando `go run <nome_do_arquivo>` (e.g.: `go run 01_singly_list.go`).
